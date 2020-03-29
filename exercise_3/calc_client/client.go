package main

import (
	"fmt" 
	"log"
	"context"
	"google.golang.org/grpc"
	"github.com/exercise_3/calc"
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

	c := calcpb.NewCalcServiceClient(cc)
	doCalcAverage(c) 
	
}


func doCalcAverage(c calcpb.CalcServiceClient){
	fmt.Println("Starting to do a Compute average Streaming RPC..")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil{
		log.Fatalf("error while opening streaming:%v", err)

	}
	numbers := []int32{3,5,9,54,23}

	for _ , number := range numbers{
		fmt.Println("Sending number: %v \n", number)
		stream.Send(&calcpb.ComputeAverageRequest{
			Number : number,
		})
	}
	
	
	
	res, err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("error while recivibg response from  LongGreet:%v", err)

	}

	fmt.Println("The Average is: %v\n", res.GetNumber())
	
}