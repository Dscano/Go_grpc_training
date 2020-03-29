package main

import (
	"fmt" 
	"log"
	"context"
	"google.golang.org/grpc"
	"github.com/exercise_1/sumpb_unary"
)

func main(){
	fmt.Println("Welcome to Sum")
	//Opening an insecure connection with 50051
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("could not connetct:%v", err)
	}
	// when all main is done the cc.close is done
	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)
	doUnary(c)
	
}

func doUnary(c sumpb.SumServiceClient){
	fmt.Println("Starting to do Unary RPC..")
	req := &sumpb.SumRequest{
			Summing: &sumpb.Sum{
				Number: []int32{3, 10, 50},
			},
		}
	res, err :=c.Sum(context.Background(), req)
	if err != nil{
		log.Fatalf("error while calling Greet RpC:%v", err)
	}

	log.Printf("The result is: %v",res.Result)

}