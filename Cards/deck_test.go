package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	t.Logf("Testing newDeck")
	hand := newDeck()
	if len(hand) != 56 {
		t.Errorf("Expected Deck Length of 56, but got %v", len(hand))
	}
	if hand[0] != "One of Spades" {
		t.Errorf("Expected One of Spades to be first card in a new deck")
	}
	if hand[len(hand)-1] != "Ace of Diamonds" {
		t.Errorf("Expected Ace of Diamonds to be last card in a new deck")
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	t.Logf("Testing Save to file and New deck from file")
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	_, err := os.ReadFile("_decktesting")
	if err != nil {
		t.Errorf("Could not read deck saved to file")
	}
	d2 := newDeckFromFile("_decktesting")
	if len(d2) != 56 {
		t.Errorf("Expected deck read from file to have length of 56, but got %v", len(d2))
	}
	os.Remove("_decktesting")
}
