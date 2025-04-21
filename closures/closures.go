package main

import "fmt"

// A closure is a function that returns another function, allowing access the values within the original function after the orginal function has been dropped

func intSeq() func() int {
	i := 0
	//this is an anon func being returned
	// An anon func is one without a name
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt());
	fmt.Println(nextInt());
	fmt.Println(nextInt());

	newInts := intSeq()
	fmt.Println(newInts())
}