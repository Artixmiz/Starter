package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("Hello World", "hello"))
	fmt.Println(strings.HasSuffix("Hello World", "Hello"))
}
