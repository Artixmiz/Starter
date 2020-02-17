package main

import (
	"fmt"
	"time"
)

func main() {
	data := 10
	go func() {
		data = 20

	}()
	go func() {
		fmt.Println(data)
	}()
	time.Sleep(time.Millisecond)
}
