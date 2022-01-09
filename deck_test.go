package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	nCards := len(cardSuites) * len(cardValues)

	if len(d) != nCards {
		t.Errorf("Expected deck with %d cards but found %d", nCards, len(d))
	}
	cardN := 0
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			if d[cardN] != string(value+" of "+suite) {
				t.Errorf("(%d)Expected card %s  but found %s", cardN, string(value+" of "+suite), d[cardN])
			}
			cardN++
		}
	}
}
