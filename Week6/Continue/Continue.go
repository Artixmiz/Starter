package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("before")
		continue
		fmt.Println("after")
	}
	fmt.Println("next statement")
}
