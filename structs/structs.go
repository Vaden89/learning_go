package main

import "fmt"

// Structs are typed collections of fields, typically used in grouping data together to form records, they remind me of inferences fromm typescript
type person struct {
	name string
	age  int
}

// This function creates a new person taking in name as parameter and returning a pointer to the created record, it is idiomatic in go to have a contructor function for creating structs
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

// GO allows for structs to have methods, these are functions encased within structs
type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) prem() int {
	return 2*r.width + 2*r.height
}

func main() {
	//Basic syntax to create a new struct
	fmt.Println(person{name: "Jonathan", age: 19})

	//When initializing a struct if  you omit a value it is zeroed according to it's data type
	fmt.Println(person{name: "Isaac"})

	//Use of the struct constructor function
	fmt.Println(newPerson("Cifer"))

	//We can access values in a struct using the dot notation
	friend1 := person{name: "Daniel", age: 21}
	fmt.Println(friend1.name)

	//We can also do the same using a pointer to the struct
	friend1Pointer := &friend1
	fmt.Println(friend1Pointer.age)

	//Values within structs can be mutated
	friend1Pointer.age = 19
	fmt.Println(friend1Pointer.age)

	r := rect{width: 10, height: 5}

	//We use dot notation to run methods
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.prem())

	//It also works with pointers too
	rp := &r
	fmt.Println("Area: ", rp.area())
	fmt.Println("Perimeter: ", rp.prem())
}
