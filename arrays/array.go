package main

import "fmt"

func main() {
	//Initialize an array, if you only initialize an array all it's elements would contain just zero values depending on the data type though [0,0,0,0,0]

	var numArr [5]int
	fmt.Println("Empty array: ", numArr)

	var numArrLength = len(numArr)
	fmt.Println("The length of the number array is: ", numArrLength)

	//Initialize and declare an array using this syntax
	b := [5]int{1,2,3,4,5}
	fmt.Println("A new array of numbers: ", b)

	//or using this syntax

	var c = [2]int{0,1}
	fmt.Println("New array declared using a different method: ", c)

	//You can specifiy the index of an element you're declare in an array using <index>:<value>, any index skipped the value is then zeroed 
	b = [...]int{1,2, 4: 10}
	fmt.Println("Array declared with index element specified: ", b)

	//There are also two dimensional array "Basically arrays that contain other arrays"
	var twoDimensionalArr = [2][3]int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Println("This is a two dimensional array: ", twoDimensionalArr)
}
