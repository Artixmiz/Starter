package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("import = ")
	var text string
	fmt.Scan(&text)
	fmt.Println("a", strings.Count(text, "a"))

}
