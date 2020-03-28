package main

import (
	"fmt" 
	"net"
	"log"
	"context"
	"google.golang.org/grpc"
	"github.com/exercise_1/sumpb_unary"
)

//this is the server "object"
type server struct{}

func(*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error){
	fmt.Println("Greet function was invoked with %v\n", req)
	var result int32
	for _, number := range req.GetSumming().GetNumber(){
		result+=number
	}
	res := &sumpb.SumResponse{
		Result: result,

	}
	return res, nil
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
	sumpb.RegisterSumServiceServer(s , &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}