package main

import "fmt"

//A recursive function is a function that calls it's self
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main(){
	fmt.Println(fib(7))
}