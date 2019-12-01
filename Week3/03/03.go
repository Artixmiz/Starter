package main

import (
	"fmt"
)

func main() {
	fmt.Print("import = ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	d := (a / b)
	fmt.Println(a > d)
	fmt.Printf("%V / %V > %V = %V", a, b, c, (a/b) > c)

}
