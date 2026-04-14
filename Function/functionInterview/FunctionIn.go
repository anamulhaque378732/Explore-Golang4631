package main

import "fmt"

// parameter vs argument

func addTwoNumber(a, b int) { // a,b parameter
	c := a + b
	fmt.Println(c)

}

//1. first order function

/*
 //  i. standard function or named function
 //  ii. anonymous function
 // iii. IIFE
 //  iv. function expression

/*
/functional paradigm -> haskel, racket math -> logic (discrete math)
/ 1. first order logic
/ 2. higher order logic
/
/
//// *** logic ****
/
/ 1. object(people, animal, car)
/2. property (color, student)
/3. Relation ()
/
/
/*/

//2. higher order function / first class function  , 3 rules must be apply any rules
/*
// 1.  parameter => function
// 2. function return
// 3. both
**/

var bb = "4444, any type" // bb = first class citizen
var ad = addTwoNumber     // first class citizen ,first class function

// function parameter

// ********   //function parameter = callback function//   ******

func processOperation(a, b int, operation func(x int, y int)) {
	operation(a, b)
}

// function return

func call() func(x int, y int) {

	return add
}

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {

	// addTwoNumber(41, 6) // arguments => 41, 6
	processOperation(45, 23, add)

	sum := call()
	sum(4, 56)

}
