package main

import "fmt"

func main() {
	fmt.Print("input :")
	var text string
	fmt.Scan(&text)
	fmt.Println(`read "`, text, `" from standard input`)
}
