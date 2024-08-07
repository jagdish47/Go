package main

import (
	"fmt"
	"time"
)

//Order Struct

type order struct{

	id string
	amount float32
	status string
	createdAt time.Time

}

//receiver type
func (o *order) changeStatus(status string){

	o.status = status

}


//it's like constructor
func newOrder(id string, amout float32, status string) *order{
	myOrder:= order{
		id:id,
		amount: amout,
		status: status,
		//if you're not giving any value it will give you nil

	}

	return &myOrder
}


func main(){


	//if you dont set any field, default value is zero value
	//int => 0 float=>0, string=>"", boolean=>false
	// myOrder:= order{
	// 	id:"1",
	// 	amount: 33.33,
	// 	status: "received",
	// 	//if you're not giving any value it will give you nil

	// }

	myOrder := newOrder("1",33.33, "received")

	myOrder.changeStatus("confirmed")

	//you can assign values later as well.
	myOrder.createdAt = time.Now()
	//get the things
	fmt.Println(myOrder.status)

	fmt.Println("Order Struct", myOrder)


	//another way to create struct

	language := struct {
		name string
		isGood bool
	}{"golang",true}

	//give values in same order


}