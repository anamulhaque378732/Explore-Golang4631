package packageScope

// obeously write a function or variable name start capital letter

import "fmt"

func AddSomething(x int, y int) {
	fmt.Println(x + y)
}

func Sum() {
	fmt.Println("Anamul + sumona")
}

func Multiple(x int, y int) {
	fmt.Println(x * y)
}

var Money int = 100

// how to import , terminal command "go mod init example.com"
