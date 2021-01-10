package main

import (
	"context"
	"fmt"
	"github.com/ferza17/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type server struct {
}

func (*server) SumData(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("SumData function was invoked with: %v\n", req)
	res := &calculatorpb.SumResponse{
		Result: req.Sum.GetSum1() + req.Sum.GetSum2(),
	}

	return res, nil
}

func (*server) SumManyTimes(req *calculatorpb.SumManyTimesRequest, stream calculatorpb.SumService_SumManyTimesServer) error {
	fmt.Printf("SumManyTimes was invoked with: %v", req)
	total := int(req.GetTotal())
	divisor := 2

	for total > 1 {
		if total%divisor == 0 {
			if err := stream.Send(&calculatorpb.SumManyTimesResponse{Result: int32(divisor)}); err != nil {
				log.Fatalf("Error when streaming: %v", err)
				return err
			}
			total /= divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v", divisor)
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func (*server) AvgLongTimes(stream calculatorpb.SumService_AvgLongTimesServer) error {
	var total int32
	var divider int

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// End Of File

			// Business Logic
			var result float64
			result = float64(total) / float64(divider)
			// End

			return stream.SendAndClose(&calculatorpb.AvgLongResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Unable to read client stream: %v", err)
			return err
		}

		// divider +1 if not in end of file
		total += req.GetNum()
		divider += 1
	}
}

func (*server) FindMaximum(stream calculatorpb.SumService_FindMaximumServer) error {
	maximum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error While Reading stream: %v", err)
			return err
		}

		// Business Logic
		number := req.GetNumber()
		if number > maximum {
			maximum = number
			if err := stream.Send(&calculatorpb.FindMaximumResponse{Maximum: maximum}); err != nil {
				log.Fatalf("Error while sending stream: %v", err)
				return err
			}
		}

	}

}

func main() {
	fmt.Println("About to start Server...")
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
