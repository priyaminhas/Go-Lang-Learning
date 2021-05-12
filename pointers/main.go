package main

import "fmt"

func main()  {
	fmt.Println("main function")
	i := 10
	p := &i
	//reads i through p
	fmt.Println(*p)
	//set by value call
	setByValue(10)
	fmt.Println(i)

	//set by reference call
	setByReference(&i)
	fmt.Println(i)
}
func setByValue(i int)   {
	i = 20
}
func setByReference(p *int ){
	*p = 20
}
