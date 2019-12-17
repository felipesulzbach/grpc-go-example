package main

import (
  "context"
  "fmt"
  "io"
  "log"
  "time"

  "github.com/grpc-go-course/calculator/calculatorpb"

  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
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

  log.Println(">>")
  doClientStreaming(c)
  log.Println("<<")

  log.Println(">>")
  doBidirectionalStreaming(c)
  log.Println("<<")

  log.Println(">>")
  doSquareRoot(c)
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

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
  log.Println("CLIENT STREAMING - Starting...")

  stream, err := c.ComputeAverage(context.Background())
  if err != nil {
    log.Fatalf("Error while open stream: %v", err)
  }

  numbers := []int32{3, 5, 9, 54, 23}
  for _, number := range numbers {
    log.Printf("Sending number: %v\n", number)
    stream.Send(&calculatorpb.ComputeAverageRequest{
      Number: number,
    })
  }

  res, err := stream.CloseAndRecv()
  if err != nil {
    log.Fatalf("Error while receiving response: %v", err)
  }
  log.Printf("Response from Average: %v", res.GetAverage())

  log.Println("CLIENT STREAMING - Completed.")
}

func doBidirectionalStreaming(c calculatorpb.CalculatorServiceClient) {
  log.Println("BIDIRECTIONAL STREAMING - Starting...")

  stream, err := c.FindMaximum(context.Background())
  if err != nil {
    log.Fatalf("Error while open stream: %v", err)
  }

  waitc := make(chan struct{})

  // send go routine
  go func() {
    numbers := []int32{4, 7, 2, 19, 4, 6, 32}
    for _, number := range numbers {
      log.Printf("Sending number: %v\n", number)
      stream.Send(&calculatorpb.FindMaximumRequest{
        Number: number,
      })
      time.Sleep(1000 * time.Millisecond)
    }
    stream.CloseSend()
  }()

  // receive go routing
  go func() {
    for {
      res, err := stream.Recv()
      if err == io.EOF {
        break // It has reached the end of the stream.
      }
      if err != nil {
        log.Fatalf("Error while receiving stream: %v", err)
        break // It has reached the end of the stream.
      }
      maximum := res.GetMaximum()
      log.Printf("BidirectionalStreaming received: %v...\n", maximum)
    }
    close(waitc)
  }()

  <-waitc

  log.Println("BIDIRECTIONAL STREAMING - Completed.")
}

func doSquareRoot(c calculatorpb.CalculatorServiceClient) {
  log.Println("SQUAREROOT STREAMING - Starting...")

  // correct call
  doErrorCall(c, 10)

  // error call
  doErrorCall(c, -2)

  log.Println("SQUAREROOT STREAMING - Completed.")
}

func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
  res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: n})
  if err != nil {
    err, ok := status.FromError(err)
    if ok {
      log.Printf("Error message from server: %v - %v.\n", err.Code(), err.Message())
      if err.Code() == codes.InvalidArgument {
        log.Println("We probaly sent a negative number!")
        return
      }
    } else {
      log.Fatalf("Error while calling SquareRoot: %v", err)
      return
    }
  }
  log.Printf("Result of square root of %v: %v\n", n, res.GetNumberRoot())
}
