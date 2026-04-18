package main

import "fmt"

/* Internal memory
/1. code segment //(global function, main function, init function and all function) => always read only
/
/2. data segment // global memory(global variable)
/
/3. stack (call init and allocate some space this name 'stack frame' . when call main function and allocate some space this name "main stack", then when you call any function this is which function stack)
/
/4.heap // GC => garbage collector ()
**/

var a = 10

const b = 12 //constant

var p = 100

func multiple(x int, y int) {

	fmt.Println(x * y)

}

func call() {
	var add = func(a int, b int) {

		fmt.Println("other function inside function", a+b)
	}
	add(4, 9)
	add(p, b)
}

func main() {
	// multiple(5, 6) //
	// multiple(a, 5) //
	call()
	fmt.Println(b)

}

func init() {
	fmt.Println("Hello")
}

// *
// when code run thats  2 phases
// 1. compilation phases(create a binary file in compiler, main file "00101001100")
// 2. execution phases
//
//
// create binary file command "go build ---filename---"
//then run command "./---filename--"
//
//  *//
