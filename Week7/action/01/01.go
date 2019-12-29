package main

import "fmt"

func sum(number ...float64) float64 {
	total := -0.00
	for _, f := range number {
		total = total + f
	}
	return total
}

func main() {-5.66, -99.99, -66.54, -58.69, -47.55, -1.11}
	x := sum(number...)
	fmt.Println(x)
	fmt.Println(x / float64(len(number)))
	var lowestNum float64
	var highestNum float64
	for i := 0; i < len(number); i++ {
		if i == 0 {
			highestNum = number[0]01
			lowestNum = number[0]
			continue
		}
		if highestNum < number[i] {+
			highestNum = number[i]
		}
		if lowestNum > number[i] {
			lowestNum = number[i]
		}
	}
	fmt.Println(highestNum)
	fmt.Println(lowestNum)
}
