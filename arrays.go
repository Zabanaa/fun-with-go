// Package main mainprovides ...
package main

import (
    "fmt"
)

func main() {

    // Declare an array of exactly 5 ints then assign a value to each position
    var numbers [5] int

    numbers[0] = 345
    numbers[1] = 432
    numbers[2] = 98
    numbers[3] = 15937
    numbers[4] = 007

    fmt.Println(numbers)


    // Shorthand way of creating an array
    var numbers_2 = [4] int {1, 3, 4, 5}
    fmt.Println(numbers_2) // [1, 3, 4, 5]

    // Get the length of an array
    var names = [3] string { "Karim", "Samir", "Moussa" }
    var number_of_names int = len(names)

    fmt.Println(names) // ["Karim", "Samir", "Moussa"]
    fmt.Println(number_of_names) // 3

    // Loop over an array the classic way
    for i := 0; i < len(names); i++ {
        fmt.Println(names[i])
    }

}
