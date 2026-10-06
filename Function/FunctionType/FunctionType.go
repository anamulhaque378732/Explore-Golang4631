package main

import "fmt"

var (
	a = 10
)

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

// 5. callback function

// 6. variadic function

// 8. init function - you can not invoked this, computer  invoked this automatically (at first  invoked init function then  main then global and  others function)

func init() {
	fmt.Println("I am the first function that is executed first")
	fmt.Println(a) //10
	a = 20
}

// 9. closure-close over funtion

// 10. Defer function

// 11. receiver function

// 12.  IIFE - immediately invoked function expression

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

}
