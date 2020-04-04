package main

import (
	"fmt" 
	"net"
	"log"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	certFile := "/home/diaccio/go/src/github.com/grpc-SSL/ssl/server.crt"
	keyFile := "/home/diaccio/go/src/github.com/grpc-SSL/ssl/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil{
		log.Fatalf("Failed loading certificates: %v", err)
		return
	}

	opts := grpc.Creds(creds)
	//grpc server
	s := grpc.NewServer(opts)
	// we pass the server as a struct
	greetpb.RegisterGreetServiceServer(s , &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}