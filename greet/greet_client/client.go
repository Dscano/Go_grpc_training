package main

import (
	"fmt" 
	"log"
	"google.golang.org/grpc"
	"github.com/greet/greetpb"
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
	fmt.Println("Created client:%f",c)
}