package main

import "fmt"

func main() {
	for {
		fmt.Println("before")
		break
		fmt.Println("after")
	}
	fmt.Println("next staement")
}
