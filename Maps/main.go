package main

import "fmt"

func main()  {
	m := make(map[int]string)
	m[0] = "Priya"
	m[12] = "Minhas"
	fmt.Println("m:",m )

	//insert in map
	m[3] = "Manraj"
	fmt.Println("m:",m )

	//retrive from map
	elem := m[12]
	fmt.Println(elem)

	//delete from map
	delete(m,12)
	fmt.Println("m:",m )

	//check if item is present in map
	elemNew,ok := m[12]
	fmt.Printf("Element %s at 12 is %t ",elemNew,ok)

}
