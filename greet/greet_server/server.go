package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/_dev/grpc-go-example/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	resp := &greetpb.GreetResponse{
		Result: result,
	}
	return resp, nil
}

func main() {
	fmt.Println("Server starting...")

	// Creating the port of GRPC server...
	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Creating GRPC server...
	s := grpc.NewServer()

	// Registring de GreetService in GRPC server...
	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Println("Server running...")

	// Binding the port to GRPC server...
	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
