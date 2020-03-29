package main

import(

	"fmt"
	"github.com/Go-protobuf/simple/simple-proto"
)

func main(){
	doSimple()
}

func doSimple() simplepb.SimpleMessage {
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

	return sm
}
