package main

import "testing"

func TestFind(t *testing.T) {
	t.Run("Hello World", func(t *testing.T) {
		want := true
		got := Find("Hello World")
		if got != want {
			t.Errorf("\n got %v want %v", got, want)
		}
	})

	t.Run("Hello Solar System", func(t *testing.T) {
		want := false
		got := Find("Hello Solar System")
		if got != want {
			t.Errorf("\n got %v want %v", got, want)
		}
	})

	t.Run("Hello Lunar System", func(t *testing.T) {
		want := false
		got := Find("Hello Lunar System")
		if got != want {
			t.Errorf("\n got %v want %v", got, want)
		}
	})

	t.Run("Hello We Are Cs", func(t *testing.T) {
		want := true
		got := Find("Hello We Are Cs")
		if got != want {
			t.Errorf("\n got %v want %v", got, want)
		}
	})
}
