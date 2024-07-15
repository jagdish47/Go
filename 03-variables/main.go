package main

import "fmt"


func main(){


	var name string = "golang"
	fmt.Println(name);

	//go automatically infer the type of it
	var greet = "How are you"
	fmt.Println(greet)


	//if we dont want to write type and var this is shorthand
	again:="againtry47"
	fmt.Println(again)


	//if we are declare variable but want to assign value later must give type to it.
	var one string
	one = "Hey one"
	fmt.Println(one)


	//above are valid for all datatype int, float, boolean, 
}