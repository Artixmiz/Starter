package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}
	x := add(10, 20)
	fmt.Println(x)
}
