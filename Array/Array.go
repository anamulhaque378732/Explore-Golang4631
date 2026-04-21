package main

import "fmt"

// array

var arrString = [5]string{"Anamul", "Sumona", "Raihan", "Nayem", "Momin"}

var arrJobs = [5]string{"Developer", "Ai", "UX/UI"}

func main() {
	// declare an array
	var arr2 = [6]int{4, 5, 6, 4, 7, 8}
	arr3 := [5]int{4, 5, 6, 7}
	var arr [6]int
	arr[1] = 6
	arr[0] = 13
	arr[5] = 3
	arr[4] = 55

	fmt.Println(arr, arr3)

	fmt.Println(arr2)

	fmt.Println(arrString)
	fmt.Println(arrJobs)
}
