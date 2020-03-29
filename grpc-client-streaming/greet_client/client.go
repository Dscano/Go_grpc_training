package main

import (
	"fmt" 
	"log"
	"context"
	"time"
	"google.golang.org/grpc"
	"github.com/grpc-client-streaming/greetpb_client_streaming"
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
	doClientStreaming(c) 
	
}


func doClientStreaming(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do a Client Streaming RPC..")
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Pippo",
		    },

		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Maestro",
		    },

		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Matrax",
		    },

		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Pimpi",
		    },
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil{
		log.Fatalf("error while calling LongGreet:%v", err)

	}
	//we iterate over our slice and send each message individually
	for _,req := range requests{
		fmt.Println("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000*time.Millisecond)

	}
	res, err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("error while recivibg response from  LongGreet:%v", err)

	}
	fmt.Println("LongGreet Response: %v\n", res)
	
}