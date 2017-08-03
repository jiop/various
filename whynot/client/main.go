package main

import (
	"context"
	"log"

	pb "github.com/jiop/various/whynot/resources"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewResourceClient(conn)
	user, err := c.GetUser(context.Background(), &pb.None{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", user)
}
