package main

import "fmt"

// function type , "closure"

const num1 = 10 // constant
var p = 100

func outer() func() {
	money := 100
	age := 25
	fmt.Println("Age = ", age)

	show := func() { //closure
		money = money + num1 + p
		fmt.Println(money)
	}
	return show

}
func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()

	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("===bank===")
}

// *
// when code run thats  2 phases
// 1. compilation phases/compile time(create a binary file in compiler, main file "00101001100")
// 2. execution phases/ run time
//
//
// create binary file command "go build ---filename---"
//then run command "./---filename--"
//
//  *//
