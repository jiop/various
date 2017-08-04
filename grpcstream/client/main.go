package main

import (
	"flag"
	"log"
	"time"

	pb "github.com/jiop/various/grpcstream/resources"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	mode = flag.String("mode", "one", "one, list_stream or rand_stream")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer conn.Close()

	client := pb.NewResourcesClient(conn)

	switch *mode {
	case "one":
		getData(client)
	case "list_stream":
		streamListData(client)
	case "rand_stream":
		streamRandData(client)
	}
}

// Ask one data to the server
func getData(client pb.ResourcesClient) {
	log.Println("-----------------GetData-----------------")

	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Millisecond)
	defer cancel()

	r, err := client.GetData(ctx, &pb.None{})
	if err != nil {
		log.Printf("err: %v", err)
	}
	log.Printf("Random number from server: %v", r)
}

// Ask data list to the server for the maximum duration of the context
func streamListData(client pb.ResourcesClient) {
	log.Println("-----------------ListData-----------------")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stream, err := client.StreamListData(ctx, &pb.None{})
	if err != nil {
		log.Printf("err: %v", err)
		return
	}
	for {
		res, err := stream.Recv()
		if grpc.ErrorDesc(err) == "EOF" {
			log.Println("Server sent all his random number.")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println("random number from server: ", res)
	}
}

// Ask random data to the server for the duration of the context
func streamRandData(client pb.ResourcesClient) {
	log.Println("-----------------RandData-----------------")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stream, err := client.StreamRandData(ctx, &pb.None{})
	if err != nil {
		log.Printf("err: %v", err)
		return
	}

	c := make(chan *pb.Data)
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				return
			}
			c <- res
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("Client reach timeout.")
			close(c)
			return
		case res := <-c:
			log.Println("Random number from server:", res)
		}
	}
}
