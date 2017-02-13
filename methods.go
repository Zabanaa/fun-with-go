package main

import (
	"fmt"
)

type Club struct {
	name, stadium, manager string
}

func (c *Club) whoAmI() {
	fmt.Printf("We are %s, our stadium is %s. Please meet %s, our lovely manager", c.name, c.stadium, c.manager)
}

// Use pointers to structs in methods that will mutate the values

func main() {

	arsenal := Club{"Arsenal FC", "Emirates Stadium", "Arsene Wenger"}
	arsenal.whoAmI()

}
