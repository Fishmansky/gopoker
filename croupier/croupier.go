package croupier

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

// best possible hands in descending order
var BestHands = []string{"Royal Flush", "Straight Flush", "Four of kind", "Full House", "Flush", "Streigh", "Three of kind", "Two pairs", "Pair", "High Card"}

var HandTypes = map[string]Hand{
	// _ for discard
	// * for wildcard
	"Royal Flush":    {Cards: []string{"10.*", "11.*", "12.*", "13.*", "14.*"}},
	"Straight Flush": {Cards: []string{"X.C", "X+1.C", "X+2.C", "X+3.C", "X+4.C"}}, // X for figure, C for Color
	"Four of kind":   {Cards: []string{"X.*", "X.*", "X.*", "X.*", "*.*"}},
	"Full House":     {Cards: []string{"X.*", "X.*", "X.*", "Y.*", "Y.*"}},
	"Flush":          {Cards: []string{"*.C", "*.C", "*.C", "*.C", "*.C"}},
	"Streigh":        {Cards: []string{"X.*", "X+1.*", "X+2.*", "X+3.*", "X+4.*"}},
	"Three of kind":  {Cards: []string{"X.*", "X.*", "X.*", "Y.*", "Z.*"}},
	"Two pairs":      {Cards: []string{"X.*", "X.*", "Y.*", "Y.*", "_.*"}},
	"Pair":           {Cards: []string{"X.*", "X.*", "_.*", "_.*", "_.*"}},
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

type Pair []int

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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

func CardRaks(cards []int) []string {
	convertedCards := []string{}
	for _, card := range cards {
		switch card {
		case 11:
			convertedCards = append(convertedCards, "J")
		case 12:
			convertedCards = append(convertedCards, "Q")
		case 13:
			convertedCards = append(convertedCards, "K")
		case 14:
			convertedCards = append(convertedCards, "A")
		default:
			convertedCards = append(convertedCards, strconv.Itoa(card))
		}
	}
	return convertedCards
}

func CompareRank(card1 string, card2 int) bool {
	return Rank(card1) == strconv.Itoa(card2)
}

func CompareSuit(card1 string, card2 int) bool {
	return false
}

func GetHandType(cards []int) string {
	for result, order := range HandTypes {
		for i, card := range cards {
			if CompareRank(order.Cards[i], card) {
				break
			}
			return result
		}
	}
	return ""
}

func FindOrder(handCards []int, tableCards []int) map[int][]int {
	orderedCards := make(map[int][]int, 0)
	allcards := []int{}
	allcards = append(allcards, handCards...)
	allcards = append(allcards, tableCards...)
	sort.Ints(allcards)
	for i, card := range allcards {
		orderedCards[card] = []int{card}
		for _, nextcard := range allcards[i:] {
			if nextcard == orderedCards[card][len(orderedCards[card])-1]+1 {
				orderedCards[card] = append(orderedCards[card], nextcard)
			}
		}
		if !contains(orderedCards[card], handCards[0]) && !contains(orderedCards[card], handCards[1]) || len(orderedCards[card]) != 5 {
			delete(orderedCards, card)
		}
	}
	// TODO:
	// 1. find a way to ranks convertion looseless
	// 2. check if there is more than one order - if so return best one

	return orderedCards
}

func FindSameKind(handCards []int, tableCards []int) map[int]Pair {
	// returns map of pair
	// handcard is key
	// table card indexes as stored as values
	sameKind := make(map[int]Pair, 0)
	for _, handCard := range handCards {
		for i, tableCard := range tableCards {
			if tableCard == handCard {
				sameKind[handCard] = append(sameKind[handCard], i)
			}
		}
	}
	return sameKind
}

func (t *Table) EvaluateHand(h *Hand) string {
	// get cards from hand and table and convert them to int values
	convertedHandCards := CardValues(h.Cards)
	convertedTableCards := CardValues(t.CommunityCards)

	// look for card pairs or groups in order
	kinds := FindSameKind(convertedHandCards, convertedTableCards)
	order := FindOrder(convertedHandCards, convertedTableCards)
	fmt.Println(kinds)
	fmt.Println(order)
	if len(order) > 0 {
		for _, o := range order {
			return GetHandType(o)
		}
	}
	if len(kinds) == 1 {
		return "Pair"
	}

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
