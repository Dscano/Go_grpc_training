package main

import(

	"fmt"

	"github.com/Go-protobuf/enums/simple-proto"
)

func main(){

	doEnum()
}


func doEnum(){
	em := enumpb.EnumMessage{
		Id:         45,
		DayOfTheWeek:   enumpb.DayOfTheWeek_THURSDAY,
	}

	fmt.Println("Print Enum Message istance",em)

	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY
	fmt.Println("Print the new Simple Message istance",em)

	fmt.Println("The ID is:", em.GetId())
}
