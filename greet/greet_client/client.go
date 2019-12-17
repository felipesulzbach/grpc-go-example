package main

import (
  "context"
  "fmt"
  "io"
  "log"
  "time"

  "github.com/_dev/grpc-go-example/greet/greetpb"

  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
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

  log.Println(">>")
  doBidirectionalStreaming(c)
  log.Println("<<")

  log.Println(">>")
  doUnaryWithDeadline(c, 5*time.Second) // should complete
  doUnaryWithDeadline(c, 1*time.Second) // should timeout
  log.Println("<<")

  log.Println("Client stoped.")
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
      break
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
  log.Printf("LongGreet Response: %v", resp)

  log.Println("CLIENT STREAMING - Completed.")
}

// Bidirectional Streaming API
func doBidirectionalStreaming(c greetpb.GreetServiceClient) {
  log.Println("BIDIRECTIONAL STREAMING - Starting...")

  // We create a stream by invoking the client.
  stream, err := c.GreetEveryone(context.Background())
  if err != nil {
    log.Fatalf("Error while creating stream: %v", err)
    return
  }

  requests := []*greetpb.GreetEveryoneRequest{
    &greetpb.GreetEveryoneRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Felipe 01",
      },
    },
    &greetpb.GreetEveryoneRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Felipe 02",
      },
    },
    &greetpb.GreetEveryoneRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Felipe 03",
      },
    },
    &greetpb.GreetEveryoneRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Felipe 04",
      },
    },
    &greetpb.GreetEveryoneRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Felipe 05",
      },
    },
  }

  waitc := make(chan struct{}) // Wait for the channel to close.

  // We send a bunch of messages to the client (go routine).
  go func() {
    // function to send a bunch of messages
    for _, req := range requests {
      log.Printf("GreetEveryone sent: %v...\n", req)
      stream.Send(req)
      time.Sleep(1000 * time.Millisecond)
    }
    stream.CloseSend()
  }()

  // We receive a bunch of messages from the client (go routine).
  go func() {
    // function to receive a banch of messages
    for {
      res, err := stream.Recv()
      if err == io.EOF {
        break // It has reached the end of the stream.
      }
      if err != nil {
        log.Fatalf("Error while receiving stream: %v", err)
        break // It has reached the end of the stream.
      }
      log.Printf("GreetEveryone received: %v...\n", res)
    }
    close(waitc)
  }()

  // Block until everything is done.
  go func() {
  }()

  <-waitc

  log.Println("BIDIRECTIONAL STREAMING - Completed.")
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
  fmt.Println("Starting to do a UnaryWithDeadline RPC...")
  req := &greetpb.GreetWithDeadlineRequest{
    Greeting: &greetpb.Greeting{
      FirstName: "Stephane",
      LastName:  "Maarek",
    },
  }
  ctx, cancel := context.WithTimeout(context.Background(), timeout)
  defer cancel()

  res, err := c.GreetWithDeadline(ctx, req)
  if err != nil {

    statusErr, ok := status.FromError(err)
    if ok {
      if statusErr.Code() == codes.DeadlineExceeded {
        fmt.Println("Timeout was hit! Deadline was exceeded")
      } else {
        fmt.Printf("unexpected error: %v", statusErr)
      }
    } else {
      log.Fatalf("error while calling GreetWithDeadline RPC: %v", err)
    }
    return
  }
  log.Printf("Response from GreetWithDeadline: %v", res.Result)
}
