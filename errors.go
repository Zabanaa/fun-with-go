package main

import (
	"errors"
	"fmt"
)

// By convention, errors are always the last return value from a function
// and have a type "error", which is a built in interface

func addOne(number int) (int, error) {

	if number == 42 {
		// In this case, we have an error, so we return -1 for the int
		// and errors.new() along with a message for the error
		return -1, errors.New("Wrong number fam, sorry")
	}

	// there was no error, so we return the result and nil for the error
	// if the value for error is nil, then it means everything went fine
	return number + 1, nil
}

// We can also create custom error types and add methods that Impletment the
// Error() interface like so

type wrongNumberError struct {
	argument int
	message  string
}

func (err *wrongNumberError) Error() string {
	return fmt.Sprintf("Error: %d - %s", err.argument, err.message)
}

func addThree(number int) (int, error) {

	if number == 42 {
		// In this case we create a new struct containing the argument
		// and a custom error message
		// but notice that we return an address to the struct
		err := wrongNumberError{number, "Sorry but you passed a wrong number"}
		return -1, &err
	}

	return number + 3, nil
}

func main() {

	var wrongNumber int = 42
	var ages = []int{23, 45, wrongNumber}

	for _, age := range ages {

		result, err := addOne(age)

		if err != nil {
			fmt.Println("Failed", err)
		} else {
			fmt.Println("Passed", result)
		}

	}

	for _, age := range ages {

		result, err := addThree(age)

		if err != nil {
			fmt.Println("Failed", err)
		} else {
			fmt.Println("Passed", result)
		}

	}

}
