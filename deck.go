package main

import (
	"fmt"
	"gameCard"
	"os"
	"strconv"
	"strings"
)

func main() {
	// The os.Args always contains the path to the executable so program args are at index 1+
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	// If we have an arg and it matches the min length, use it
	if args != nil {
		var valueToConvert = args[0]
		fmt.Println("valToConvert", valueToConvert)

		if cardCount, err := strconv.Atoi(strings.Replace(valueToConvert, " ", "", -1)); err != nil {
			fmt.Printf("Error 1: %+v", err)
		} else {
			cardDeck := gameCard.MakeRandomizedSimpleDeck(cardCount, "Randomized Simple")
			fmt.Println(cardDeck)
		}
	} else {
		deckConfig := gameCard.DeckConfig{
			TotalCards:    18,
			NumBuffCards:  6,
			NumNerfCards:  6,
			NumScoreCards: 6,
		}
		cardDeck := gameCard.MakeConfiguredDeck(deckConfig, "Configured")
		fmt.Println(cardDeck)
	}

}
