package main

import (
	"fmt" 
	"net"
	"log"
	"context"
	"io"
	"google.golang.org/grpc"
	"github.com/exercise_3/calc"
)

//this is the server "object"
type server struct{}

func(*server) ComputeAverage( stream calcpb.CalcService_ComputeAverageServer) error{
	fmt.Printf("ComputeAverage function was invoked with a streaming request\n")
	sum := int32(0)
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum)/float64(count)
			// we have finished reading the client stream
			return stream.SendAndClose(&calcpb.ComputeAverageResponse{
				Number: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		sum += req.GetNumber()
		count++
	}

}


func main(){
	fmt.Println("Starting Server")
	// conncetion creation
	lis,err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil{
		log.Fatalf("Failed to listen: %v", err)
	}

	//grpc server
	s := grpc.NewServer()
	// we pass the server as a struct
	calcpb.RegisterCalcServiceServer(s , &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
