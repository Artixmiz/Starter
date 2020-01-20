package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func main() {
	std := student{name: "goku"}
	p := &std
	(*p).age = 18
	p.email = "goli@super.saiya"

	fmt.Println(std)
}
