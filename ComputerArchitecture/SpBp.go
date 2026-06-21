package main

import "fmt"

func add(x int, y int) int {
	var result int
	result = x + y

	return result
}

// others name of stack frame in function frame

func main() {
	var a int = 10
	var sum = add(a, 4)

	fmt.Println(sum)
}
