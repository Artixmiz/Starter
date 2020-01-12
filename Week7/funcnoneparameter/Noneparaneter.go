package main

import "fmt"

func sum(number ...int) int {
	total := 0
	for _, n := range number {
		total = total + n
	}
	return total

}

func main()  {
	x := sum(1,3,5,7,9)
	fmt.Println(x)

	y :=sum()
	fmt.Println(y)
}