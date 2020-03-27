package main

import (
	"fmt" 
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/greet/greetpb"
)

//this is the server "object"
type server struct{}

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