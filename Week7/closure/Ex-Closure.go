package main

import "fmt"

func makeEven() func() int {
	even := 0
	return func() int {
		even = even + 2
		return even
	}
}

func main() {
	nextEven := makeEven()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
