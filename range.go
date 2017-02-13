// Package main provides ...
package main

import (
	"fmt"
)

func main() {

	var clubs = [3]string{"Arsenal", "Everton", "Tottenham"}

	// Using range to loop over an array and its indices
	for index, club := range clubs {
		fmt.Println(index, club)
	}

	// Using range to loop over an array and ignoring the indices
	for _, club := range clubs {
		fmt.Println(club)
	}

	var ages = map[string]int{"Karim": 25, "Zizou": 44}

	// Using range to loop over key value pairs in a map
	for key, value := range ages {
		fmt.Printf("%s -> %d\n", key, value)
		// Note that the order of printing is not guaranteed
	}

	// Using range to only loop over a map's keys
	for key := range ages {
		fmt.Println("Key:", key)
	}

	// Using range on a string
	for index, character := range "Hello" {
		// This will print the ANSII code of every character
		fmt.Println(index, character)

	}

	for index, character := range "Hello" {
		fmt.Println(index, string(character))
		// This will print the actual character
	}

}
