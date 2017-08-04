package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	pb "github.com/jiop/various/grpcstream/resources"
)

const (
	port = 50051
	base = 100
)

type server struct {
	savedDataLength int
	savedData       []*pb.Data
}

func (*server) GetData(ctx context.Context, none *pb.None) (*pb.Data, error) {
	d := time.Duration(rand.Intn(base)) * time.Millisecond
	select {
	case <-time.After(d):
		return &pb.Data{Name: "test"}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *server) StreamListData(none *pb.None, stream pb.Resources_StreamListDataServer) error {
	ctx := stream.Context()
	for _, v := range s.savedData {
		log.Printf("current data: %v", v)
		d := time.Duration(10 * time.Millisecond)
		select {
		case <-time.After(d):
			err := stream.Send(v)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return io.EOF
}

func (s *server) StreamRandData(none *pb.None, stream pb.Resources_StreamRandDataServer) error {
	ctx := stream.Context()
	for {
		d := time.Duration(10 * time.Millisecond)
		select {
		case <-time.After(d):
			v := rand.Intn(100)
			err := stream.Send(&pb.Data{Name: fmt.Sprintf("%d", v)})
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func createServer() *server {
	res := make([]*pb.Data, 0)
	length := rand.Intn(1000)
	for i := 0; i < length; i++ {
		res = append(res, &pb.Data{Name: fmt.Sprintf("rand: %d", rand.Intn(1000))})
	}
	return &server{savedDataLength: length, savedData: res}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterResourcesServer(grpcServer, createServer())
	grpcServer.Serve(lis)
}
