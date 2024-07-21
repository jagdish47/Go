package main

import (
	"fmt"
	"maps"
)

func main(){



	data := make(map[string]string)


	data["name"] = "jagdish kumawat"
	data["age"] = "24"
	data["mail"] = "jagdishkumawat81@gmail.com"
	data["gender"] = "male"



	fmt.Println(data["name"])
	fmt.Println(data["kk"]) //if not availabe such key will return 0, false, "" based on datatype of value


	fmt.Println(len(data)) //give length total

	delete(data, "name") //to delte value using key
	clear(data) //to clear the data

	fmt.Println(data) 

 
	// -----------------------------

	//another way to create the map

	m := map[string]int{"price":343, "size":34, "phone":9549392949}
	ma := map[string]int{"price":343, "size":34, "phone":9549392949}

	fmt.Println(m)



	k, ok := m["prices"]

	//true/false   value

	if ok {
		fmt.Println("All Ok brother",k,  ok)
	}else{
		fmt.Println("Not Ok baba",k,  ok)
	}

	//How to compare two maps


	fmt.Println(maps.Equal(m, ma))


}