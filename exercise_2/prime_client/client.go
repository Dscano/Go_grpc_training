package main

import (
	"fmt" 
	"log"
	"io"
	"context"
	"google.golang.org/grpc"
	"github.com/exercise_2/prime"
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

	c := primepb.NewPrimeServiceClient(cc)
	doServerStreaming(c) 
	
}


func doServerStreaming(c primepb.PrimeServiceClient){
	fmt.Println("Starting to do a Server Streaming RPC..")
	req := &primepb.PrimeStreamingRequest{
		Priming: &primepb.Prime{
			Number: 120,
		},
	}

	resStream, err := c.PrimeStreaming(context.Background(),req)
	if err != nil{
		log.Fatalf("error while calling Prime RpC:%v", err)
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

		log.Printf("Response from Prime Streaming: %v", msg.GetResult())
	}
}
