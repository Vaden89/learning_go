package main

import "fmt"

func main() {
	//I'm thinking of slices in go as a dynamic array, when just initialized the slices contains 0 elements

	//declare
	var s []int
	fmt.Println("Uninitialized Slice: ", s)

	//We can pass a capcity to make, this creats a slice that already has a length of the number passed with all the elements being zeroed
	s = make([]int, 3)
	fmt.Println("Empty Slice: ", s)

	//Unlike arrays, slice have richer functions such as append that return the new slice with additional values
	//Append does a copy of the original slice and adds the new elements to the slice
	s = append(s, 4, 5)
	fmt.Println("Slice with appended value: ", s)

	//You can copy one slice into another using the copy functions
	c := make([]int, len(s))
	copy(c, s)

	//Slice support the slice operation, which removes a section of a slice based off of index
	// syntax: slice[low:high]
	l := s[2:5]
	fmt.Println("sl1: ", l)

	//We can slice up to (but exluding ) using;
	l = s[:3]
	fmt.Println("Slice up to index[3]: ", l)

	//We can slice from (inclusive) a specific index and stop at the end of the Slice[Arr]
	l = s[2:]
	fmt.Println("Slice from index[2]: ", l)

	//We can declare and initialize a slice in a single line by;
	t := []string{"g", "h", "i"}
	fmt.Println("declared and initialized: ", t)
	
}
