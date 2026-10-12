package main

import "fmt"

/* Internal memory
/1. code segment //(global function,  main function, init function and all function and declare a variable const keyword) => always read only
/
/2. data segment // global memory (global variable)
/
/3. stack (invoked init and allocate some space this name 'stack frame' . when invoked main function and allocate some space this name "main stack", then when you invoked any function this is which function stack)
/ when invoked a function its alocate a space in stack its call stack frame
/4. heap // GC => garbage collector ()
**/

var a = 10 // reassign this variable . its store data segment

const b = 12 //constant. you dont reassign this

var p = 100

func multiple(x int, y int) {

	fmt.Println(x*y, p)

}

func call() {
	var add = func(a int, b int) {

		fmt.Println("Other function inside function", a+b)
	}
	add(4, 9)
	add(p, b)
}

func main() {
	// multiple(5, 6)
	// multiple(a, 5)
	call()
	fmt.Println(b)

}

func init() {
	fmt.Println("Hello")
}

// *
// when code run thats  2 phases
// 1. compilation phases (create a binary file in compiler, main file "00101001100")
// 2. execution phases
//
//
// create binary file command "go build ---filename---"
// then run command "./---filename--"
//
//  *//
