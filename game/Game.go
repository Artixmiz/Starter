package Game

import (
	"fmt"
)

// varable
var Deck [52]card

type card struct {
	num  int
	suit string
}

func main() {
	var input int
	fmt.Scan(&input)
	fmt.Print(pick(input))
}

func pick(want int) card {
	for want < 0 {
		want = want + 52
	}
	for want > 51 {
		want = want - 52
	}
	num := []string{"ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suit := []string{"Diamonds", "Hearts", "Clubs", "Spades"}

	return Deck[want]
}
