package main

import "fmt"

// type of scope- 3 type of scope
//*
// 1. global scope
// 2. local scope { block scope:   if, function,switch} , inside {} all are local scope
// 3. package scope
//
//
// *//

// global scope
var num1 = 20
var num2 = 100

// this function is global scope
func add(x int, y int) {
	// function scope/ local scope/block scope in this function
	z := x + y
	// z := x + num1 //possible , cause num1 is global scope
	// z := num2 + num1 //possible , cause num1, num2 is global scope

	// z := x + num3 //error
	fmt.Println(z)
}

// don't access main scope value in global scope, but global scope value access in global and main scope

func main() {
	// local scope in main function
	var num3 = 30
	// var num4 = 40

	// add(num1, num2) //two variable are global scope
	// add(num3, num4) //two variable are main scope
	// add(num1, num3) // one variable global others main scope
	// add(num2, num3) // one variable global others main scope
	// add(num2, num4) // one variable global others main scope

	// add(num1, z) //one value global scope, others value global function scope, so you cannot access this value

	if num3 >= 18 {
		//p is local scope
		p := 3
		fmt.Println("i am matured boy i have", p, "friend") //local scope
	}

	switch num2 {
	case 1:
		fmt.Println("this is one") //local scope
	}

}
