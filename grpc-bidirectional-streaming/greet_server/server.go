package main

import (
	"fmt" 
	"net"
	"log"
	"io"
	"google.golang.org/grpc"
	"github.com/grpc-bidirectional-streaming/greetpb_bidirectional"
)

//this is the server "object"
type server struct{}

func(*server) GreetEveryone( stream greetpb.GreetService_GreetEveryoneServer) error{
	fmt.Printf("GreetEveryone function was invoked with a streaming request\n")

	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result   = "Hello " + firstName + "! "
		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil{
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
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
	greetpb.RegisterGreetServiceServer(s , &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}