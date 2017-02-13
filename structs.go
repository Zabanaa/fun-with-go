package main

import (
	"fmt"
)

// Declare a struct here
type Club struct {
	name    string
	stadium string
	manager string
}

// Pointers to structs are automatically dereferenced
func changeClub(club *Club, newName string) {
	// The two statements below are actually the same
	club.name = newName
	(*club).name = newName
}

func main() {

	// Longhand notation
	var arsenal Club
	arsenal.name = "Arsenal FC"
	arsenal.stadium = "Emirates Stadium"
	arsenal.manager = "Arsene Wenger"

	fmt.Println(arsenal)

	// Shorthand notation
	myClub := Club{name: "Real Madrid CF", manager: "Zizou"}
	fmt.Println("My club is", myClub.name)

	// We can also pass pointers to structs around
	changeClub(&myClub, "West Ham FC")

	fmt.Println("My club is now", myClub.name)

}
