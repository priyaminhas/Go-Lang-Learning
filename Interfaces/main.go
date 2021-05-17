package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Width,Height float64
}



type Circle struct {
	Radius float64
}
type Square struct {
	Side float64
}
func (r Rectangle) Area() float64{
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64{
	return  2 * (r.Width + r.Height)
}
func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64  {
	return 2 * math.Pi * c.Radius
}
func (s Square) Area() float64  {
	return s.Side * s.Side
}
func (s Square) Perimeter() float64  {
	return 2 * (s.Side + s.Side)

}
func getArea(sh Shape){
	fmt.Println(sh.Area())
}

func getPerimeter(sh Shape){
	fmt.Println(sh.Perimeter())
}
func main()  {
	r := Rectangle{3,4}
	c := Circle{5}
	s := Square{10}

	getArea(r)
	getArea(c)
	getArea(s)

	getPerimeter(c)

	//empty interface
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "priya"
	describe(i)
}
func describe(i interface{})  {
	fmt.Printf("(%v %T)\n",i,i)
}
