package main

import (
	"fmt"
)

// Pointer means address of memory

// Pointer or address of memory (ram)

type User struct {
	Name   string
	Age    int
	Salary float64
	Foods  []string
}

// Pass by value

func print(numbers [3]int) {
	fmt.Println(numbers)
}

// Pass by reference

func print1(numbers *[3]int) {
	fmt.Println(numbers)
}

func main() {

	// x := 30
	// fmt.Println(x)
	//When you print address you call &variable computer ans hexadecimal format

	// printAddress := &x //&x means  address of x
	// // fmt.Println(printAddress)

	// valueAtAddress := *printAddress //When you find the value of thus object you call *variable
	// fmt.Println(*printAddress) // find the value

	// *printAddress = 20 // update variable bye pointer thrue
	// fmt.Println(x)

	// fmt.Println(printAddress)
	// fmt.Println("value at the address", valueAtAddress)

	// arr := [3]int{121, 232, 343}

	// print(arr) //Pass bye value

	// print1(&arr) //Pass by reference

	obj := User{
		Name:   "anamul",
		Age:    26,
		Salary: 25013.00,
		Foods:  []string{"bread", "rice", "Meat"},
	}

	user := &obj
	user2 := &obj

	fmt.Println(user)   //pass bye value
	fmt.Println(*user2) // pass bye reference

	fmt.Println(user.Name)

	fmt.Println(obj)

}
