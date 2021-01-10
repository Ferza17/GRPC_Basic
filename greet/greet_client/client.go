package main

import (
	"context"
	"fmt"
	"github.com/ferza17/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello i'm a client")
	// grpc automatically SSL, but now use .WithInsecure because it run in localhost
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreatServiceClient(cc)
	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)

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
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)

	}()

	// Block until everything is done
	<-waitc
}
