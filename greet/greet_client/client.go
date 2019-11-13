package main

import (
  "context"
  "io"
  "log"

  "github.com/_dev/grpc-go-example/greet/greetpb"

  "google.golang.org/grpc"
)

func main() {
  log.Println("Client running...")
  cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Could not connect: %v", err)
  }

  // "defer" command so that the connection closes only at the end.
  defer cc.Close()

  // Creating client...
  c := greetpb.NewGreetServiceClient(cc)
  //fmt.Println("Created client: %f", c)

  log.Println(">>")
  doUnary(c)
  log.Println("<<")

  log.Println(">>")
  doServerStreaming(c)
  log.Println("<<")
}

// Unary API
func doUnary(c greetpb.GreetServiceClient) {
  log.Println("UNARY - Starting...")

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

  log.Println("UNARY - Completed.")
}

// Server Streaming API
func doServerStreaming(c greetpb.GreetServiceClient) {
  log.Println("SERVER STREAMING - Starting...")

  req := &greetpb.GreetManyTimesRequest{
    Greeting: &greetpb.Greeting{
      FirstName: "Felipe",
      LastName:  "Sulzbach",
    },
  }
  resp, err := c.GreetManyTimes(context.Background(), req)
  if err != nil {
    log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
  }

  for { // Runs in a loop to consume the entire stream.
    msg, err := resp.Recv()
    if err == io.EOF {
      break // It has reached the end of the stream.
    }
    if err != nil {
      log.Fatalf("Error while reading stream: %v", err)
    }
    log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
  }

  log.Println("SERVER STREAMING - Completed.")
}
