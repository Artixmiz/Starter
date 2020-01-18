package cording

import ("fmt"
		"testing"
)
func Testcording(t *Testing.T){
	testCases := []struct {
		input []int
		want int
	}{
		{[]int{-1, -2,-3,-5,-6}, -1}
	}
	for _, test := range testCases {
		got := cording(test.input[0])
		if got != test.want {
			t.Errorf("\nunexpected\n\tgot: %v\n\twant: %v", got, test.want)
		}
	}
}