package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// Here we are adding a method for the Student struct type
func  (s Student) getAge() int {
	return s.age
}

//pointer to be passed when mutating the values of struct
func (s *Student) setAge(age int){
	//either you can directly use s or *s, both work similarly on struct
	s.age = age
}

func main() {
	s1 := Student{"Tim", []int{70, 90, 80, 97}, 19}
	fmt.Println(s1.getAge())
	s1.setAge(45)
	fmt.Println(s1.getAge())
}