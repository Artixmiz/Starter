package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomArray(len int) []int {
	a := make([]int, len)
	for i := 0; i <= len-1; i++ {
		a[i] = rand.Intn(len)
	}
	return a
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randomArray(5))
	randomNum := random(1, 13)
	fmt.Printf("Random Num: %d\n", randomNum)

}
