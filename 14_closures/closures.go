package main

import "fmt"

func counter() func() int{


	var count int


	return func() int{
		count+=1 //whenever we call the outer scope of fucntion alaway there for inner scope(Clouser)

		return count;
	}
}


func main(){


	increment:=counter();

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

}