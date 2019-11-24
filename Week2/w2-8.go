package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("Hello World", "hello"))
	fmt.Println(strings.HasPrefix("Hello World", "Hello"))
}
