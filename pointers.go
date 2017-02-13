package main

import (
	"fmt"
)

// 1 - This function takes a string s and changes its value
// The gotcha here is that this will create a copy of s in memory
func changeValue(s string) {
	fmt.Println("Address of s inside the function", &s)
	s = "Hello newly changed string"
}

// 2 - This function accepts a pointer to a string
// In this example we pass the value as a reference
// Meaning we will go to the memory address of whatever string is passed and change its value
func actuallyChangeValue(s *string) {
	// s in this case is a memory address
	// to actually get to the value we have to dereference it
	// like so : *s = "new value"
	fmt.Println("Address of s inside the func", s)
	*s = "Here we changed the value"
}

func main() {

	// 1 - In this case, the value of s does not change because changeValue() creates a copy of s as you can see by the two different memory addresses
	var s string = "Karim"
	fmt.Println(s)
	fmt.Println("Address of s outside the function", &s)

	changeValue(s)

	fmt.Println(s)

	// 2 - Now we are going to pass the value as reference
	var str string = "My Man"
	fmt.Println("Address of s outside the func", &str)

	actuallyChangeValue(&str)

	fmt.Println(str)
}
