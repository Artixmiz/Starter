package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[0] = 10
	fmt.Println(a, b)
	fmt.Println("---------------")
	c := []int{1, 2, 3}
	d := make([]int, len(c))
	copy(d, c)
	fmt.Println(c, d)
	c[0] = 10
	fmt.Println(c, d)
	fmt.Println("---------------")
	e := make([]int, 1)
	copy(e, c)
	fmt.Println(c, e)
}
