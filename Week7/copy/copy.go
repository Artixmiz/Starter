package main

import "fmt"

func suntract(number int) {
	number = number - 1
}

func main() {
	x := 10
	suntract(x)
	fmt.Println(x)
}
