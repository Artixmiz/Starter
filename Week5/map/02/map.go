package main

import "fmt"

func main() {
	x := make(map[string]int)
	fmt.Println(len(x))
	x["key"] = 99
	fmt.Println(len(x))
	fmt.Println(x)

	y := map[string]int{"key1": 10, "key2": 20}
	y["key3"] = 30
	fmt.Println(y)
}
