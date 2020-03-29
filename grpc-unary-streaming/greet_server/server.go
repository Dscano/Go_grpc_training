package main

import (
	"fmt" 
	"net"
	"log"
	"context"
	"google.golang.org/grpc"
	"github.com/grpc-unary-streaming/greetpb_unary"
)

//this is the server "object"
type server struct{}

func(*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Println("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
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
	greetpb.RegisterGreetServiceServer(s , &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}