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
	fmt.Println(len(numbers)) //7
	fmt.Println(cap(numbers)) //7

}

func main() {

	// slice new underlying array rule => 1024 -> incrage 100% . then insert a element increage 1025 to 2048. then can   increage 25% bye old array length.

	var slice []int          // [] length = 0 cap =  0
	slice = append(slice, 1) // [1] len = 1  cap = 1
	slice = append(slice, 2) // [1,2] len = 2  cap = 2
	slice = append(slice, 3) // [1,2,3] len = 3  cap = 4
	// slice = append(slice, 4)  // [1,2,3,4] len = 4  cap = 4
	slice2 := slice            // [1,2,3,4] len = 3  cap = 4
	slice = append(slice, 5)   // [1,2,3,5] len = 4  cap = 4
	slice2 = append(slice2, 6) // [1,2,3,6] len = 4  cap = 4
	slice[0] = 10              // [10,2,3,6] len = 4  cap = 4

	// fmt.Println(slice, len(slice), cap(slice), slice2, len(slice2), cap(slice2))

	// fmt.Println(len(slice), cap(slice))

	// another example bye slice

	x := []int{1, 2, 3, 4, 5}

	x = append(x, 6)
	x = append(x, 7)
	a := x[4:]

	y := changeSlice(a)

	// fmt.Println(x, len(x), cap(x)) // [1 2 3 4 10 6 7]  len = 7  cap = 10

	// fmt.Println(a, len(a), cap(a)) // [10 6 7]  len = 3  cap = 6

	fmt.Println(y, len(y), cap(y)) //  [10 6 7 11]  len = 4  cap = 6

	fmt.Println(x[0:7])  //[1 2 3 4 10 6 7]
	fmt.Println(x[0:8])  //[1 2 3 4 10 6 7 11]
	fmt.Println(x[0:9])  //[1 2 3 4 10 6 7 11 0]
	fmt.Println(x[0:10]) // [1 2 3 4 10 6 7 11 0 0]
	// fmt.Println(x[0:20]) //  runtime error: slice bounds out of range [:20] with capacity 10

	// variadic function,   argument not matter

	print(11, 22, 33, 44, 55, 66, 33, 33)

}
