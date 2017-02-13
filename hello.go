package main

import "fmt"

type Employee struct {
    name string
    age int
    address string
}

func change_value(p *int, new_value int) {
    *p = new_value
}

func main(){

    var x int = 23
    fmt.Println("Value of x is", x)

    change_value(&x, 982)

    fmt.Println("Value of x is", x)

    // fmt.Println(p)
    // fmt.Println(susan)
    // fmt.Println("Hello Fam")

}
