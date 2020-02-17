package main

import (
	"fmt"
	"time"
)

func say(txt string) {
	for i := 0; i < 3; i++ {
		fmt.Println(time.now(), ":", i, ":", txt)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	go say("hello")
	go say("hi")
}
