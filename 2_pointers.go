package main

import (
	"fmt"
)

func main() {

	var name string = "karim"

	fmt.Println("Name: ", name)

	p := &name
	*p = "Karim Benzema"

	fmt.Println("Name: ", name)

}
