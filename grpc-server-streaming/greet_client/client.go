package main

import (
	"fmt" 
	"log"
	"io"
	"context"
	"google.golang.org/grpc"
	"github.com/grpc-server-streaming/greetpb_server_streaming"
)

func main(){
	fmt.Println("Hello, I am a client")
	//Opening an insecure connection with 50051
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("could not connetct:%v", err)
	}
	// when all main is done the cc.close is done
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doServerStreaming(c) 
	
}


func doServerStreaming(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do a Server Streaming RPC..")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Pippo",
			LastName: "Bomber",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(),req)
	if err != nil{
		log.Fatalf("error while calling GreetManyTimes RpC:%v", err)
	}
	for{
			msg, err :=resStream.Recv()
		if err == io.EOF{
			// we've reached the end of the stream
			break
		}
		if err !=nil{
			log.Fatalf("error while reading stream:%v", err)
		}

		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}