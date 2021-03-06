package main

import (
	"context"
	"fmt"
	"github.com/ferza17/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{}

// Unary API
func (*server) Greet(ctx context.Context, req *greetpb.GreatRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet Function was invoked with: %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}
// Server Streaming API
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreatService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " Number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		if err := stream.Send(res); err != nil {
			log.Println("Unable to send result.")
		}
		// TO se actually streaming but in real project you dont want to use it
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}
// Client Streaming API
func (*server) LongGreet(stream greetpb.GreatService_LongGreetServer) error {
	fmt.Printf("LongGreet was invoked with a streaming request %v\n", stream)
	var result string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished read the client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Unable to read client stream: %v", err)
			return err
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
	return nil
}
// Bi Directional Streaming API
func (*server) GreetEveryone(stream greetpb.GreatService_GreetEveryoneServer) error  {
	fmt.Printf("GreetEveryone was invoked with a streaming request %v\n", stream)

	for  {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + " !"
		if err := stream.Send(&greetpb.GreetEveryoneResponse{Result: result}); err != nil {
			log.Fatalf("Error while sending stream.send: %v", err)
			return err
		}
	}
}

func main() {
	fmt.Println("Server about to running...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreatServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
