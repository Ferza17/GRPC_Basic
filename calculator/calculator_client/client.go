package main

import (
	"context"
	"fmt"
	"github.com/ferza17/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("About to start client...")
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to Dial: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewSumServiceClient(cc)
	//doSum(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)

}

func doSum(c calculatorpb.SumServiceClient) {
	fmt.Println("About to doSum...")
	req := &calculatorpb.SumRequest{
		Sum: &calculatorpb.Sum{
			Sum1: 3,
			Sum2: 10,
		},
	}

	res, err := c.SumData(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when SumData: %v", err)
	}
	log.Printf("Result Data -> %v	", res)
}

func doServerStreaming(c calculatorpb.SumServiceClient) {
	fmt.Println("About to Streaming...")
	req := &calculatorpb.SumManyTimesRequest{Total: 120}

	resStream, err := c.SumManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling SumManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// Reached the end of stream
			break
		}

		if err != nil {
			log.Fatalf("Error While Reading Stream: %v", err)
		}

		log.Printf("Response from SumManyTimes -> %v", msg)
	}
}

func doClientStreaming(c calculatorpb.SumServiceClient) {
	fmt.Println("about to start Client Streaming RPC...")

	request := []*calculatorpb.AvgLongRequest{
		{Num: 1},
		{Num: 2},
		{Num: 3},
		{Num: 4},
	}

	stream, err := c.AvgLongTimes(context.Background())
	if err != nil {
		log.Fatalf("Unable to call AvgLongTimes")
	}

	for _, req := range request {
		fmt.Println("Sending Request : ", req)

		if err := stream.Send(req); err != nil {
			log.Fatalln("Unable to send request : ", req)
		}

		// Dont do this in real project / production.
		time.Sleep(1 * time.Second)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Unable to Close and Receive response : ", err)
	}
	fmt.Println("AvgLongTimes Response : ", response)

}

func doBiDiStreaming(c calculatorpb.SumServiceClient) {
	fmt.Println("about to start BiDi Streaming RPC...")

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while FindMaximum: %v", err)
	}

	waitc := make(chan struct{})

	// send go routine
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, num := range numbers {
			fmt.Println("Sending number: ", num)
			if err := stream.Send(&calculatorpb.FindMaximumRequest{Number: num}); err != nil {
				break
			}
			// dont do this in production; just to make sure its work
			time.Sleep(1 * time.Second)
		}

		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Error while close and send request: %v", err)
			return
		}
	}()

	// receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receive stream: %v", err)
			}

			maximum := res.GetMaximum()
			fmt.Printf("Receive a new Maximum of...: %v\n", maximum)
		}
		close(waitc)
	}()

	<-waitc

}
