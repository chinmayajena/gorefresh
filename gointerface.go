package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main()  {
	//:= is defining without var keyword and works inside method.
	//cannot be used in package level variabel declaration
	c1:=circle{4}
	r1:=rect{5,7}
	shapes:=[]shape{c1,r1}

	//here _ is index. Because index will be not be used, golang asks user to use _
	for _, shape:=range shapes {
		fmt.Println(shape.area())
	}
}