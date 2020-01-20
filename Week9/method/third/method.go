package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func (std student) growUp(i int) {
	std.age = std.age + i
}

func main() {
	var a student
	a.age = 18
	fmt.Println(a.age)
	a.growUp(20)
	fmt.Println(a.age)
}
