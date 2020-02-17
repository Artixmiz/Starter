package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

var Handall []string

func InHand() []string {
	Handall = []string{}
	for i := 0; i < 5; i++ {
		Hand := pick()
		Handall = append(Handall, Hand)
	}
	Handall = append(Handall, pick())
	return Handall
}
func pick() string {
	num := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	suit := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	newCard := fmt.Sprint(num[rand.Intn(13)] + suit[rand.Intn(len(suit))])
	return newCard
}
func Checkcard() int {
	var pair int

	for i := 0; i < len(Handall); i++ {
		for j := i + 1; j < len(Handall); j++ {
			first := Handall[i]
			second := Handall[j]
			if strings.Contains(first[:1], second[:1]) {
				pair++
			}
		}
	}
	return pair
}
func ThrowAway() {
	for Checkcard() != 3 {
		var throw int
		fmt.Print("Choose which index to throw away: ")
		fmt.Scan(&throw)

		for i := 0; i < 6; i++ {
			if throw-1 == i {
				Handall[i] = pick()
			} else {
				Handall[i] = Handall[i]
			}
		}
		fmt.Println(Handall)
		fmt.Println(Checkcard())
	}
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}
func main() {
	var pair int
	fmt.Println(InHand())
	ThrowAway()
	fmt.Println(pair)
}
