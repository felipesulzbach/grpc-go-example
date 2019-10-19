package main

import (
	"fmt"
	"log"
	"net"

	"github.com/_dev/grpc-go-example/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello World!")

	// Creating the port of GRPC server...
	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Creating GRPC server...
	s := grpc.NewServer()

	// Registring de GreetService in GRPC server...
	greetpb.RegisterGreetServiceServer(s, &server{})

	// Binding the port to GRPC server...
	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
