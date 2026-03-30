package main

import "fmt"

// data type

/*
1. numeric <1......n>
1.1- integer number 10 <int> int, int8, int16, int32,int64, uint,uint8,uint16,uint32, uint64

1.2-floating point number 10.5,13.124 <float> float32,or float64
1.3-


2.boolean <true,false> <bool>
3.string <"a...z"> <string>
4.

*/

func main() {
	//**
	// int
	// float
	// bool
	// string
	// */

	// variable declaration in a block
	var (
		aaa int
		baa int    = 1
		caa string = "Hello"
	)

	// declare a initial variable
	var x int = 10
	var y float32 = 30.35
	var age = 10
	var isGo = true
	var str = "Anamul"
	var myName string
	myName = "anamul"
	a := 10
	aa := true
	aa = false
	const num = 1100
	fmt.Println(aa, x, y, age, isGo, str, num, a, myName, aaa, baa, caa)
	a = 100

}

var initial int = 15

// declare multiple variable
var a, b, c, d int = 1, 2, 3, 4

const MYNAME = "Sumona"
