package main

import (
	"fmt"
	"log"
	"net"
	"time"

	co "github.com/jiop/various/grpc_agents/connection"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = 49000

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
	co.RegisterPingServer(grpcServer, s)

	log.Println("board server listening on port", port)
	grpcServer.Serve(lis)
}

func (s *boardServer) Send(ctx context.Context, message *co.PingMessage) (*co.None, error) {
	log.Println("new agent listening on", message.Port)
	s.agents = append(s.agents, message.Port)
	return &co.None{}, nil
}

func (s *boardServer) logAgents(t time.Duration) {
	for {
		select {
		case <-time.After(t):
			log.Printf("%v", s.agents)
		}
	}
}

func (s *boardServer) gcAgents(t time.Duration, checker func(string) bool) {
	for {
		select {
		case <-time.After(t):
			var newAgents []string
			for _, a := range s.agents {
				if checker(a) {
					newAgents = append(newAgents, a)
				}
			}
			s.agents = newAgents
		}
	}
}

func checkAgentStatus(port string) bool {
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer conn.Close()

	client := co.NewStatusClient(conn)
	st, err := client.Get(context.Background(), &co.None{})
	return err == nil && st.Value == "up"
}

func main() {
	me := newBoardServer()

	go me.logAgents(1 * time.Second)
	go me.gcAgents(1*time.Second, checkAgentStatus)

	me.start()
}
