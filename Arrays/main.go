package main

import "fmt"

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
}
