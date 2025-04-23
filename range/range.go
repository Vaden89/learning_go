package main

import "fmt"

//Range can used to iterate over the built in data types of the language
func main() {
	//Declaring and initializing a slice
	nums := []int{2, 3, 4}
	sum := 0

	//using the range keyword on a slice allows us access to elements in the slice and their index, in this case we burn the index using the blank identifier as it's not being used
	for _, num := range nums {
		sum += num
	}

	fmt.Println("The sum of values in the slice: ", sum)

	kvs := map[string]string{"a": "apple", "b": "ball"}

	//We can also use the range keyword on maps to allow for us to iterate over a map, the range keyword gives us access to the key and value of an entry in a map
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//We can also use the range keyword on strings, in go strings are made up of runes, these runs are integars that represent unicode points, strings are basically an array of runes, meaning we can access the index and the data at that index {rune}
	for i, c := range "isaac" {
		fmt.Println(i, c)
	}
}
