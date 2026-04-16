package main

import "fmt"

/* Internal memory
/1. code segment //(global function, main function, init function and all function)
/
/2. data segment // global memory(global variable)
/
/3. stack (call init and allocate some space this name 'stack frame' . when call main function and allocate some space this name "main stack", then when you call any function this is witch function stack)
/
/4.heap // GC => garbage collector ()
**/

var a = 10

func multiple(x int, y int) {

	fmt.Println(x * y)

}

func main() {
	multiple(5, 6) //
	multiple(a, 5) //
}

func init() {
	fmt.Println("Hello")
}
