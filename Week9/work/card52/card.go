package main

import "fmt"

type deck struct {
	suit string
	num  string
}

func main() {
	var arrnum [10000]deck
	arrnum[0] = deck{"spades", "A"}
	arrnum[1] = deck{"spades", "2"}
	arrnum[2] = deck{"spades", "3"}
	arrnum[3] = deck{"spades", "4"}
	arrnum[4] = deck{"spades", "5"}
	arrnum[5] = deck{"spades", "6"}
	arrnum[6] = deck{"spades", "7"}
	arrnum[7] = deck{"spades", "8"}
	arrnum[8] = deck{"spades", "9"}
	arrnum[9] = deck{"spades", "10"}
	arrnum[10] = deck{"spades", "J"}
	arrnum[11] = deck{"spades", "Q"}
	arrnum[12] = deck{"spades", "K"}
	fmt.Println(arrnum[9999])

}
