// Package main provides ...
package main

import (
    "fmt"
)

func main() {

    var legal_age   int = 18
    // var student_age int = 17
    var student_age int = 23

    // If - else
    if student_age < legal_age {
        fmt.Println("You are not allowed to drive")
    } else {
        fmt.Println("You are legally allowed to drive, enjoy.")
    }

    switch legal_age {
        case 18:
            fmt.Println("You are allowed to drive with another adult")
        case 21:
            fmt.Println("You are allowed to drive on your own")
        case 23:
            fmt.Println("You must pay for insurance on your own now")

        default:
            fmt.Println("Whuuh ?!")
    }

}
