package main

import "fmt"

//as a integer you can pass any number of arguments
//if you want ot pass any type use interface{} like now you can pass anything


// func sum(nums ...interface{}) you can do like this now you can pass anything to this

func sum(nums ...int) int{

	fmt.Println(nums) // [10 20 30 43 43 4 3 43 4 3 3 33 3] convert into array and give you array
	total:=0

	for _, num:=range nums{
		total += num
	}

	return total
}

func main(){



	tot := sum(10,20,30,43,43,4,3,43,4,3,3,33,3)

	fmt.Println("Total Is :: ", tot);



}


//passing n number of argument to a function is known as variadic function