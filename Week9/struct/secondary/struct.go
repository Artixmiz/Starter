package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func main() {
	var a student
	a.name = "goku"
	a.age = 18
	a.email = "Goku@super.saiya"
	b := student{"gohan", 3, "gohan@super.saiya"}
	c := student{name: "videl", email: "videl@dauhter.satan"}
	d := student{age: 20}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
