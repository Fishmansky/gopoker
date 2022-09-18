package croupier

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var HandsOrder = map[string]Hand{
	"Royal Flush":    {Cards: []string{"10.*", "J.*", "Q.*", "K.*", "A.*"}},
	"Straight Flush": {Cards: []string{"X.C", "X+1.C", "X+2.C", "X+3.C", "X+4.C"}}, // X for figure, C for Color
	"Four of kind":   {Cards: []string{"X.*", "X.*", "X.*", "X.*", "*.*"}},
	"Full House":     {Cards: []string{"X.*", "X.*", "X.*", "Y.*", "Y.*"}},
	"Flush":          {Cards: []string{"*.C", "*.C", "*.C", "*.C", "*.C"}},
	"Streigh":        {Cards: []string{"X.*", "X+1.*", "X+2.*", "X+3.*", "X+4.*"}},
	"Three of kind":  {Cards: []string{"X.*", "X.*", "X.*", "Y.*", "Z.*"}},
	"Two pairs":      {Cards: []string{"X.*", "X.*", "Y.*", "Y.*", "Z.*"}},
	"Pair":           {Cards: []string{"X.*", "X.*", "Y.*", "Z.*", "V.*"}},
	"High Card":      {Cards: []string{"X.*", "Y.*", "Z.*", "V.*", "Q.*"}},
}

type Table struct {
	Deck           []string
	CommunityCards []string
}

type Hand struct {
	PlayerName string
	Cards      []string
}

func Rank(card string) string {
	return strings.Split(card, ".")[0]
}

func Suit(card string) string {
	return strings.Split(card, ".")[1]
}

func CardValues(cards []string) []int {
	convertedCards := []int{}
	for _, card := range cards {
		var cardVal int
		switch Rank(card) {
		case "J":
			cardVal = 11
		case "Q":
			cardVal = 12
		case "K":
			cardVal = 13
		case "A":
			cardVal = 14
		default:
			intVal, err := strconv.Atoi(Rank(card))
			if err != nil {
				panic(err)
			}
			cardVal = intVal
		}
		convertedCards = append(convertedCards, cardVal)
	}
	return convertedCards
}

func GroupCards(cards *[]string) []string {
	return []string{}
}

func (t *Table) EvaluateHand(h *Hand) string {
	// get cards from hand and table
	allCards := []string{}
	allCards = append(allCards, t.CommunityCards...)
	allCards = append(allCards, h.Cards...)
	// convert cards to int values
	// convertedCards := CardValues(allCards)
	// look for card pairs or groups in order

	return "High Card"
}

func (t *Table) EvaluateHands(hands ...*Hand) Hand {

	return Hand{}
}

func (t *Table) String() string {
	return strings.Join(t.Deck, " ")
}

func (h *Hand) String() string {
	return strings.Join(h.Cards, " ")
}

func (h *Hand) Show() {
	fmt.Printf("______\n")
	fmt.Printf("|   %v|\n", Suit(h.Cards[0]))
	fmt.Printf("|%v\n", Rank(h.Cards[0]))
	fmt.Printf("______\n")
	fmt.Printf("|   %v|\n", Suit(h.Cards[1]))
	fmt.Printf("|%v\n", Rank(h.Cards[1]))
}

func (t *Table) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(t.Deck), func(i, j int) {
		t.Deck[i], t.Deck[j] = t.Deck[j], t.Deck[i]
	})
}

func (t *Table) Deal2(h *Hand) {
	h.Cards = t.Deck[:2]
	t.Deck = t.Deck[2:]
}

// napisz napisz funkcję znajdującą najlepszy układ kart na ręce i stole
// napisz testy do tej funkcji
