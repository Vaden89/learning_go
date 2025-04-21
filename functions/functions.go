package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// You can group the type of multiple parameters
func multiply(a, b, c int) int {
	return a * b * c
}

// you can have multiple values being returned, almost like when you returning a turple of values in rust
func vals() (int, int) {
	return 3, 7
}

//We can have something called a variadic function, which allows our function to accept an abiritary number of arguments

func sum(nums ...int) {
	fmt.Println("These are the numbers being added: ", nums)
	total := 0

	//range nums returns two values: the index of the element being the first and the value being the second, in this case we burn the index
	for _, num := range nums {
		total += num
	}
	fmt.Println("The total is: ", total)
}

//Note that you have to explicitly return something in go, it doesn't automatically return the last experession like in rust

func main() {
	result := plus(1, 2)
	fmt.Println("The result of 1 + 2 is: ", result)

	result = multiply(3, 4, 5)
	fmt.Println("The result of 3 * 4 * 5 is: ", result)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//You can also burn on of the values being returned using the _, "blank indentifier"
	_, c := vals()
	fmt.Println(c)

	sum(1, 2, 3)

	nums := []int{4, 5, 6}
	//You can also pass a slice as the parameter, but you'd have to spread it using ... operator
	sum(nums...)
}
