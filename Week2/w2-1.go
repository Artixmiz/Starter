package main

import "fmt"

func main() {
	fmt.Print(`\n \t Backticks`)
	fmt.Print("\n \t double quote")
	fmt.Println()
	fmt.Println(`\n \t Backticks`[0])
	fmt.Println(len(`\n \t Backticks`))
	fmt.Println("Hello" + "World")
}
