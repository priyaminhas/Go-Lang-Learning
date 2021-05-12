package main

import (
	"fmt"
	"math/cmplx"
)

//variables without specify the type as its implicit with assignment
var i, j = 1, 2
//all basic types of go, factored into a block
var (
	boolVar bool
	stringVar string
	z  complex128 = cmplx.Sqrt(-5 + 12i)
	emptyVar int
)
//all int int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//byte => alias for uint8

//rune => alias for int32
// represents a Unicode code point

//float32 float64

//complex64 complex128

func main() {
	var c, python, java = true, false, "no!"
	//inside function short assignment is possible
	k := 100
	fmt.Println(i, j, c, python, java,k)
	const Pi = 3.14

	boolVar =  true
	stringVar = "Priya this is first string"
	fmt.Println(boolVar)
	fmt.Println(stringVar)
	fmt.Println(z)
	//not assigned any value
	fmt.Println(emptyVar)
	//constant
	fmt.Println("Hello World, I am a constant ",Pi)
}