package Create_file

import "testing"

func Test_readtxt(t *testing.T) {
	tests := []struct {
		want int
	}{
		{3000},
	}
	for _, test := range tests {

		got := len(readtxt())
		if got != test.want {
			t.Errorf("\nunexpected\n\tgot: %v\n\twant: %v", got, test.want)
		}

	}
}
