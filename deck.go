package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func newDeckFromFile(fileName string) deck {

	fileContent, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("erro: ", error)
		os.Exit(1)
	}
	return deck(strings.Split(string(fileContent), ","))

}

func (d deck) Print() {

	for _, card := range d {
		fmt.Println(card)
	}
}

// returns (hand, remainingDeck)
func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")

}

func (d deck) saveToFile(fileName string) error {
	e := ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
	return e
}
