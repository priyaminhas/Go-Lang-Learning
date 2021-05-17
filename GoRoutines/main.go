package main

import "fmt"

func main()  {
	go routineWork("Going to Park")
	routineWork("Making Bed")
}
func routineWork(s string){
	fmt.Printf("This is my work : %s",s)
}
