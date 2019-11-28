package main

import "fmt"

func main() {
	fmt.Println("input :")
	var name string
	var age int
	n, e := fmt.Scanln(&name, &age)
	fmt.Println("name ", name)
	fmt.Println("age ", age)
	fmt.Println("number of items successfully scanned", n)
	fmt.Println("error", e)
}
