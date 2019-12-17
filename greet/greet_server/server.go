package main

import (
  "bytes"
  "context"
  "fmt"
  "io"
  "log"
  "net"
  "strconv"
  "time"

  "github.com/_dev/grpc-go-example/greet/greetpb"

  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
)

type server struct{}

// Unary API
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
  log.Printf("Greet - function was invoked with %v.", req)

  firstName := req.GetGreeting().GetFirstName()
  var result bytes.Buffer
  result.WriteString("Hello ")
  result.WriteString(firstName)
  result.WriteString("!")

  response := &greetpb.GreetResponse{
    Result: result.String(),
  }

  log.Println("Greet - returned.")
  return response, nil
}

// Server Streaming API
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
  log.Printf("GreetManyTimes - function was invoked with %v.", req)

  firstName := req.GetGreeting().GetFirstName()
  for i := 0; i < 10; i++ {
    var result bytes.Buffer
    result.WriteString("Hello ")
    result.WriteString(firstName)
    result.WriteString(" number ")
    result.WriteString(strconv.Itoa(i))
    result.WriteString("!")

    response := &greetpb.GreetManyTimesResponse{
      Result: result.String(),
    }
    stream.Send(response)
    time.Sleep(1000 * time.Millisecond) // wating 1 second.
  }

  log.Println("GreetManyTimes - returned.")
  return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
  log.Println("LongGreet - function was invoked with streaming.")

  var result bytes.Buffer
  for { // Runs in a loop to consume the entire stream.
    request, err := stream.Recv()
    if err == io.EOF {
      log.Println("LongGreet - returned.")
      return stream.SendAndClose(&greetpb.LongGreetResponse{
        Result: result.String(),
      })
    }
    if err != nil {
      log.Fatalf("Error while reading stream: %v", err)
    }

    firstName := request.GetGreeting().GetFirstName()
    result.WriteString("Hello ")
    result.WriteString(firstName)
    result.WriteString("! ")
  }
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
  log.Println("GreetEveryone - function was invoked with streaming.")

  var result bytes.Buffer
  for { // Runs in a loop to consume the entire stream.
    req, err := stream.Recv()
    if err == io.EOF {
      return nil
    }
    if err != nil {
      log.Fatalf("Error while reading stream: %v", err)
      return err
    }

    firstName := req.GetGreeting().GetFirstName()
    result.WriteString("Hello ")
    result.WriteString(firstName)
    result.WriteString("! ")

    err = stream.Send(&greetpb.GreetEveryoneResponse{
      Result: result.String(),
    })
    if err != nil {
      log.Fatalf("Error while sending stream: %v", err)
      return err
    }
  }
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
  fmt.Printf("GreetWithDeadline function was invoked with %v\n", req)
  for i := 0; i < 3; i++ {
    if ctx.Err() == context.Canceled {
      // the client canceled the request
      fmt.Println("The client canceled the request!")
      return nil, status.Error(codes.Canceled, "the client canceled the request")
    }
    time.Sleep(1 * time.Second)
  }
  firstName := req.GetGreeting().GetFirstName()
  result := "Hello " + firstName
  res := &greetpb.GreetWithDeadlineResponse{
    Result: result,
  }
  return res, nil
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

  // Registring de GreetService in GRPC server...
  greetpb.RegisterGreetServiceServer(s, &server{})

  log.Println("SERVER - Running...")

  // Bidirectional the port to GRPC server...
  if err := s.Serve(list); err != nil {
    log.Fatalf("Failed to serve: %v", err)
  }
}
