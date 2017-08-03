/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"log"
	"os"

	pb "github.com/jiop/grpctests/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	defaultAge  = "9999"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	age := defaultAge
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	if len(os.Args) > 2 {
		age = os.Args[2]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name, Age: age})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	c2 := pb.NewDatabaseClient(conn)
	r2, err := c2.GetUsers(context.Background(), &pb.None{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Users: %s", r2.Users)
}
