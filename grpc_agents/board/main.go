package main

import (
	"fmt"
	"log"
	"net"
	"time"

	ad "github.com/jiop/various/grpc_agents/advertise"
	"github.com/jiop/various/grpc_agents/status"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = 49000

type agent struct {
	port int
}

type boardServer struct {
	agents []string
}

func newBoardServer() *boardServer {
	return &boardServer{}
}

func (s *boardServer) start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	ad.RegisterAdvertiseServer(grpcServer, s)

	log.Println("board server listening on port", port)
	grpcServer.Serve(lis)
}

func (s *boardServer) Send(ctx context.Context, message *ad.AdvertiseMessage) (*ad.AdvertiseAck, error) {
	log.Println("new agent listening on", message.Port)
	s.agents = append(s.agents, message.Port)
	return &ad.AdvertiseAck{Ok: true}, nil
}

func checkAgentStatus(port string) bool {
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer conn.Close()

	client := status.NewStatusClient(conn)
	st, err := client.Get(context.Background(), &status.None{})
	if err != nil || st.Value != "up" {
		return false
	}
	return true

}

func main() {
	me := newBoardServer()
	go func() {
		for {
			select {
			case <-time.After(2 * time.Second):
				log.Printf("%v", me.agents)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				var newAgents []string
				for _, a := range me.agents {
					if checkAgentStatus(a) {
						newAgents = append(newAgents, a)
					}
				}
				me.agents = newAgents
			}
		}
	}()

	me.start()
}
