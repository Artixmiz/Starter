package main

import "fmt"

func compute(fn func(int, int) int) int {
	return fn(3, 4)
}
func main() {
	sum := func(x, y int) int {
		return x + y
	}
	subtract := func(x, y int) int {
		return x - y
	}
	x := compute(sum)
	y := compute(subtract)
	fmt.Println(x)
	fmt.Println(y)
}
