package main

type Hand struct {
	cards    []int
	couples []int
	
func (Hand *Hand) getCard() {
	card := rand.Intn(13) + 1
	Hand.cards = append(Hand.cards,card)
}

fun main( ){
	n := rand.Intn(1)
	fmt.Println()
}