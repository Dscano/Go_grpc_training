package main

import (
	"fmt" 
	"log"
	"context"
	"time"
	"io"
	"google.golang.org/grpc"
	"github.com/grpc-bidirectional-streaming/greetpb_bidirectional"
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
	doBiDiStreaming(c) 
	
}

func doBiDiStreaming(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do a BiDi Streaming RPC..")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil{
		log.Fatalf("error while calling LongGreet:%v", err)
		return
	}
	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Pippo",
		    },

		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Maestro",
		    },

		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Matrax",
		    },

		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Pimpi",
		    },
		},
	}

	waitc := make(chan struct{})
	//function to send a bunch of messages
	go func(){
		for _, req:= range requests {
			fmt.Println("Sending message : %v \n", req)
			stream.Send(req)
			time.Sleep(1000*time.Millisecond)
		}
		stream.CloseSend()
	}()

	//function to receive a bunch of messages
	go func(){
		for{
			res, err := stream.Recv()
			if err == io.EOF{
				break;
			}
			if err != nil{
				log.Fatalf("error while recivig :%v", err)
				break;
			}
			fmt.Println("Received: %v \n", res.GetResult())
		}
	}()
	<-waitc	
}