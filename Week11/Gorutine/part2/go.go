package main

import "fmt"

func say(txt string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, ":", txt)
	}
}

func main() {
	go say("hello")
	go say("hi")
	var input string
	fmt.Sprintln(&input)
}
