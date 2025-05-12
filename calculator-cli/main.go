package main

import (
	"fmt"
	"strings"
)

func add(a uint, b uint) {
	var result = a + b
	fmt.Printf("The sum of %v and %v is: %v\n", a, b, result)
}

func subtract(a uint, b uint){
	var result = a - b
	fmt.Printf("%v subtracted by %v results in: %v\n", a,b, result)
}

func divide(a uint, b uint){
	var result = a / b
	fmt.Printf("%v divided by %v results in: %v\n", a,b, result)
}

func multiply(a uint, b uint){
	var result = a * b
	fmt.Printf("%v multiplied by %v results in: %v\n", a,b, result)

}

func main() {
	for {

		fmt.Println("Welcome to your cli calculator!!!")
		fmt.Println("What would you like to do? Pick an option!!")
		fmt.Printf("1. Add \n2. Subtract \n3. Divide \n4. Multiply \n")

		var selectedOption string
		fmt.Scan(&selectedOption)

		selectedOption = strings.ToLower(selectedOption)

		if selectedOption != "add" && selectedOption != "subtract" && selectedOption != "multiply" && selectedOption != "divide" {
			fmt.Println("You did not select a valid option! please select a valid option next time")
			fmt.Println("#############################")
			continue
		}

		var value1, value2 uint
		fmt.Printf("Input the first value you want to %v: ", selectedOption)
		fmt.Scan(&value1)

		fmt.Printf("Input the second value you want to %v:", selectedOption)
		fmt.Scan(&value2)
		

		if(selectedOption == "add"){
			add(value1, value2)
			continue
		}

		if(selectedOption =="subtract"){
			subtract(value1, value2)
			continue
		}

		if(selectedOption =="divide"){
			divide(value1, value2)
			continue
		}

		if(selectedOption =="multiply"){
			multiply(value1, value2)
			continue
		}
	}
}
