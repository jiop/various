package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	pb "github.com/jiop/various/radixcp/request"
	"google.golang.org/grpc"
)

var (
	city       = flag.String("city", "Paris 01", "a city name")
	serverHost = flag.String("server", "localhost:49123", "gRPC server host")
)

func transformStr(s string) string {
	return strings.ToLower(strings.Replace(s, " ", "_", -1))
}

func main() {
	flag.Parse()

	if *city == "" || *serverHost == "" {
		log.Fatal("missing required argument.")
	}

	conn, err := grpc.Dial(*serverHost, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRequestClient(conn)
	cpMessage, err := client.GetCP(context.Background(), &pb.CityMessage{Name: transformStr(*city)})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cpMessage.Postalcode)
}
