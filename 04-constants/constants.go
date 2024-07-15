package main

import (
	"fmt"
	"reflect"
)

//you can declare variable outside of function
const age = 30;
var greet string = "Good morning"

// one:=123 -> you can't do it, outside the main function

func main(){



	const name string = "jagdish"

	// name = "asnhu" -> you cann't redecalre the constant variable give you error


	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(greet)


	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)

	fmt.Println(reflect.TypeOf(port))
	fmt.Println(reflect.TypeOf(host))

}