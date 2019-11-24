package main

import (
	"fmt"
)

func main() {
	fmt.Println("10 is of type %T\n", 10)
	fmt.Println("10 is of type %T\n", string(10))
	fmt.Println("10 is of type %T\n", float64(10))
}
