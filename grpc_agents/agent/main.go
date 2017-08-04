package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/context"

	co "github.com/jiop/various/grpc_agents/connection"
	"google.golang.org/grpc"
)

type statusServer struct {
	port string
}

type crash interface {
	Error() string
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
	co.RegisterStatusServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		c <- &classicCrash{}
	}
}

func (s *statusServer) Get(ctx context.Context, in *co.None) (*co.StatusMessage, error) {
	return &co.StatusMessage{Value: "up"}, nil
}

func newStatusServer() *statusServer {
	return &statusServer{
		port: fmt.Sprintf("500%02d", rand.Intn(100)),
	}
}

func (s *statusServer) advertiseBoard(boardPort string) error {
	conn, err := grpc.Dial("localhost:"+boardPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer conn.Close()

	client := co.NewPingClient(conn)
	_, err = client.Send(context.Background(), &co.PingMessage{Port: s.port})
	return err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	me := newStatusServer()

	if err := me.advertiseBoard("49000"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Agent serving on port:", me.port)
	c := make(chan crash)
	me.start(c)
}
