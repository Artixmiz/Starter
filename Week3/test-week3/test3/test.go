package main

import "fmt"

func main() {
	fmt.Print("input :")
	var Num int
	var Float float32
	var Strings string
	fmt.Scanf("%d,%f,%s", &Num, &Float, &Strings)
	fmt.Printf("= %T \n", Num)
	fmt.Printf("= %T \n", Float)
	fmt.Printf("= %T \n", Strings)
}
