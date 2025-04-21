package main

import "fmt"

//Note that maps are basically like objects from javascript but only allow for a single, key and value type
func main() {
	//Create a map using the make func 
	m := make(map[string]int)

	//Setting key/value pairs using the traditional key value pair syntax
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	value1 := m["k1"]
	fmt.Println("The value for k1 is: ", value1)

	//If the value accessed  is not contained within the map we get a zero value returned 
	value3 := m["k3"]
	fmt.Println("The value for k3 is: ", value3)

	//We can get the length of a map by calling len() on it, the length of a map being the total number of key value pairs
	fmt.Println("The length of map m is: ", len(m))

	//We can remove a key/value pair from a map using the delete(), taking in the map and the value to be removed as arguments
	delete(m, "k2")
	fmt.Println("Map after value has been deleted: ", m)

	//We can removed all Key value pairs from a map using clear(), taking in the map as an argument
	clear(m)
	fmt.Println("Empty map: ", m)

	//If we do not care about getting the value of a key and just want to check if it exists in the map we can use this;
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	//We can declare and initialize a map using;
	me := map[string]string{
		"firstName": "Isaac",
		"lastName": "Shosanya",
	}
	fmt.Println("This is my user entity: ", me)
}