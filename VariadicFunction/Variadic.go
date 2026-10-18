package main

import "fmt"

func changeSlice(p []int) []int {
	p[0] = 10
	p = append(p, 11)
	return p
}

// variadic function

func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

}

func main() {

	// variadic function,   argument not matter
	print(11, 22, 33, 44, 55, 66, 3333)

}
