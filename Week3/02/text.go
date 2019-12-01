package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("import = ")
	var text string
	fmt.Scan(&text)
	fmt.Print(strings.Split(text, "fuck "))

}
