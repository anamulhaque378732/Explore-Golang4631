package main

import "fmt"

// array

var arrString = [6]string{"Anamul", "Sumona", "Raihan", "Nayem", "Momin", "Nayem"}

const a int = 25 //read only. go to data segment

var arrJobs = [5]string{"Developer", "Ai", "UX/UI", "data science", "machine learning"}

var name = [4]string{"Anamul", "Sumona my love", "12"}

func main() {
	fmt.Println(a)
	// declare an array
	var arr2 = [6]int{4, 5, 6, 4, 7, 8}
	arr3 := [5]int{4, 5, 6, 7}
	var arr [6]int
	arr[1] = 6
	arr[0] = 13
	arr[5] = 3
	arr[4] = 55
	arr[5] = 12
	fmt.Println(arr)

	fmt.Println(arr2, arr3)

	fmt.Println(arrString)
	fmt.Println(arrJobs)
	fmt.Println(name[1])
}
