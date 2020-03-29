package main

import(

	"fmt"
	"log"
	"os"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	"github.com/Go-protobuf/simple/simple-proto"
)

func main(){

	sm := doSimple()
	writeToFile("simple.bin",sm)
	sm2 := doSimple()
	readFromFile("simple.bin",sm2)
	fmt.Println("Info contained into file:",sm2)

}

func writeToFile(fname string, pb proto.Message){

	out, err := proto.Marshal(pb)
	if(err != nil){
		log.Fatalln("Can't serialise to bytes", err)
		os.Exit(1)
		
	}
	// set in file mode
	if err := ioutil.WriteFile(fname,out,0644) ; err != nil{
		log.Fatalln("Can't write to file",err)
		os.Exit(1)
	}

	fmt.Println("Data has been written!")
}

func readFromFile(filename string, pb proto.Message) error{

	in, err := ioutil.ReadFile(filename)
	if(err != nil){
		log.Fatalln("Something went wrong when reading a file", err)
		os.Exit(1)
	}

	err2 := proto.Unmarshal(in, pb)

	if(err2 != nil){
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err)
		os.Exit(1)
	}
	return nil
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
