package main

import "fmt"

// most question ask  interview in this topic *** slice ***

func main() {
	// Declare slice way

	var arr = [6]string{"This", "is", "a", "Go", "Interview", "questins"}

	// // slice has 3 parameter  pointer, length, capasity

	// way- 1. slice from in existing array

	slice := arr[1:4] // ["is", "a", "Go"] //  , length = 3, capasity = 5

	slice2 := arr[3:5] //    length = 2, capasity = 3

	//  Way - 2.  Slice to slice, when you make slice to slice . subslice use highest capasity of main slice

	slice3 := slice[1:2] // ["a"] //    length = 1, capasity = 4
	slice4 := slice[1:5] //["a" ,"Go", "Interview" ,"questins"]  length = 4, capasity = 4

	fmt.Println(arr)
	fmt.Println(slice, slice2, slice3, slice4)
	fmt.Println(len(slice), len(slice3), len(slice4))

	// Way - 3.  slice literal : // when you cann't fixed the array length it's called slice literal. when you declare a slice bye slice literal then length and capasity are same

	sliceInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //length : 9 , capasity : 9

	fmt.Println("slice", sliceInt, "len", len(sliceInt), "Cap", cap(sliceInt))

	//  Way - 4 slice bye make  function with length

	sliceByeMakeFunction := make([]int, 3) // [0,0,0] length = 3 cap = 3
	sliceByeMakeFunction[0] = 5            // [5,0,0] length = 3 cap = 3

	fmt.Println(sliceByeMakeFunction, len(sliceByeMakeFunction), cap(sliceByeMakeFunction))

	//  Way - 5 slice bye make function with length and capasity

	anotherSliceByeMakeFunction := make([]int, 3, 5) //length = 3 , cap = 5 [0,0,0]

	anotherSliceByeMakeFunction[2] = 10 //length = 3 , cap = 5 [0,0,10]

	anotherSliceByeMakeFunction = append(anotherSliceByeMakeFunction, 10) //[0,0,10,10] length = 4 , cap = 5

	fmt.Println(anotherSliceByeMakeFunction, len(anotherSliceByeMakeFunction), cap(anotherSliceByeMakeFunction))

	//  Way -  6 Empty slice or nil slice

	var slice6 []int //length = 0 , cap = 0

	slice6 = append(slice6, 4)             //len = 1, cap = 1
	slice6 = append(slice6, 44)            //len = 1, cap = 1
	slice6 = append(slice6, 121, 232, 343) // length = 3 , cap = 3
	// fmt.Println(slice6, len(slice6), cap(slice6), apend)
	fmt.Println(slice6)

}
