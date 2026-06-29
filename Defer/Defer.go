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

/*
* ----Named return values function (rules)---
* 1. All codes execute
* 2. defer function store a magic box name *******"Defer list pointer"******
* 3. when return --> all defer function complete  before return
* 4. return maned variables values
*
* ----Just return type---
* 1. All codes execute
* 2. defer function store
* 3. return values are evaluated at this time (store the return values)
* 4. all defer function execute and complete
*
 */

//1.  named return value

func sum(a int, b int) (Total int) {
	Total = a + b
	return
}

func calculate() (result int) { //named return
	fmt.Println("First", result) //0

	defer func() {
		result = result + 10
		fmt.Println("defer", result)
	}() //15

	defer func(a int) {
		fmt.Println("ami", a)
	}(result) // 5

	defer fmt.Println("third", 5) //5

	result = 5
	fmt.Println("Second", result)

	return //15
}
func calculate2() (result int) { //named return
	fmt.Println("First", result) //0

	defer func() {
		result = result + 10
		fmt.Println("defer", result)
	}() //15

	defer func(a int) {
		fmt.Println("ami", a)
	}(result) // 0  Arguments to a deferred function call are evaluated immediately when the defer statement is executed, not when the deferred function runs.

	defer fmt.Println("third", 5) //5

	result = 5
	fmt.Println("Second", result)

	return //15
}

var x = 10

// func calc() int { // type return
// 	result := 0
// 	fmt.Println("First2", result)

// 	defer func() { //store another space with closure form
// 		result = result + 10
// 		fmt.Println("defer2", result)
// 	}()

// 	result = 5

// 	fmt.Println("Second2", result)
// 	return result
// }

func main() {
	a := calculate()
	fmt.Println("main first", a) //15

	// b := calc()
	// fmt.Println("main Second", b) //5

	// res := sum(4, 5)

	// fmt.Println("Total", res)

	// fmt.Println(sum(4, 41))
	// a()
}
