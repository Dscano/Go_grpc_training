package main

import(

	"fmt"
	"log"
	"os"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	"github.com/Go-protobuf/addressbook/proto"
)

func main(){

	ad := doAddressbook()
	writeToFile("addressbook.bin",ad)
	ad2 := doAddressbook()
	readFromFile("addressbook.bin",ad2)
	fmt.Println("Info contained into file:",ad2)
}


func doAddressbook()*addresspb.Person{
	ad := addresspb.Person {
			Name: "Pippo",
			Id : 1,
			Email: "pippo@gmail.com",
			Phones:[]*addresspb.Person_PhoneNumber {
				&addresspb.Person_PhoneNumber {
					Number : "0583-579864",
					Type : addresspb.Person_HOME,
				},
			},

		}

	fmt.Println("Print addressbook istance", ad)

	return &ad

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
