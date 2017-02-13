// Package main provides ...
package main

import (
	"fmt"
)

func main() {

	var clubs [6]string
	clubs[0] = "Real Madrid CF"
	clubs[1] = "Bayern Munich"
	clubs[2] = "Paris Saint Germain"
	clubs[3] = "Arsenal FC"
	clubs[4] = "Tottenham Hotspur"
	clubs[5] = "Sevilla FC"

	// Get items from position 0 up until position 3
	var firstThree = clubs[:3]
	fmt.Println(firstThree)

	// Get all elements of the array starting at position 3
	var lastThree = clubs[3:]
	fmt.Println(lastThree)

	// loop through the last Three elements
	var sliceLength int = len(lastThree)

	for i := 0; i < sliceLength; i++ {
		fmt.Println(lastThree[i])
	}
}
