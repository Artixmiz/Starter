package main

import "fmt"

func main() {
	panic("Hello panic")
	text := recover()
	fmt.Println(text)
}
