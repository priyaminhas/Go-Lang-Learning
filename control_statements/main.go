package main

import (
	"fmt"
	"math"
	"runtime"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	j := 5.0
	//if else statements
	if v:= j + math.Pi ; v <20 {
		fmt.Println("v is less than 20")
	} else{
		fmt.Println("v is greater than 20")
	}

	//switch statements
	switch os:= runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)

	}
	//defer statement
	for i := 1; i < 10 ; i++{
		defer fmt.Println(i)
	}
	fmt.Println("done main function")
}
