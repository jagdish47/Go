package main

import "fmt"


func main(){

	// var age int = 18

	// age:=15
	// if age >= 18{
	// 	fmt.Println("The Person is an Adult")
	// }else{
	// 	fmt.Println("The Person is not an Adult")
	// }


	// age:=55
	// if age >= 18{
	// 	fmt.Println("Adult guy")
	// }else if age >= 12{
	// 	fmt.Println("Teenager guy")
	// }else if age < 12 {
	// }else{
	// 	fmt.Println("Child baby")
	// }


	var role = "admin"
	var hasPermission = false

	if(role == "admin" && hasPermission){
		fmt.Println("Yes, authorized")
	}


	//go does not have Ternary Operator, you wll have to use normal if else.


}