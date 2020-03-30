package main

import (
	"fmt" 
	"net"
	"log"
	"math"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"github.com/errors/square"
	"google.golang.org/grpc/codes"
)

//this is the server "object"
type server struct{}

func(*server) SquareRoot( ctx context.Context, req *squarepb.SquareRootRequest) (*squarepb.SquareRootResponse,error){
	fmt.Printf("Received SquareRoot RPC\n")
	number := req.GetNumber() 

	if number<0 {
		return nil, status.Errorf(
		codes.InvalidArgument,
		fmt.Sprintf("Recived a negative number: ", number),
		)
	}

	return &squarepb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil

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
	squarepb.RegisterSquareRootServiceServer(s , &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
