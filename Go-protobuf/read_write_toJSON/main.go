package main

import(

	"fmt"
	"log"
	"os"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/Go-protobuf/simple/simple-proto"
)

func main(){

	sm := doSimple()

	smAsString := toJSON(sm)
	
	fmt.Println("Convertion to JSON:",smAsString)
	smread := &simplepb.SimpleMessage{}
	fromJSON(smAsString,smread)
	fmt.Println("Successfuly create a proto structure:",smread)

}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if(err != nil){
		log.Fatalln("Can't convert to JSON", err)
		os.Exit(1)
		
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in,pb)
	if(err != nil){
		log.Fatalln("Culden't unmarshal the JSON", err)
		os.Exit(1)
		
	}
}


func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println("Print Simple Message istance",sm)

	sm.Name = "I renamed you"
	fmt.Println("Print the new Simple Message istance",sm)

	fmt.Println("The ID is:", sm.GetId())

	return &sm
}
