package main

import "fmt"

//there is no while keyword in go.
//for looping only for we use, using for we create while loop as well
func main(){



	// 1
	//while loop(using for loop)
	// i:=1
	// for i <=5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 2
	// infinite loop
	// for{}

	// 3
	// for loop
	// for i:=1; i<=5; i++{
	// 	if i == 3{
	// 		break
	// 	}
	// 	if i == 1{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// 4
	// Range 
	for i:= range 10 {
		fmt.Println(i);
	}

}



