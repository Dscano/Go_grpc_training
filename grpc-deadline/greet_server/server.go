package main

import (
	"fmt" 
	"net"
	"log"
	"context"
	"time"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc"
	"github.com/grpc-deadline/greetpb_deadline"
)

//this is the server "object"
type server struct{}

func(*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Println("Greet with Deadline function was invoked with %v\n", req)
	for i := 0; i < 3 ; i++ {
		if ctx.Err() == context.Canceled {
			// the client cancelled the request
			fmt.Println("The client canceled the request")
			return  nil , status.Error(codes.Canceled,"the client canceled the request")
		}
		time.Sleep(1*time.Second)
	}

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