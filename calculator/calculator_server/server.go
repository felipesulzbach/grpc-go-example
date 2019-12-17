package main

import (
  "context"
  "fmt"
  "io"
  "log"
  "math"
  "net"

  "github.com/_dev/grpc-go-example/calculator/calculatorpb"

  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
  fmt.Printf("Received Sum RPC: %v\n", req)
  firstNumber := req.FirstNumber
  secondNumber := req.SecondNumber
  sum := firstNumber + secondNumber
  res := &calculatorpb.SumResponse{
    SumResult: sum,
  }
  return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
  log.Printf("Received PrimeNumberDecomposition RPC: %v\n", req)
  number := req.GetNumber()
  divisor := int64(2)

  for number > 1 {
    if number%divisor == 0 {
      stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
        PrimeFactor: divisor,
      })
      number = number / divisor
    } else {
      divisor++
      log.Printf("Divisor has increased to %v\n", divisor)
    }
  }
  return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
  log.Println("Received ComputeAverage RPC.")

  sum := int32(0)
  count := 0
  for {
    req, err := stream.Recv()
    if err == io.EOF {
      average := float64(sum) / float64(count)
      return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
        Average: average,
      })
    }
    if err != nil {
      log.Fatalf("Error while reading client stream: %v", err)
    }
    sum += req.GetNumber()
    count++
  }
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
  log.Println("Received FindMaximum RPC.")

  maximum := int32(0)
  for {
    req, err := stream.Recv()
    if err == io.EOF {
      return nil
    }
    if err != nil {
      log.Fatalf("Error while reading client stream: %v", err)
      return err
    }
    number := req.GetNumber()
    if number > maximum {
      maximum = number
      err = stream.Send(&calculatorpb.FindMaximumResponse{
        Maximum: maximum,
      })
      if err != nil {
        log.Fatalf("Error while sending client stream: %v", err)
        return err
      }
    }
  }
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
  log.Println("Received SquareRoot RPC.")

  number := req.GetNumber()
  if number < 0 {
    return nil, status.Errorf(
      codes.InvalidArgument,
      fmt.Sprintf("Received a negative number: %v!", number),
    )
  }
  return &calculatorpb.SquareRootResponse{
    NumberRoot: math.Sqrt(float64(number)),
  }, nil
}

func main() {
  log.Println("SERVER - Starting...")

  // Creating the port of GRPC server...
  list, err := net.Listen("tcp", "0.0.0.0:50051")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  // Creating GRPC server...
  s := grpc.NewServer()

  // Registring de CalculatorService in GRPC server...
  calculatorpb.RegisterCalculatorServiceServer(s, &server{})

  log.Println("SERVER - Running...")

  // Binding the port to GRPC server...
  if err := s.Serve(list); err != nil {
    log.Fatalf("Failed to serve: %v", err)
  }
}
