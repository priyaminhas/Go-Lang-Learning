package main

import "fmt"

func sum(s []int) int{
	sum := 0
	for _, v := range s{
		sum += v
	}
	return sum
}
func sumChannel(s []int,c chan int)  {
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum
}
func main()  {
	s := []int{4,5,3,6,2}
	result := sum(s)
	//result of simple sum
	fmt.Println(result)

	//result of channel sum
	c := make(chan int)
	go sumChannel(s,c)
	x := <-c
	fmt.Println(x)
}

