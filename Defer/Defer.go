// Defer

package main

import "fmt"

func a() {
	i := 0 // 0
	fmt.Println("First", i)

	defer fmt.Println("Second", i) // 0 call  go run time   // defer function store linked list,
	i = i + 1
	fmt.Println("Third", i)        //1
	defer fmt.Println("Fourth", i) //1
}

//1.  named return value

func sum(a int, b int) (Total int) {
	Total = a + b
	return
}

func calculate() (result int) {
	fmt.Println("First", result)
	defer func() {
		result = result + 10
		fmt.Println("defer", result)
	}()
	result = 5
	fmt.Println("Second", result)
	return
}

func calc() int {
	result := 0
	fmt.Println("First2", result)
	defer func() {
		result = result + 10
		fmt.Println("defer2", result)
	}()
	result = 5
	fmt.Println("Second2", result)
	return result
}

func main() {
	a := calculate()
	fmt.Println("main first", a)

	b := calc()
	fmt.Println("main Second", b)

	// res := sum(4, 5)

	// fmt.Println("Total", res)

	// fmt.Println(sum(4, 41))
	// a()
}
