package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/jiop/various/whynot/resources"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
	base = 100
)

type server struct{}

func concurrentFunc() string {
	c := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(base)) * time.Millisecond)
		c <- "Emma"
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(base)) * time.Millisecond)
		c <- "Albert"
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(base)) * time.Millisecond)
		c <- "James"
	}()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case <-timeout:
		case result := <-c:
			return result
		}
	}
	return "None"
}

func (s *server) GetUser(ctx context.Context, in *pb.None) (*pb.User, error) {
	return &pb.User{
		Name:      concurrentFunc(),
		Age:       "Age",
		Firstname: "Firstname",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterResourceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
