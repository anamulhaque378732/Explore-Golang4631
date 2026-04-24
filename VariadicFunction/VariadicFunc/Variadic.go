package main

import "fmt"

func changeSlice(p []int) []int {
	p[0] = 10
	p = append(p, 11)
	return p
}

// variadic function

func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

}

func main() {
	// ------------------------1--------------------------
	// var x []int //[], len = 0 cap = 0

	// x = append(x, 1) //[1], len = 1 cap = 1
	// x = append(x, 2) //[1,2], len = 2 cap = 2
	// x = append(x, 3) //[1,2,3], len = 3 cap = 4

	// //slice  underlying array rule => 1024 --> increase 100% . 1024 < 1025 to n,  then increase 25%

	// x = append(x, 4) //[1,2,3], len = 4 cap = 4

	// y := x // [1,2,3], len = 3 cap = 4
	// x = append(x, 5)

	// // fmt.Println(x, y) //[1,2,3,5] [1,2,3]

	// y = append(y, 6)

	// x[0] = 10

	// fmt.Println(x) //[10,2,3,6]
	// fmt.Println(y) //[10,2,3,6]

	// x = append(x, 8)
	// y = append(y, 9)
	// fmt.Println(x, y) //[1,2,3,6,8] [1,2,3,6,9]

	// -----------------2--------------------------

	x := []int{1, 2, 3, 4, 5}

	x = append(x, 6)

	// fmt.Println(x) //[1,2,3,4,5,6]

	x = append(x, 7)

	// fmt.Println(x) //[1,2,3,4,5,6,7]

	a := x[4:]

	// fmt.Println(a) // [5,6,7] len=3 cap =6

	y := changeSlice(a) //

	// fmt.Println(x)      //[1,2,3,4,10,6,7]
	fmt.Println(y) //[10,6,7,11]
	// fmt.Println(x[0:8]) //[1 2 3 4 10 6 7 11]

	// fmt.Println(x[0:9]) //[1 2 3 4 10 6 7 11 0]

	// fmt.Println(x[0:10]) //[1 2 3 4 10 6 7 11 0 0]

	// fmt.Println(x[0:11]) //runtime  error

	// variadic func,   argument not matter
	print(11, 22, 33, 44, 55, 66, 3333)

}
