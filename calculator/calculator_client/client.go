package main

import (
	"context"
	"fmt"
	"github.com/ferza17/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main()  {
	fmt.Println("About to start client...")
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to Dial: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewSumServiceClient(cc)
	//doSum(c)
	doServerStreaming(c)

}

func doSum(c calculatorpb.SumServiceClient)  {
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

func doServerStreaming(c calculatorpb.SumServiceClient)  {
	fmt.Println("About to Streaming...")
	req := &calculatorpb.SumManyTimesRequest{Total: 120}

	resStream, err := c.SumManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling SumManyTimes RPC: %v", err)
	}
	for  {
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
