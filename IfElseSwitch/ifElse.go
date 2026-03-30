package main

import "fmt"

//operator >,<,>=,<=,==
// and &&
// or =||
// not= !

func main() {
	// age := 60
	// if age > 18 {
	// 	fmt.Println("You are eligible to married")
	// } else if age < 18 {
	// 	fmt.Println("you are not eligible to be married , but love someone")
	// } else if age == 18 {
	// 	fmt.Println("You are just a teenager, not eligible to be married")
	// }

	// sex := "male"
	// salary := 30000
	// if age == 20 && sex == "male" {
	// 	fmt.Println("you are ready to be married")
	// }

	// if age < 30 || salary > 40000 {
	// 	fmt.Println("You are ready to married")
	// } else {
	// 	fmt.Println("you are not ready to married")
	// }

	// isPretty := false

	// if !isPretty {
	// 	fmt.Println("you are not married to ready")
	// }

	a := 1
	switch a {
	case 1:
		fmt.Println("a is 1")

	case 2, 3:
		fmt.Println("a is either 2, 3")
	default:
		fmt.Println("a is neither 1 nor 2 or 3")
	}

}
