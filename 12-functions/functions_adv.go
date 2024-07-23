package main

import "fmt"

//you can pass a function to another function as parameter
func processIt(fn func(a int) int) {
	resutl := fn(10)

	fmt.Println(resutl)
}

//you can return a fucntion as well

//telling that this will return a fuction
func didIt() func(a int) int{

	// the function to return
	return func(a int)int{
		return 10+10
	}

}

func main() {

	//First Class function you can pass a fucntion to variable
	fn := func(a int) int {
		return a
	}

	processIt(fn)

	// called first function it will return a functin
	res := didIt()
	fmt.Println(res);

	// calling returned function
	resul := res(10)

	//returned function returning a value printing that
	fmt.Println(resul)

}