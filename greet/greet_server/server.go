package main

import (
  "bytes"
  "context"
  "log"
  "net"
  "strconv"
  "time"

  "github.com/_dev/grpc-go-example/greet/greetpb"

  "google.golang.org/grpc"
)

type server struct{}

// Unary API
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
  log.Printf("Greet function was invoked with %v", req)

  firstName := req.GetGreeting().GetFirstName()
  var buffer bytes.Buffer
  buffer.WriteString("Hello ")
  buffer.WriteString(firstName)

  response := &greetpb.GreetResponse{
    Result: buffer.String(),
  }
  return response, nil
}

// Server Streaming API
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
  log.Printf("GreetManyTimes function was invoked with %v", req)

  firstName := req.GetGreeting().GetFirstName()
  for i := 0; i < 10; i++ {
    var buffer bytes.Buffer
    buffer.WriteString("Hello ")
    buffer.WriteString(firstName)
    buffer.WriteString(" number ")
    buffer.WriteString(strconv.Itoa(i))

    response := &greetpb.GreetManyTimesResponse{
      Result: buffer.String(),
    }
    stream.Send(response)
    time.Sleep(1000 * time.Millisecond) // wating 1 second.
  }
  return nil
}

func main() {
  log.Println("SERVER - starting...")

  // Creating the port of GRPC server...
  list, err := net.Listen("tcp", "0.0.0.0:50051")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  // Creating GRPC server...
  s := grpc.NewServer()

  // Registring de GreetService in GRPC server...
  greetpb.RegisterGreetServiceServer(s, &server{})

  log.Println("SERVER - running...")

  // Binding the port to GRPC server...
  if err := s.Serve(list); err != nil {
    log.Fatalf("Failed to serve: %v", err)
  }
}
