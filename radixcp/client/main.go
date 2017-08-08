package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jiop/various/radixcp/request"
	"google.golang.org/grpc"
)

const serverPort = "49123"

func main() {
	conn, err := grpc.Dial("localhost:"+serverPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRequestClient(conn)
	cpMessage, err := client.GetCP(context.Background(), &pb.CityMessage{Name: "colombes"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cpMessage.Postalcode)
}
