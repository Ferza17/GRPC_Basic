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

func main()  {
	fmt.Println("About to start client...")
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to Dial: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewSumServiceClient(cc)
	//doSum(c)
	//doServerStreaming(c)
	doClientStreaming(c)

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

func doClientStreaming(c calculatorpb.SumServiceClient)  {
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

	for _, req := range request{
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
