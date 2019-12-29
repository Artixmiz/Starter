package main

import "fmt"

func main() {
	defer fmt.Println("Hello World")
	var a map[int]int
	a[0] = 1
	fmt.Println(a)
}
