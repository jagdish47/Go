package main

import "fmt"

//whenever passing liek this is is passing by value like one new copy will create, main variable like passing thing not change.




func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
}


//Passing by reference : means we are passing the memory address of it, and change the value at that memory address so it will change value to that address

func changeNumRef(num *int){
	*num = 10
	fmt.Println("Changing by Ref :: ", num); //0xc00000a0b8
	fmt.Println("Changing by  :: ", *num); //10

}

func main() {

	num:=1


	// fmt.Println("Memory address :: ", &num)  // 0xc00000a0b8
	// changeNum(num)
	// fmt.Println("After change num : ", num)

	changeNumRef(&num)
	fmt.Println("After changing :: ", num)

}