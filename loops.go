// Package main provides ...
package main

import (
    "fmt"
)

func main() {

    var count int = 12

    // Loop with an iterator (Classic)
    for i := 0; i < count; i++ {
        fmt.Println(i)
    }

    // For with a single condition (While loop)
    for count < 23{
       fmt.Println(count)
       count++
    }

    // Infinite Loop
    for {
        fmt.Println("Hello Karim")
        break
    }

}
