package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/context"

	ad "github.com/jiop/various/grpc_agents/advertise"
	pb "github.com/jiop/various/grpc_agents/status"
	"google.golang.org/grpc"
)

type statusServer struct {
	port string
}

type crash interface {
	Error() string
}
type randomCrash struct{}

func (*randomCrash) Error() string {
	return "randomCrash"
}

type classicCrash struct{}

func (*classicCrash) Error() string {
	return "classicCrash"
}

func (s *statusServer) start(c chan crash) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterStatusServer(grpcServer, s)

	// go func() {
	// 	select {
	// 	case <-time.After(3 * time.Second):
	// 		c <- &randomCrash{}
	// 	}
	// }()

	if err := grpcServer.Serve(lis); err != nil {
		c <- &classicCrash{}
	}
}

func (s *statusServer) Get(ctx context.Context, in *pb.None) (*pb.StatusMessage, error) {
	return &pb.StatusMessage{Value: "up"}, nil
}

func newStatusServer() *statusServer {
	return &statusServer{
		port: fmt.Sprintf("500%02d", rand.Intn(100)),
	}
}

type advertiseClient struct {
	client ad.AdvertiseClient
	port   string
}

func advertiseBoard(boardPort string, clientPort string) error {
	conn, err := grpc.Dial("localhost:"+boardPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer conn.Close()

	client := ad.NewAdvertiseClient(conn)
	ack, err := client.Send(context.Background(), &ad.AdvertiseMessage{Port: clientPort})
	if err != nil {
		return err
	}
	if ack.Ok != true {
		return errors.New("Board server did not acknowledge")
	}
	return err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	me := newStatusServer()

	if err := advertiseBoard("49000", me.port); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Agent serving on port:", me.port)
	c := make(chan crash)
	go me.start(c)
	for {
		select {
		case v := <-c:
			if v.Error() == "randomCrash" {
				log.Println("randomCrash in the status server.")
				return
			}
			if v.Error() == "classicCrash" {
				log.Println("classicCrash in the status server.")
				return
			}
		}
	}
}
