package main

import "fmt"

func main() {
	var p *int
	i := 42
	fmt.Println("value", i)
	p = &i
	*p = 3
	fmt.Println("value", i)
}
