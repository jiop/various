package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/jiop/various/radix/radix"
	pb "github.com/jiop/various/radixcp/request"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const csvFile = "csv/laposte_hexasmal.csv"

var port = flag.String("port", "49123", "port")

func transformStr(s string) string {
	return strings.ToLower(strings.Replace(s, " ", "_", -1))
}

func loadRadix(csvFile string, radixStore *radix.Tree) error {
	f, err := os.Open(csvFile)
	defer f.Close()

	if err != nil {
		return err
	}
	r := csv.NewReader(f)
	r.Comma = ';'
	_, err = r.Read()
	if err != nil {
		return err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		radixStore.Insert(transformStr(record[1]), record[2])
	}

	return nil
}

type server struct {
	radixStore *radix.Tree
}

func (s *server) run(c chan<- error) {
	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRequestServer(grpcServer, s)

	log.Println("board server listening on port", *port)
	if err := grpcServer.Serve(lis); err != nil {
		c <- err
	}
}

func (s *server) GetCP(ctx context.Context, message *pb.CityMessage) (*pb.PostalCodeMessage, error) {
	res, _ := s.radixStore.Get(message.Name)
	return &pb.PostalCodeMessage{Postalcode: res.(string)}, nil
}

func (s *server) loadData(csvFile string) error {
	f, err := os.Open(csvFile)
	defer f.Close()
	if err != nil {
		return err
	}

	r := csv.NewReader(f)
	r.Comma = ';'

	if _, err = r.Read(); err != nil {
		return err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		s.radixStore.Insert(transformStr(record[1]), record[2])
	}

	return nil
}

func newServer() *server {
	return &server{radixStore: radix.New()}
}

func main() {
	flag.Parse()

	s := newServer()
	if err := s.loadData(csvFile); err != nil {
		log.Fatal(err)
	}

	c := make(chan error)
	go s.run(c)

	for {
		select {
		case v := <-c:
			log.Fatal(v)
		}
	}
}
