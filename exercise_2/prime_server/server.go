package main

import (
	"fmt" 
	"net"
	"log"
	"context"
	"google.golang.org/grpc"
	"github.com/exercise_2/prime"
)

//this is the server "object"
type server struct{}

func(*server) Prime(ctx context.Context, req *primepb.PrimeRequest) (*primepb.PrimeResponse, error){
	fmt.Println("Prime function was invoked with %v\n", req)
	number := req.GetPriming().GetNumber()
	res := &primepb.PrimeResponse{
		Result: number,

	}
	return res, nil
}


func (*server)  PrimeStreaming(req *primepb.PrimeStreamingRequest, stream primepb.PrimeService_PrimeStreamingServer) error {
	fmt.Println("Prime function was invoked with %v\n", req)
	number := req.GetPriming().GetNumber()
    k := int32(2)
	for number > 1{
		if number % k == 0 {
			stream.Send(&primepb.PrimeStreamingResponse{
				Result: k,
		})
			number = number/k
		} else{
			k++
			fmt.Println("Divisor has increased to %v",k)
		}
		
	}

	return nil
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
	primepb.RegisterPrimeServiceServer(s , &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}