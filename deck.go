package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "clubs"}
	cardValues := []string{"One", "Two", "Three"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards

}
func (d deck) Print() {

	for _, card := range d {
		fmt.Println(card)
	}
}
