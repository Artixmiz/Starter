package main

import "fmt"

func main() {
	n, e := fmt.Print("Hello", "World", 123, 456)
	fmt.Print("number of bytes weitten :", n)
	fmt.Print("write error encountered :", e)
}
