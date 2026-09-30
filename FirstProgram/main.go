package main

// must need package
// main function package  name must be main
import "fmt"

// fmt => format , built in package

// func = function , main function
func multiple (a int ,b int) int {
	return a * b
}

func add (a int ,b int) int {
	return a + b
}



func main() { 

	fmt.Println("Hello, world")
	fmt.Println(add( 122, 3 ))
}
