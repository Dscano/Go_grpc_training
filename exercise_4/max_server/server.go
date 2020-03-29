package main

import (
	"fmt" 
	"net"
	"log"
	"io"
	"google.golang.org/grpc"
	"github.com/exercise_4/max"
)

//this is the server "object"
type server struct{}

func(*server) FindMax( stream maxpb.MaxService_FindMaxServer) error{
	fmt.Printf("Max function was invoked with a streaming request\n")
	max := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		number := req.GetNumber()
		if number > max {
				max = number
			    sendErr := stream.Send(&maxpb.MaxResponse{
				Result: max,
				})
			if sendErr !=nil {
				log.Fatalf("Error while sending data to client stream: %v", err)
				return nil
			}
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
	maxpb.RegisterMaxServiceServer(s , &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
