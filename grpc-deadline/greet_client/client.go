package main

import (
	"fmt" 
	"log"
	"context"
	"time"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc"
	"github.com/grpc-deadline/greetpb_deadline"
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

	doUnaryDeadline(c, 5* time.Second) //should complete
	doUnaryDeadline(c, 1* time.Second) //should uncomplete
	
	
}

func doUnaryDeadline(c greetpb.GreetServiceClient, timeout time.Duration){
	fmt.Println("Starting to do Unary with Deadline RPC..")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Pippo",
			LastName: "Bomber",
		},
	}
	ctx, cancel:= context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err :=c.GreetWithDeadline( ctx, req)
	if err != nil{
		statusErr,ok :=status.FromError(err)
		if ok{
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Printf("Time out was hit! Deadline was exceeded")
			} else {
				fmt.Printf("unexpected error: ", statusErr)
			}
		} else {
			log.Fatalf("error while calling Greet with Deadline RpC:%v", err)
		}
		return
	}
	
	log.Printf("Response from Greet: %v",res.Result)

}