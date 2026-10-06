package main

import "fmt"

// parameter vs argument

//1. first order function

/*
 //  i. Standard function or named function
 //  ii. Anonymous function
 // iii. IIFE
 //  iv. Function expression



/*
/functional paradigm -> haskel, racket math -> logic (discrete math)
/ 1. first order logic
/ 2. higher order logic
/
/
//// *** logic ****
/
/ 1. Object(people, animal, car)
/ 2. Property (color, student)
/ 3. Relation ()
/
/
/*/

// 2. Higher order function / first class function  , 3 rules must be apply any rules
/*
// 1.  Parameter => Function
// 2. Function return
// 3. Both
**/

// First order function

func addTwoNumber(a, b int) { // a,b parameter
	c := a + b
	fmt.Println(c)

}

var bb = "4444, any type" // bb = first class citizen assign a data is 1st class  citizen

var ad = addTwoNumber     // first class citizen , first class function

// function parameter

// ********   // function parameter = callback function // ******

// Higher Order function has a parameter bye a function

func processOperation(a int, b int, operation func(x int, y int)) {
	operation(a, b)
}
func add(x int, y int) {
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

// function has a parameter bye a function and tetrun a function

func arithemeticOperation(a int, b int, subs func(d int, e int)) func(x int, y int) {
	subs(a, b)
	return divide
}

func divide(a int, b int) {
	d := a / b
	fmt.Println(d)
}

func main() {

	// addTwoNumber(41, 6) // arguments => 41, 6

	// add, divide its a call back function

	// processOperation(45, 23, add)

	// sum := call()
	// sum(4, 56)
	// div := arithemeticOperation(45, 3, divide) /
	// div(45, 5)

}
