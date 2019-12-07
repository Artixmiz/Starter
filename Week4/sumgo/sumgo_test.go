package sumgo

import "testing"

func TestSumgo(t *testing.T) {

	want := 1
	got := plus(1, 1)
	if got != want {
		t.Error("error1+1")
	}
}
