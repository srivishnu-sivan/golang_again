package main

import (
	"fmt"
	"time"

	
)


type orders struct{
	id string
	amount float32
	status string
	createAt time.Time // nanosecond precision

}

// ? recevier type
func (o *orders) chagneStatus(status string){
	o.status = status
}

func (o orders) getAmount() float32{
return o.amount
}


func main(){

	myOrder := orders{
		id : "1",
		amount : 50.00,
		status : "received",
		createAt: time.Now(),
	}
	myOrder.chagneStatus("confirmed")

	fmt.Println(myOrder)


	amount  := myOrder.getAmount()
	fmt.Println(amount)

}