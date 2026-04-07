package main

import "fmt"

var (
	a = 100
	b = 200
	c = 300
)

func printNum(num int) {
	fmt.Println(num)
}
func addThreeNumber(x int, y int, z int) {
	res := x + y + z
	printNum(res)
}

func main() {
	addThreeNumber(a, b, c)
}

// func printNum(num int) {
// 	fmt.Println(num)
// }
