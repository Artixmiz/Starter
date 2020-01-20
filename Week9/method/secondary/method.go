package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func (std student) introduce() {
	fmt.Println("Hello my name is", std.name)
}

type pupil struct {
	address string
	std     student
}

func main() {
	goku := student{name: "goku"}
	pup := pupil{std: goku}
	pup.std.introduce()
}
