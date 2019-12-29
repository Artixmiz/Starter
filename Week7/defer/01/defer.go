package main

import "fmt"

func say() {
	fmt.Println("Hi Alina")
}

func main() {
	defer say()
	fmt.Println("Hello World")
}
