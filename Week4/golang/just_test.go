package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Hello World", func(t *testing.T) {
		want := 30
		got := Add(10, 20)
		if got != want {
			t.Errorf("\n got %v \n want %v")
		}
	})
}
