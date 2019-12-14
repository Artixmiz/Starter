package main

import "fmt"

func main() {
	x := [6]string{"4", "3", "2", "d", "e", "f"}
	fmt.Println(x[2])
	fmt.Println(x[0:2])
	fmt.Println(len(x))
	fmt.Println(len(x), cap(x))
}
