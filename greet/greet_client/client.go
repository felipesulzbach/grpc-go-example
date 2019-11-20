package main

import (
  "context"
  "io"
  "log"
  "time"

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

  log.Println(">>")
  doClientStreaming(c)
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
  log.Printf("Greet Request: %v...", req)
  resp, err := c.Greet(context.Background(), req)
  if err != nil {
    log.Fatalf("Error while calling Greet RPC: %v", err)
  }
  log.Printf("Greet Response: %v", resp.Result)

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
  log.Printf("GreetManyTimes Request: %v...", req)
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
    log.Printf("GreetManyTimes Response: %v", msg.GetResult())
  }

  log.Println("SERVER STREAMING - Completed.")
}

// Client Streaming API
func doClientStreaming(c greetpb.GreetServiceClient) {
  log.Println("CLIENT STREAMING - Starting...")

  requests := []*greetpb.LongGreetRequest{
    &greetpb.LongGreetRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Felipe 01",
      },
    },
    &greetpb.LongGreetRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Felipe 02",
      },
    },
    &greetpb.LongGreetRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Felipe 03",
      },
    },
    &greetpb.LongGreetRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Felipe 04",
      },
    },
    &greetpb.LongGreetRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Felipe 05",
      },
    },
  }

  stream, err := c.LongGreet(context.Background())
  if err != nil {
    log.Fatalf("Error while reading stream: %v", err)
  }

  // We iterate over our slice and send each message individually.
  for _, req := range requests {
    log.Printf("LongGreet Request: %v...", req)
    stream.Send(req)
    time.Sleep(1000 * time.Millisecond)
  }

  resp, err := stream.CloseAndRecv()
  if err != nil {
    log.Fatalf("Error while calling LongGreet RPC: %v", err)
  }
  log.Fatalf("LongGreet Response: %v", resp)

  log.Println("CLIENT STREAMING - Completed.")
}
