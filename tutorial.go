package main

import "fmt"

type Point struct {
	x int32
	y int32
}


func changeX(pt Point)  {
	pt.x = 100
}

func main() {
	p1 := Point{1, 2}
	fmt.Println(p1)
	changeX(p1)
	fmt.Println(p1)


}