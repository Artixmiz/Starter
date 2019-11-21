package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Hello World", "hi"))
	fmt.Println(strings.ContainsAny("Hello World", "Hi"))
}
