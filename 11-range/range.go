package main

import "fmt"

//iterate over data structure(we use range())


func main(){

	//ARRAY

	nums := []int{10,20,3,4,5,3,3}

	//Iterating using for loop
	for i:=0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}


	sum := 0

	//very easy syntax like directly getting element and index using range
	for idx, ele := range nums{
		sum+=ele;
		fmt.Println(idx, " -> " , ele)
	}

	fmt.Println("Total : ", sum)

	// -------------------------------------------
	//MAP
	//You can't use traditional for loop to iterate the map in GO you have to use the range function to iterate the map.

	m := map[string]string{"name":"Jagdish kumawat", "gmail":"jagdishkumawat81@gmail.com"}
	
	for k,v := range m{
		fmt.Println(k, " -> ",v)
	}



	//STRING

	name:="Jagdish Kmwt"


	for s:=0; s < len(name);s++{
		fmt.Println(s, " :: ", string(name[s]))
		//it will print the index and the ascii value of string name[s] print the ascii value
		//string(name[s]) -> print actual character
	}


	for i, ch := range name{
		fmt.Println(i, " -> ", string(ch))
	}




}