package main

import "fmt"

func say() {
	fmt.Println("Hello")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func main() {
	say()
	greet("GOKU")
}
