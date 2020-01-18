package main

import (
	"fmt"
	"math/rand"
)

var rem []card = make([]card, 100)
var code []int = make([]int, 100)
var pick [53]card
var i int
var n int

type card struct {
	num  string
	suit string
}


func main() {
	pick[1] = card{"2", "Spades"}
	pick[2] = card{"2", "Clubs"}
	pick[3] = card{"2", "Diamonds"}
	pick[4] = card{"2", "Hearts"}
	pick[5] = card{"3", "Spades"}
	pick[6] = card{"3", "Clubs"}
	pick[7] = card{"3", "Diamonds"}
	pick[8] = card{"3", "Hearts"}
	pick[9] = card{"4", "Spades"}
	pick[10] = card{"4", "Clubs"}
	pick[11] = card{"4", "Diamonds"}
	pick[12] = card{"4", "Hearts"}
	pick[13] = card{"5", "Spades"}
	pick[14] = card{"5", "Clubs"}
	pick[15] = card{"5", "Diamonds"}
	pick[16] = card{"5", "Hearts"}
	pick[17] = card{"6", "Spades"}
	pick[18] = card{"6", "Clubs"}
	pick[19] = card{"6", "Diamonds"}
	pick[20] = card{"6", "Hearts"}
	pick[21] = card{"7", "Spades"}
	pick[22] = card{"7", "Clubs"}
	pick[23] = card{"7", "Diamonds"}
	pick[24] = card{"7", "Hearts"}
	pick[25] = card{"8", "Spades"}
	pick[26] = card{"8", "Clubs"}
	pick[27] = card{"8", "Diamonds"}
	pick[28] = card{"8", "Hearts"}
	pick[29] = card{"9", "Spades"}
	pick[30] = card{"9", "Clubs"}
	pick[31] = card{"9", "Diamonds"}
	pick[32] = card{"9", "Hearts"}
	pick[33] = card{"10", "Spades"}
	pick[34] = card{"10", "Clubs"}
	pick[35] = card{"10", "Diamonds"}
	pick[36] = card{"10", "Hearts"}
	pick[37] = card{"Jack", "Spades"}
	pick[38] = card{"Jack", "Clubs"}
	pick[39] = card{"Jack", "Diamonds"}
	pick[40] = card{"Jack", "Hearts"}
	pick[41] = card{"Queen", "Spades"}
	pick[42] = card{"Queen", "Clubs"}
	pick[43] = card{"Queen", "Diamonds"}
	pick[44] = card{"Queen", "Hearts"}
	pick[45] = card{"King", "Spades"}
	pick[46] = card{"King", "Clubs"}
	pick[47] = card{"King", "Diamonds"}
	pick[48] = card{"King", "Hearts"}
	pick[49] = card{"Ace", "Spades"}
	pick[50] = card{"Ace", "Clubs"}
	pick[51] = card{"Ace", "Diamonds"}
	pick[52] = card{"Ace", "Hearts"}
redo:
	fmt.Print("Choose a number (Choose 0 to exit program): ")
	fmt.Scan(&i)
	if i != 0 {
		if CheckCode(i) == 0 {
			var rem1 = pick[rand.Intn(53)]
			rem[n] = rem1
			code[n] = i
			n++
			fmt.Println(rem1)
		}
		goto redo
	}
}
func CheckCode(k int) int {

	for j := 0; j <= n; j++ {
		if code[j] == k {
			fmt.Println(rem[j])
			return 1
		}
	}
	return 0
}