package packageScope

import "fmt"

func AddSomething(x int, y int) {
	fmt.Println(x + y)
}

func Sum() {
	fmt.Println("Anamul + sumona")
}

var Money int = 100

// how to import , terminal command "go mod init example.com"
