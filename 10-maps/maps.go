package main

import "fmt"

//maps are like Object

func main(){


	//define map
	var m = make(map[string]string) 
	//map is keyword
	//key -> string
	//value -> string

	//adding key and values
	m["name"] = "jagdish kumawat"
	m["location"] = "Udaipur, Rajasthan"

	//getting values
	//if key is not present it will return 0 value based on datatype if bool false, if string empty string 
	fmt.Println(m["name"])
	fmt.Println(m["location"])

}