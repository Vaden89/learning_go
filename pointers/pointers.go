package main

import "fmt"

//Pointers are a means of passing a reference "address" to a value stored in memory, without passing the actual data being stored

//In this case the parameter is passed by value, hence creating a copy of the value which can be manipulated independently of  the original value
func passByVal(ival int) {
	ival = 0
}

//In this case the parameter is passed by reference, meaning that it takes in a pointer as parameter, the pointer code inside then gets the actual values in memory and assigns it a new value
func passByRef(iptr *int) {
	//pointer
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial value: ", i)

	passByVal(i)
	fmt.Println("After the value has been passed by value: ", i)

	//we use the & symbol before the value to give a pointer to the value i
	passByRef(&i)
	fmt.Println("After the referenve to the value has been passed: ", i)
}
