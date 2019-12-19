package main

import "fmt"

func main() {
	score := 75
	if score > 80 {
		fmt.Println("A")
	} else if score > 70 {
		fmt.Println("B")
	} else if score > 60 {
		fmt.Println("c")
	} else if score > 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
