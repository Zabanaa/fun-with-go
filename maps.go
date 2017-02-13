// Package main provides ...
package main

import (
	"fmt"
)

func main() {

	// Create an empty map of ints
	var myMap = make(map[string]int)

	// Add key value pairs to the map
	myMap["route"] = 45
	myMap["numberOfWins"] = 9

	// Get the number of key value pairs in the map
	var mapLength int = len(myMap)

	fmt.Println(myMap)
	fmt.Println("The length of myMap is", mapLength)

	// Create a map of strings using shorthand notation
	clubs := map[string]string{"London": "Arsenal", "Paris": "PSG"}
	fmt.Println(clubs)

}
