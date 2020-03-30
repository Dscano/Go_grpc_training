package main

import (
	"fmt" 
	"log"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"github.com/errors/square"
	"google.golang.org/grpc/codes"
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

	c := squarepb.NewSquareRootServiceClient(cc)
	doErrorUnary(c) 
	
}

func doErrorUnary (c squarepb.SquareRootServiceClient){
	fmt.Println("Starting to do a SquareRoot Streaming RPC..")
	//correct call
	doSquare(c, 10)
	//error call
	doSquare(c, -2)
}


func doSquare(c squarepb.SquareRootServiceClient, n int32){
	
	
	res, err := c.SquareRoot(context.Background(), &squarepb.SquareRootRequest{
			Number:n,
	})
	if err != nil{
		respErr,ok := status.FromError(err)
		if ok {
			// actuall error from gRPCuser error
			fmt.Println("Error message from server:\n",respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably send a negative number")
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: ", err)

		}
	}
	fmt.Println("Result of square root of : ", n, res.GetNumberRoot())	
}
