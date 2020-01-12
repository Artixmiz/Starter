package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	x, y := swap(10, 5)
	fmt.Println(x, y)
}
