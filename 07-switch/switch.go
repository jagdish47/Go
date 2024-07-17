package main

import (
	"fmt"
	"time"
)


func main(){




	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().Clock())

	


	var operation string = "*"

	one:=10
	two:=2


	switch operation {

	case "+":

		fmt.Println(one + two)
		// break -> no need of break, the go handle it itself

	case "-":
		fmt.Println(one - two)

	case "*":
		fmt.Println(one * two)

	case "/":
		fmt.Println(one / two)

	default:
		fmt.Println("Invalid Operator Entered, ☹️")

	}


}