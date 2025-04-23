package main

import (
	"fmt"
	"unicode/utf8"
)

//In go strings are a slice of bytes, these bytes are used to represent texts that have been encoded in UTF-8, elements we would typically refer to as characters are known as runes
func main() {
	//In go string literals are just UTF-8 encoded texts
	const name = "isaac"

	//We can get the length of a string same way we can get the length of a slice since strings are basically slices of bytes
	fmt.Println("Len: ", len(name))

	//If we index into a string we get the raw byte value of each character at each index

	//This is fast becoming my favorite way to write for loops
	for i := range name {
		fmt.Printf("%x ", name[i])
	}

	//We can get the number of runes in a string using the utf8 package
	fmt.Println("\nRune Count: ", utf8.RuneCountInString(name))
}
