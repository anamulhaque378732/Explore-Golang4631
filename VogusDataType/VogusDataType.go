package main

import (
	"fmt"
)

func main() {

	var a int = 100 // 8 bit -128 to 127 store

	var b int16 = 165         // 16 bit -32768 to 32767
	const d int32 = 588445665 // 32 bit -2,147,483,648 to 2,147,483,647
	var c int = 35            // default 64 bit , default computer bit
	const x uint8 = 10        // unsigned 0 to and only positive number (0 to 255)
	const y = 255

	var e uint16 = 5544 // 0 to 65,535
	var r uint32 = 215554
	var t, u float32 = 25.215555, 254.326544
	var p float64 = 6556.5544
	var s, f bool = true, false

	// rune ---> alias for int32 bit (unicode point) like as symbol , single character "%c" , bangla character,

	heard := "🧡"

	fmt.Printf("%c\n", heard) // \n mean new line

	fmt.Printf("%d", a) // integer number print

	fmt.Printf("%f", p) // floating number print

	fmt.Printf("%.2f", u) // floating number print after point in 2 digit

	fmt.Printf("%v\n", s) //
	var ss string = "My name is anamul"

	fmt.Printf("%s\n", ss) // bool print

	fmt.Printf("%T\n", ss) //check type

	fmt.Println(a, b, c, x, y, e, r, t, u, p, s, f, heard)

}
