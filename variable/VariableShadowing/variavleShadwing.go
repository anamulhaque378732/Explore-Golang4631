package main

import "fmt"

var (
	a    int    = 10
	name string = "anamul"
)

func main() {
	age := 30

	if age > 18 {
		a := 46
		fmt.Println(a) //47

	} else {

		fmt.Println(a + 20)
	}

	fmt.Println(a) //10
}
