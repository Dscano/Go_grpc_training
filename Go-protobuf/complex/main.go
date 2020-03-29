package main

import(

	"fmt"
	"github.com/Go-protobuf/complex/simple-proto"
)

func main(){

	doComplex()
}


func doComplex(){
	cm := complexpb.ComplexMessage{
		OneDummy:    &complexpb.DummyMessage{
			Id : 1,
			Name: "First Message",
		},

		MultipleDummy:[]*complexpb.DummyMessage{
			&complexpb.DummyMessage{
			Id : 2,
			Name: "Second Message",

		},
		&complexpb.DummyMessage{
			Id : 3,
			Name: "Third Message",

		},
			
	},
}

	fmt.Println("Print Enum Message istance",cm)

}
