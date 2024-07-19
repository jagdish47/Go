package main

import "fmt"

//slice -> dynamic array (change size whenever add items, no fix size)
//provide good methods, most used in go

func main(){

	//it's same like array but you don't have to give size
	//empty slice always be nil([]) like null
	var num[]int
	fmt.Println(num == nil) //to check weather is empty or not(nil or not)
	fmt.Println(len(num)) //0



	// ---------------------------------------------------------------------
	//this is how we can give initial size, it's not nil now
	var nums = make([]int, 0) //if you give size than 0 will be there and if you add new item 0 will be there and next item will be append


	fmt.Println(len(nums))
	fmt.Println(cap(nums)) //for now the capacity is 5, once you add new item it will(capacity) increase automatically

	//to add values in slice, it's syntax look strange but that's how we do in go
	nums = append(nums, 10)
	nums = append(nums, 20)
	nums = append(nums, 30)

	

	fmt.Println(nums)

	// -----------------------------------------------------------------------
	// best way to create slice is

	number:=[]int{}
	fmt.Println(len(number)) //0
	fmt.Println(cap(number)) //0
	fmt.Println(number) //[]

}