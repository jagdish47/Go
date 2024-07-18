package main

import "fmt"



func main() {




	var nums [5]int  //[0,0,0,0,0]
	fmt.Println("lenght of array : ",len(nums))


	nums[0] = 10
	nums[1] = 20

	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])

	fmt.Println(nums) //[10 20 0 0 0]


	var names[5]string
	fmt.Println(names)//[    ] empty string

	var is[5]bool
	fmt.Println(is)//[false false false false false]

	var deci[5]float32
	fmt.Println(deci)//[0 0 0 0 0]

	//int -> 0
	//bool -> false
	//string -> ""(empty string)



	// -----------------------------------------------------------------


	//to declare it in single line

	one:=[5]int{10,4,5,6,3}
	two:=[4]string{"jagdish","anshu","daya","himmu"}
	three:=[3]bool{true,false,false}
	four:=[4]float32{33.2, 22.2, 10, 20}


	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)


	// -------------------------------------------

	// 2d array
	mat:=[3][3]int{{1,2,3},{6,7,8},{3,3,3}}
	fmt.Println(mat)




	//we have to know size already than we can use array
	//otherwise we use slice(dynamic array) in go lang

	//mostly we use slice in go lang 
	//Memory Optimization
	//constant time access


}