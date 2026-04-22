package main

import (
	"fmt"
)

// pointer or address of memory (ram)
type User struct {
	Name   string
	Age    int
	Salary float64
	Foods  []string
}

// pass by value

func print1(numbers [3]int) {
	fmt.Println(numbers)
}

// pass by reference
func print(numbers *[3]int) {
	fmt.Println(numbers)
}

func main() {

	// x := 30
	// fmt.Println(x)
	// //when you print address you call &variable computer ans hexadecimal format
	// printAddress := &x
	// valueAtAddress := *printAddress //when you find the value of thus object you call *variable

	// *printAddress = 23
	// fmt.Println(x)
	// fmt.Println(printAddress)
	// fmt.Println("value at the address", valueAtAddress)

	// arr := [3]int{121, 232, 343}

	// print1(arr) //pass bye value
	// print(&arr) //pas by reference

	obj := User{
		Name:   "anamul",
		Age:    26,
		Salary: 25013.00,
		Foods:  []string{"anamul", "rice"},
	}
	user := &obj

	// fmt.Println(*user)
	// fmt.Println(user.Name)

	// fmt.Println(obj)
	fmt.Println(user)
}
