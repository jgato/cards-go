package main

import "fmt"

type bot interface {
	getGreetings() string
}

type spanishBot struct{}
type englishBot struct{}

func main() {
	var sp spanishBot
	var eb englishBot
	printGreetings(sp)
	printGreetings(eb)

}

func printGreetings(bot bot) {
	fmt.Print(bot.getGreetings())
}

func (spanishBot) getGreetings() string {
	return "hola tio"
}

func (englishBot) getGreetings() string {
	return "hi there"
}
