package main

import "fmt"

type student struct {
	fName string
	lName string
	age int
}
func (s student) fullname() {
	fmt.Printf("The full name is %s %s",s.fName,s.lName)
}
func main()  {
	s1 := student{"Manraj","Singh",3}
	s1.fullname()
}
