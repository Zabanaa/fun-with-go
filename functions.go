// Package main provides ...
package main

import (
	"fmt"
)

// 1 - Void function that does not return any value
func printSomething() {
	fmt.Println("Something")
}

// 2 - Function with a return value
func addTwo(number int) int {
	var result int
	result = number + 2
	return result
}

// 3 - Function with multiple return values
func greeting(name string) (string, bool) {
	var isError bool = false
	var new_name string = "Hello there " + name
	return new_name, isError
}

// 4 - Function with multiple NAMED returned values
func greetingTwo(name string) (greet string, isError bool) {
	greet = "Greetings " + name
	isError = false
	return
}

func main() {

	// 4 - Multiple named returned values
	var rimk string = "Rimk Poto"
	rimk, err := greetingTwo(rimk)
	fmt.Println(rimk, "Error:", err)

	// 3 - Multiple return values
	var myName string = "Karim"
	fmt.Println("My name is", myName)

	new_name, err := greeting(myName)
	fmt.Println(new_name, "Error:", err)

	// 2 - Single return value
	var myNum int = 2
	fmt.Println("My num is ", myNum)
	myNum = addTwo(myNum)
	fmt.Println("My num is now", myNum)

	// 1 - Void function
	printSomething()
}
