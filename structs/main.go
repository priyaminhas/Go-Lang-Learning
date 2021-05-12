package main

import "fmt"

type student struct {
	age int
	fName string
}

func main()  {
	fmt.Println("main function start")
	var a student
	//setting the values of struct from an external function
	a.setStudent(10,"Manraj")
	//as the age is set by value so its not changed below
	fmt.Println(a)
	aPoint := &a
	aPoint.setStudent(4,"Manraj")
	fmt.Println(a)

	//pointer to struct
	var b student
	p := &b
	p.age = 3
	p.fName = "sauhraab"
	fmt.Println(b)
}
func (a student) setStudent(x int, y string) student{
	a.age = x
	a.fName = y
	return a
}
func (aPoint *student) setStudentByReference(x int, y string) {
	aPoint.age = x
	aPoint.fName = y
}
