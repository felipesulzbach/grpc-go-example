package main

import (
  "context"
  "fmt"
  "io"
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

  log.Println(">>")
  doUnary(c)
  log.Println("<<")

  log.Println(">>")
  doServerStreaming(c)
  log.Println("<<")
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
  log.Println("UNARY - Starting...")

  req := &calculatorpb.SumRequest{
    FirstNumber:  6,
    SecondNumber: 60,
  }
  res, err := c.Sum(context.Background(), req)
  if err != nil {
    log.Fatalf("Error while calling Sum RPC: %v", err)
  }
  log.Printf("Response from Sum: %v", res.SumResult)

  log.Println("UNARY - Completed.")
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
  log.Println("SERVER STREAMING - Starting...")

  req := &calculatorpb.PrimeNumberDecompositionRequest{
    Number: 125465465,
  }
  res, err := c.PrimeNumberDecomposition(context.Background(), req)
  if err != nil {
    log.Fatalf("Error while calling Sum RPC: %v", err)
  }
  for { // Runs in a loop to consume the entire stream.
    msg, err := res.Recv()
    if err == io.EOF {
      break // It has reached the end of the stream.
    }
    if err != nil {
      log.Fatalf("Error while reading stream: %v", err)
    }
    log.Printf("Response from PrimeNumber: %v", msg.GetPrimeFactor())
  }

  log.Println("SERVER STREAMING - Completed.")
}
