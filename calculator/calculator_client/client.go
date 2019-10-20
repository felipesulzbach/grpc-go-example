package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-go-course/calculator/calculatorpb"

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
	c := calculatorpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  6,
		SecondNumber: 60,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}
