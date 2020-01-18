package Game

import (
	"testing"
)

func Testpick(t *testing.T) {
	testCases := []struct {
		input int
		want  string
	}{
		{0, "ace Diamonds"},
		{1, "2 Diamonds"},
		{2, "3 Diamonds"},
		{3, "4 Diamonds"},
		{17, "5 Hearts"},
		{18, "6 Hearts"},
		{19, "7 Hearts"},
		{20, "8 Clubs"},
		{34, "9 Clubs"},
		{35, "10 Clubs"},
		{49, "J Spades"},
		{50, "Q Spades"},
		{51, "K Spades"},
		{0, "ace Diamonds"},
		{20, "8 Clubs"},
		{50, "Q Spades"},
	}
	for _, test := range testCases {
		got := pick(test.input)
		if got != test.want {
			t.Errorf("\nunexpected\n\tgot: %v\n\twant: %v", got, test.want)
		}
	}
}
