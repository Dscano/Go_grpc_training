package main

import (
	"fmt" 
	"log"
	"context"
	"google.golang.org/grpc"
	"github.com/grpc-unary-streaming/greetpb_unary"
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
	doUnary(c)
	
}

func doUnary(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do Unary RPC..")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Pippo",
			LastName: "Bomber",
		},
	}
	res, err :=c.Greet(context.Background(), req)
	if err != nil{
		log.Fatalf("error while calling Greet RpC:%v", err)
	}

	log.Printf("Response from Greet: %v",res.Result)

}