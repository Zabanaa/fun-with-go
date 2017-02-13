// Package main provides ...
package main

import (
    "fmt"
)

func main() {

    // Variable declaration + assignement
    var number int
    number = 4
    fmt.Println("Number 1 is", number)

    // Increment Values

    number++
    fmt.Println("Number 1 is now", number)

    // Shorthand
    var number_2 int = 3
    fmt.Println("number_2 is", number_2)

    // Decrement Values
    number_2--
    fmt.Println("number_2 is now", number_2)


    // In this case, the type is implied
    number_3 := 454
    fmt.Println("number_3 is", number_3)

    // Multiply Values
    number_3 *= 2
    fmt.Println("number_3 is now", number_3)

    // Divide Values
    number_3 /= 4
    fmt.Println("number_3 is now", number_3)

    // Variable Types

    var name string         = "Karim"
    var age  int            = 25
    var isProgrammer bool   = true
    var livesInJapan bool   = false
    var num byte            = 235
    // var num byte            = 3454 // Not Allowed

    // Constants
    const API_KEY string = "somesecretapikey"

    // API_KEY = "somechangedstring" // Not Allowed: Cannot Assign to API_KEY



    fmt.Println("API_KEY is ", API_KEY)
    fmt.Println("num is a byte", num)
    fmt.Println("Does Karim live in Japan ?", livesInJapan)
    fmt.Println("Is Karim a Programmer", isProgrammer)
    fmt.Println("Karim is ", age)
    fmt.Println("Hey my name is ", name)


}
