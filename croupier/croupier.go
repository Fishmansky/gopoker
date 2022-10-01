package croupier

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// best possible hands in descending order
var BestHands = []string{"Royal Flush", "Straight Flush", "Four of kind", "Full House", "Flush", "Streigh", "Three of kind", "Two pairs", "Pair", "High Card"}

var HandTypes = map[string]Hand{
	// _ for discard
	// * for wildcard
	"Royal Flush":    {Cards: []string{"A.*", "K.*", "Q.*", "J.*", "10.*"}},
	"Straight Flush": {Cards: []string{"X.C", "X-1.C", "X-2.C", "X-3.C", "X-4.C"}}, // X for figure, C for Color
	"Flush":          {Cards: []string{"*.C", "*.C", "*.C", "*.C", "*.C"}},
	"Streigh":        {Cards: []string{"X.*", "X-1.*", "X-2.*", "X-3.*", "X-4.*"}},
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

func pop(s []string, e string) []string {
	for i, a := range s {
		if a == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return []string{}
}

func containsStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func RankStr(card string) string {
	return strings.Split(card, ".")[0]
}

func RankInt(card string) int {
	switch RankStr(card) {
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		intVal, err := strconv.Atoi(RankStr(card))
		if err != nil {
			panic(err)
		}
		return intVal
	}
}

func SuitStr(card string) string {
	return strings.Split(card, ".")[1]
}

func CardsIntValues(cards []string) []int {
	convertedCards := []int{}
	for _, card := range cards {
		var cardVal int
		switch RankStr(card) {
		case "J":
			cardVal = 11
		case "Q":
			cardVal = 12
		case "K":
			cardVal = 13
		case "A":
			cardVal = 14
		default:
			intVal, err := strconv.Atoi(RankStr(card))
			if err != nil {
				panic(err)
			}
			cardVal = intVal
		}
		convertedCards = append(convertedCards, cardVal)
	}
	return convertedCards
}

func CardValueString(card string) string {
	Suit := SuitStr(card)
	switch RankStr(card) {
	case "J":
		return "11." + Suit
	case "Q":
		return "12." + Suit
	case "K":
		return "13." + Suit
	case "A":
		return "14." + Suit
	default:
		return card
	}
}

func CompareRank(card1 string, card2 string) bool {
	return RankInt(card1) == RankInt(card2)
}

func CompareSuit(card1 string, card2 string) bool {
	return SuitStr(card1) == SuitStr(card2)
}

func GetHandTypeStr(cards []string) string {
	sorted := []string{}
	cards = SortCardsDesc(cards, sorted)
	result := []int{}
	for i, besthand := range BestHands {
		result[i] = 0
		for i, card := range cards {
			if CompareRank(card, HandTypes[besthand].Cards[i]) {
				result[i] += 1
			}
		}
		if result[i] == 5 {
			return besthand
		}
	}
	return ""
}

func SortCardsDesc(cards, sorted []string) []string {
	if len(cards) == 0 {
		return sorted
	}
	for len(cards) > 0 {
		HighestCard := ""
		maxInt := 0
		for _, card := range cards {
			if RankInt(card) > maxInt {
				maxInt = RankInt(card)
				HighestCard = card
			}
		}
		cards = pop(cards, HighestCard)
		sorted = append(sorted, HighestCard)
	}
	return sorted
}

func FindOrder(handCards []string, tableCards []string) map[string][]string {
	orderedCards := make(map[string][]string, 0)
	allcards := []string{}
	allcards = append(allcards, handCards...)
	allcards = append(allcards, tableCards...)
	sorted := []string{}
	allcards = SortCardsDesc(allcards, sorted)
	for i, card := range allcards {
		orderedCards[card] = []string{card}
		for _, nextcard := range allcards[i:] {
			if RankInt(nextcard) == RankInt(orderedCards[card][len(orderedCards[card])-1])-1 {
				orderedCards[card] = append(orderedCards[card], nextcard)
			}
		}
		if !containsStr(orderedCards[card], handCards[0]) && !containsStr(orderedCards[card], handCards[1]) || len(orderedCards[card]) != 5 {
			delete(orderedCards, card)
		}
	}
	return orderedCards
}

func FindSameKind(handCards []string, tableCards []string) map[string][]string {
	// returns map of pair
	// handcard is key
	// table cards are stored as values
	sameKind := make(map[string][]string, 0)
	for _, handCard := range handCards {
		for _, tableCard := range tableCards {
			if RankStr(tableCard) == RankStr(handCard) {
				sameKind[handCard] = append(sameKind[handCard], tableCard)
			}
		}
	}
	return sameKind
}

func (t *Table) EvaluateHand(h *Hand) string {
	kinds := FindSameKind(h.Cards, t.CommunityCards)
	order := FindOrder(h.Cards, t.CommunityCards)
	if len(order) > 0 {
		for _, o := range order {
			GetHandTypeStr(o)

			// TODO:
			// check all orders and return the best
		}
	}
	if len(kinds) == 1 {
		return "Pair"
		// TODO:
		// check all pairs and return the best
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
	fmt.Printf("|   %v|\n", SuitStr(h.Cards[0]))
	fmt.Printf("|%v\n", RankStr(h.Cards[0]))
	fmt.Printf("______\n")
	fmt.Printf("|   %v|\n", SuitStr(h.Cards[1]))
	fmt.Printf("|%v\n", RankStr(h.Cards[1]))
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
