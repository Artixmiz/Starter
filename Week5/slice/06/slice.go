package main

import "fmt"

func main() {
	hunmans := []string{"bulma", "ChiChi", "Videl"}
	names := []string{"goku", "gohab"}
	names = append(names, hunmans...)
	fmt.Println(names)
}
