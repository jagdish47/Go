package main

import (
	"fmt"
	"slices"
)


func main(){

	nums:=[]int{}

	nums = append(nums, 10)
	nums = append(nums, 20)
	nums = append(nums, 30)
	nums = append(nums, 40)
	nums = append(nums, 50)
	nums = append(nums, 60)

	//COPY

	var numCopy = make([]int, len(nums))
	//to copy that the size must be equal to that
	copy(numCopy, nums);

	// fmt.Println(nums)
	// fmt.Println(numCopy)

	
	//SLICE

	fmt.Print(nums[2:3]) //only index-2, 3 exclude
	fmt.Print(nums[:3]) //start not given it start from start
	fmt.Print(nums[3:]) //if you not give end than go till end, count from 0,1,2,3(start) till end


	//EQUAL


	fmt.Println(slices.Equal(nums, numCopy)) // if both same than true, otherwise false



	//2D Slice

	var mat = [][]int{{1,2,3},{4,5,6},{9,0,0}}
	fmt.Println(mat)

}