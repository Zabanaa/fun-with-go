package main

import (
	"fmt"
	"strconv"
)

// Note: in your interface, if you have methods that accept pointers
// you need to make sure that you pass a memeory address to whatever function will
// call the interface methods
// for example Staff.changeName takes a pointer to an Employee
// further down in the code we create a getAllInfo function that takes a Staff
// interface. That function needs to be passed a pointer to whatever struct we
// want to use the interface on
// example: we can't call getAllInfo(clara)
// we have to call getAllInfo(&clara)
// if it doesn't make sense
// check this out: https://gobyexample.com/interfaces
// and this as well : http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

type Staff interface {
	getPersonalInfo() string
	changeName(newName string)
}

type Address struct {
	number int
	steet  string
}

type Employee struct {
	name    string
	age     int
	address Address
}

func (employee Employee) getPersonalInfo() string {
	return "Name: " + employee.name + " - Age: " + strconv.Itoa(employee.age)
}

func (employee *Employee) changeName(newName string) {

	employee.name = newName
}

func getAllInfo(s Staff) {
	fmt.Println(s.getPersonalInfo())
}

func main() {

	address := Address{12, "north street"}
	clara := Employee{"clara", 23, address}

	cp := &clara
	getAllInfo(cp)

	// cp.changeName("hello")
	// fmt.Println(clara.getPersonalInfo())

}
