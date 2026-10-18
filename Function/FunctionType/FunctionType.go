package main

import "fmt"

var (
	a = 10
)

// custome type , receiver function use

type User struct { // read only. store code segment. compile time , code segment
	Name    string // member variable or property
	Age     int    // member variable or property
	Country string
	Id      int
}

// function type

// 1. Standard function or named function

func add(a int, b int) int {
	return a + b
}

// 2. Anonymous function

// func(a int, b int) {
// 	c := a + b
// 	fmt.Println(c)
// }

// 3. function expression or Assign function in a variable

var addTwoNumber = func(a int, b int) int {
	c := a + b
	return c
}

// 4. Higher order function or first class function

// Higher Order function has a parameter bye a function

func processOperation(a int, b int, operation func(x int, y int)) {
	operation(a, b)
}
func addSomething(x int, y int) {
	z := x + y
	fmt.Println(z)
}

// Higher order function  return  a function

func call() func(x int, y int) {
	return multipleTwoNumbwer
}

func multipleTwoNumbwer(a int, b int) {
	c := a * b
	fmt.Println(c)
}

// 5. callback function

// 6. variadic function

func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

}

// 7. init function - you can not invoked this, computer  invoked this automatically (at first  invoked init function then  main then global and  others function)

func init() {
	fmt.Println("I am the first function that is executed first")
	fmt.Println(a) //10
	a = 20
}

// 8. closure-close over funtion

// 9. Defer function

// 10. receiver function

// receiver function (you can use when you create custom type)

func (usr User) printDetails() {
	fmt.Println("Name : ", usr.Name)
	fmt.Println("Age : ", usr.Age)
}

func (usr User) call(a int) {
	fmt.Println("Name : ", usr.Name)
	fmt.Println(a)
}

// 11.  IIFE - immediately invoked function expression

// func(a int, b int) {
// 		c := a + b
// 		fmt.Println(c)
// 	}(4, 63)

func main() {
	// add()
	// fmt.Println(a)

	// In local scope, if a function is defined below, you cannot call it from above. But in global scope, you can call a function before its declaration.

	// // this function is not local scope bye main function, this function is global scope function

	fmt.Println(addTwoNumber(4, 8))

	// local scope , and shadowing

	// var addTwoNumber = func(a int, b int) {
	// 	c := a + b
	// 	fmt.Println(c)
	// }

	addTwoNumber(4, 5)

	// processOperation(45, 23, addSomething)

	// sum := call()
	// sum(4, 56)

	// variadic func,   argument not matter
	print(11, 22, 33, 44, 55, 66, 3333)

}
