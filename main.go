package main

import "fmt"

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.Print()
	fmt.Println("====")
	hand, cards := deal(cards, 3)
	cards.Print()
	fmt.Println("====")
	hand.Print()
	hand.saveToFile("my_hand.txt")
	hand2 := newDeckFromFile("my_hand.txt")
	fmt.Println("****")
	hand2.Print()
}
