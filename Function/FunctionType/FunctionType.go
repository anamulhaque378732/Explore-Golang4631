package main

import "fmt"

var (
	a = 10
)

// function type

// 1. Standard function or named function
func add() {
	fmt.Println(4 + 16)
}

// 2. Anonymous function

// func(a int, b int) {
// 	c := a + b
// 	fmt.Println(c)
// }(4, 6)

// 3. function expression or Assign function in variable

// 4. Higher order function or first class function

// 5. callback function

// 6. variadic function

// 8. init function - you can not call this, computer call this automatically (at first call init function then  main then global and  others function)

func init() {
	fmt.Println("i am the first function that executed first")
	fmt.Println(a)
	a = 20

}

// 9. closure-close over

// 10. Defer function

// 11. receiver function

// 12. IIFE- immediately invoked function expression
//

func main() {
	// add()
	// fmt.Println(a)

	// anonymous function
	// immediately invoked function expression,IIFE

	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(4, 63)

}
