package main

import "fmt"

// func add(a,b int) ing{} -> if have same type you can write like this as well
func add(a int, b int) int{
	return a + b
}

//you can return multiple values you have to just give type in bracket() that what values your returning
func getLanguages()(string, string, int){

	return "JavaScript", "C++", 100;
}




func main(){


	//  ans := add(10,5)
	//  fmt.Println(ans);

	//if something your not using whatever return you can use (_dash) there so no need to use it.
	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2 )


}