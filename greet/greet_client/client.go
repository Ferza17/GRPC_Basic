package main

import (
	"context"
	"fmt"
	"github.com/ferza17/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello i'm a client")

	tls := false
	opts := grpc.WithInsecure()
	if tls {
		// grpc automatically SSL, but now use .WithInsecure because it run in localhost
		certFile := "ssl/ca.crt" // Certificate Authority Trust Certificate
		credential, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(credential)
	}

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreatServiceClient(cc)
	doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	//doBiDiStreaming(c)
	//doUnaryWithDeadline(c, 5*time.Second) // Should Complete
	//doUnaryWithDeadline(c, 1*time.Second) // Should timeout

}

func doUnary(c greetpb.GreatServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreatRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling Greet RPC: %v", err)
	}

	log.Printf("Response From Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreatServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Fery",
			LastName:  "Aditya",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling GreatManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of stream
			break
		}

		if err != nil {
			log.Fatalf("Error While reading stream: %v", err)
		}

		log.Printf("Response from greetManyTimes -> %v", msg.GetResult())

	}
}

func doClientStreaming(c greetpb.GreatServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Fery",
				LastName:  "Reza",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ainun",
				LastName:  "Dzariah",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Adelia",
				LastName:  "Putri",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Abinaya",
				LastName:  "Putra",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Unable to connect LongGreet : %v", err)
	}

	// Iterate over slice and send each message to server
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error when sending request")
		}
		// Dont do that in production / real project
		time.Sleep(1 * time.Second)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error when receiving response from LongGreet : %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", response)
}

func doBiDiStreaming(c greetpb.GreatServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creatig stream: %v", err)
		return
	}

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Fery",
				LastName:  "Reza",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ainun",
				LastName:  "Dzariah",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Adelia",
				LastName:  "Putri",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Abinaya",
				LastName:  "Putra",
			},
		},
	}

	waitc := make(chan struct{})
	// Send a bunch of messages of the client (go routine)
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending Message: %v\n", req)
			if err := stream.Send(req); err != nil {
				log.Fatalf("Error While sending request: %v", err)
				return
			}
			time.Sleep(1 * time.Second)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Error while close and send request: %v", err)
			return
		}
	}()
	// Receive a bunch of messages from the client (go routine)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// Exit function of wait channel in go
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)

	}()

	// Block until everything is done
	<-waitc
}

func doUnaryWithDeadline(c greetpb.GreatServiceClient, timeout time.Duration) {
	fmt.Println("Starting to do a doUnaryWithDeadline RPC...")

	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "John",
			LastName:  "Doe",
		},
	}
	ctx, cancelCtx := context.WithTimeout(context.Background(), timeout)
	defer cancelCtx()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {

		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded.")
			} else {
				fmt.Println("unexpected Error: ", statusErr)
			}
		} else {
			log.Fatalf("Error While calling GreetWithDeadline RPC: %v", err)
		}
		return
	}
	log.Printf("Response From GreetWithDeadline: %v", res.Result)
}
