package main

import (
	"context"
	"fmt"
	"log"

	"github.com/_dev/grpc-go-example/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client running...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	// "defer" command so that the connection closes only at the end.
	defer cc.Close()

	// Creating client...
	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Println("Created client: %f", c)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Felipe",
			LastName:  "Sulzbach",
		},
	}
	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", resp.Result)
}
