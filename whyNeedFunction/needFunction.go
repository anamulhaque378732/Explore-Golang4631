package main

import "fmt"

// SOLID =
/*
S=Single responsibility
O=

*/

func printWelcomeMessage() {
	fmt.Println("Welcome to the application")
}

func getUserName() string {
	var name string = ""
	fmt.Println("Enter your name -")
	fmt.Scanln(&name)
	return name
}

func getToNumber() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter first number -")
	fmt.Scanln(&num1) //&=ampersand
	fmt.Println("Enter your second number -")
	fmt.Scanln(&num2)
	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func display(name string, sum int) {
	fmt.Println("Hello", name)
	fmt.Println("summation", sum)
}

func goodByeMessage() {
	fmt.Println("thank you using the application")
	fmt.Println("Good bye")
}

func main() {
	// print welcome message
	printWelcomeMessage()

	// print user name
	name := getUserName()
	num1, num2 := getToNumber()
	sum := add(num1, num2)

	// display result
	display(name, sum)
	goodByeMessage()
}

// func main() {
// 	// print welcome to the message

// 	fmt.Println("Welcome to the application")

// 	// get user name as input

// 	var name string = ""
// 	fmt.Println("Enter your name -")

// 	fmt.Scanln(&name)

// 	// fmt.Println("------", name)

// 	var num1 int
// 	var num2 int

// 	fmt.Println("Enter first number -")
// 	fmt.Scanln(&num1) //&=ampersand
// 	fmt.Println("Enter your second number -")
// 	fmt.Scanln(&num2)

// 	sum := num1 + num2

// 	// display result

// 	fmt.Println("Hello ", name)
// 	fmt.Println("summation =", sum)

// 	// print good bye

// 	fmt.Println("thank you using the application")
// 	fmt.Println("Good bye")

// }
