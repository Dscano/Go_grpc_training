package main

import (
	"fmt" 
	"log"
	"context"
	"time"
	"io"
	"google.golang.org/grpc"
	"github.com/exercise_4/max"
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

	c := maxpb.NewMaxServiceClient(cc)
	doMax(c) 
	
}


func doMax(c maxpb.MaxServiceClient){
	fmt.Println("Starting to do a Compute average Streaming RPC..")

	stream, err := c.FindMax(context.Background())
	if err != nil{
		log.Fatalf("error while opening streaming:%v", err)

	}
	waitc := make(chan struct{})
	//function to send a bunch of messages
	go func(){
		numbers := []int32{1,5,3,6,2,20,54,23}
		for _, num:= range numbers {
			fmt.Println("Sending number %v",num)
			stream.Send(&maxpb.MaxRequest{
				Number : num,
				})
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
