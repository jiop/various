package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/jiop/various/grpc_agents_2/mesh"
)

var (
	master = flag.String("master", "", "existing node")
	port   = flag.Int("port", 49000, "port")
	host   = flag.String("host", "localhost", "host")
)

type Node struct {
	host  string
	port  int
	peers map[string]bool
}

func newNode(host string, port int) *Node {
	return &Node{
		host:  host,
		port:  port,
		peers: make(map[string]bool, 0),
	}
}

func (n *Node) GetNodes(ctx context.Context, message *pb.StatusMessage) (*pb.NodesMessage, error) {
	log.Printf("Answer to GetNodes call from %s", message.Address)
	nodes := make([]*pb.StatusMessage, 0)
	for k := range n.peers {
		nodes = append(nodes, &pb.StatusMessage{Address: k})
	}
	n.peers[message.Address] = true
	return &pb.NodesMessage{Nodes: nodes}, nil
}

func (n *Node) me() string {
	return fmt.Sprintf("%s:%d", n.host, n.port)
}

func (n *Node) getNodes() error {
	for p := range n.peers {
		if n.peers[p] == false || p == n.me() {
			continue
		}
		conn, err := grpc.Dial(p, grpc.WithInsecure())
		if err != nil {
			n.peers[p] = false
			log.Printf("%v", err)
			continue
		}
		defer conn.Close()

		client := pb.NewMeshInfoClient(conn)
		nMess, err := client.GetNodes(context.Background(), &pb.StatusMessage{Address: n.me()})
		if err != nil {
			n.peers[p] = false
			log.Printf("%v", err)
			continue
		}

		for _, newNode := range nMess.Nodes {
			n.peers[newNode.Address] = true
		}
	}
	return nil
}

func (n *Node) startServer(c chan error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", n.port))
	if err != nil {
		c <- err
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	pb.RegisterMeshInfoServer(server, n)
	if err = server.Serve(lis); err != nil {
		c <- err
	}
}

func (n *Node) updatePeersList(c chan error) {
	for {
		select {
		case <-time.After(1 * time.Second):
			if err := n.getNodes(); err != nil {
				c <- err
			}
		}
	}
}

func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	me := newNode("localhost", *port)
	if *master != "" {
		me.peers[*master] = true
	}

	fmt.Println("Agent serving on:", me.me())

	channel := make(chan error)
	go me.startServer(channel)
	go me.updatePeersList(channel)

	go func() {
		for {
			select {
			case <-time.After(2 * time.Second):
				log.Printf("%v", me.peers)
			}
		}
	}()

	for {
		select {
		case err := <-channel:
			log.Fatal(err)
		}
	}
}
