package main

import (
	"fmt"
	"gameCard"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var cardDeck gameCard.Deck

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
			cardDeck = gameCard.MakeRandomizedSimpleDeck(cardCount, "Randomized Simple")
			fmt.Println(cardDeck)
		}
	} else {
		deckConfig := gameCard.DeckConfig{
			TotalCards:    18,
			NumBuffCards:  6,
			NumNerfCards:  6,
			NumScoreCards: 6,
		}
		cardDeck = gameCard.MakeConfiguredDeck(deckConfig, "Configured")
		// fmt.Println(cardDeck)
	}

	http.HandleFunc("/", deckHandler)
	http.HandleFunc("/new", newDeckHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func deckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test %s", html.EscapeString(r.URL.Path))
}

func newDeckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div><div><span>New Deck</span></div>")
	var defaultName = "Default"
	q := r.URL.Query()
	size := q["size"]
	name := q["name"]
	fmt.Printf("Len Size = %v", len(size[0]))

	if size != nil && len(size[0]) < 1 {
		fmt.Fprintf(w, "<div><span>Error : got size %v</span></div>", html.EscapeString(size[0]))
		return
	}

	if len(name) < 1 || len(name[0]) < 1 {
		fmt.Fprintf(w, "<div><span>No deck name supplied, using default</span></div>")
	}

	fmt.Printf("Name = %v", name)

	deckSize, err := strconv.Atoi(size[0])
	if err != nil {
		fmt.Fprintf(w, "Error converting %v to int", html.EscapeString(size[0]))
	} else {
		deckConfig := gameCard.DeckConfig{
			TotalCards:    deckSize,
			NumBuffCards:  6,
			NumNerfCards:  6,
			NumScoreCards: 6,
		}
		if len(name) > 0 {
			cardDeck = gameCard.MakeConfiguredDeck(deckConfig, name[0])
		} else {
			cardDeck = gameCard.MakeConfiguredDeck(deckConfig, defaultName)
		}
	}

	fmt.Fprintln(w)
	fmt.Fprintf(w, "<div><span>Deck : %v</span></div>", cardDeck.Name)
	count := 0
	for _, card := range cardDeck.Cards {
		count++
		fmt.Fprintln(w)
		fmt.Fprintf(w, "<div><span>Card %v: %v</span></div>", count, card)
	}
}
