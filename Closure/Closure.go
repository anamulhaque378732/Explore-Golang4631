package main

import "fmt"

// function type , "closure"

// stack automatic clean up, but heep does not autometic clean up

const num1 = 10 // constant
var p = 120     // reassign variable , its store in the  data segment . Its alocate run time/ execution phase

func outer() func() { // outer is 1st class function
	money := 100
	age := 25
	fmt.Println("Age = ", age) //25

	show := func() { // closure
		money = money + num1 + p
		fmt.Println(money)
	}
	return show // go compiler escape analysis of the money variable  and store in the heap

}
func call() {
	incr1 := outer()
	incr1() //230
	incr1() //360

	incr2 := outer()
	incr2() // 230
	incr2() // 360
}

func main() {
	call() // invoked call function
}

func init() {
	fmt.Println("=== bank ===") // invoked first
}

// *
// when code run thats  2 phases
// 1. compilation phases/compile time(create a binary file in compiler, main file "00101001100")
// 2. execution phases / run time
//
//
// create binary file command "go build ---filename---"
// then run command "./---filename--"
//
//  *//
