package main

import (
	"fmt"
)

func main()  {
	//simple array
	var a [4]string
	a[0] = "Priya"
	a[1] = "Minhas"
	a[2] = "Manraj"
	a[3] = "Singh"
	fmt.Println(a)

	//slice
	a1 := a[0:2]
	a2 := a[2:4]
	fmt.Println(a1,a2)

	//changing the slices make changes to array
	a2[0] = "Mandeepak"
	fmt.Println(a)

	//slice length and capacity
	s := []int{45,56,24,778,45}
	printSlice(s)

	//slicing the slice
	s = s[:3]
	printSlice(s)

	s = s[2:5]
	printSlice(s)

	//making slice with make function
	b := make([]int,0,5)
	printSlice(b)

	//append to slice
	b = append(b,12,34)
	b = append(b,121,234,756)
	printSlice(b)

	for i,v:= range b {
		fmt.Printf("b[%d] = %d  \n",i,v)
	}
}

func printSlice(s []int){
	fmt.Printf("length=%d, capacity=%d,  %v \n",len(s),cap(s),s)
}
