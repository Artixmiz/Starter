package main

import "fmt"

func say(greet string) func(string) string {
	return func(name string) string {
		return greet + name
	}
}

func main() {
	x := say("hello")

	fmt.Println(x("Gv"))
	fmt.Println(x("Gohan"))
}
