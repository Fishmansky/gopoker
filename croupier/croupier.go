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

var BestHandsOrder = map[string]int{
	"Royal Flush":    0,
	"Straight Flush": 1,
	"Four of kind":   2,
	"Full House":     3,
	"Flush":          4,
	"Streigh":        5,
	"Three of kind":  6,
	"Two pairs":      7,
	"Pair":           8,
	"High Card":      9,
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

func SameSuitCards(cards []string) bool {
	suit := SuitStr(cards[0])
	for _, card := range cards[1:] {
		if SuitStr(card) != suit {
			return false
		}
	}
	return true
}

func SortCardsDesc(cards []string) []string {
	sorted := []string{}
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
	allcards = SortCardsDesc(allcards)
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
	sameKind := make(map[string][]string, 0)
	for _, handCard := range handCards {
		sameKind[RankStr(handCard)] = append(sameKind[RankStr(handCard)], handCard)
	}
	for _, tableCard := range tableCards {
		sameKind[RankStr(tableCard)] = append(sameKind[RankStr(tableCard)], tableCard)
	}
	for kind, cards := range sameKind {
		if len(cards) < 2 {
			delete(sameKind, kind)
		}
	}
	return sameKind
}

func FindSameSuit(handCards []string, tableCards []string) map[string][]string {
	samesuitcards := make(map[string][]string, 0)
	if SameSuitCards(handCards) {
		samesuitcards[SuitStr(handCards[0])] = handCards
		for _, card := range tableCards {
			if CompareSuit(card, handCards[0]) {
				samesuitcards[SuitStr(handCards[0])] = append(samesuitcards[SuitStr(handCards[0])], card)
			}
		}
		if len(samesuitcards[SuitStr(handCards[0])]) < 5 {
			delete(samesuitcards, SuitStr(handCards[0]))
		}
	} else {
		for _, handCard := range handCards {
			samesuitcards[SuitStr(handCard)] = []string{handCard}
			for _, card := range tableCards {
				if CompareSuit(card, handCard) {
					samesuitcards[SuitStr(handCard)] = append(samesuitcards[SuitStr(handCard)], card)
				}
			}
			if len(samesuitcards[SuitStr(handCard)]) < 5 {
				delete(samesuitcards, SuitStr(handCard))
			}
		}
	}
	return samesuitcards
}

func (t *Table) EvaluateHand(h *Hand) (string, []string) {
	kinds := FindSameKind(h.Cards, t.CommunityCards)
	order := FindOrder(h.Cards, t.CommunityCards)
	suits := FindSameSuit(h.Cards, t.CommunityCards)
	pairs := 0
	threes := 0
	fours := 0
	bestHand := "High Card"
	allcards := []string{}
	allcards = append(allcards, h.Cards...)
	allcards = append(allcards, t.CommunityCards...)
	bestcards := SortCardsDesc(allcards)
	samesuit := len(suits) > 0
	inOrder := len(order) > 0
	for _, cards := range kinds {
		switch len(cards) {
		case 2:
			pairs += 1
		case 3:
			threes += 1
		case 4:
			fours += 1
		}
	}

	if samesuit && inOrder {
		for highC, cards := range order {
			if RankInt(highC) == 14 {
				bestHand = BestHands[0]
				bestcards = cards
			} else {
				bestHand = BestHands[1]
				bestcards = cards
			}
		}
	} else if !samesuit && inOrder {
		for _, cards := range order {
			bestHand = BestHands[1]
			bestcards = cards
		}

	} else if samesuit && !inOrder {
		for _, cards := range suits {
			bestHand = BestHands[4]
			bestcards = cards
		}

	}
	if fours == 1 {
		bestHand = BestHands[2]
		for _, cards := range kinds {
			bestcards = cards
		}
		for _, card := range h.Cards {
			if !containsStr(bestcards, card) {
				bestcards = append(bestcards, card)
			}
		}

	} else if threes == 1 {
		if pairs == 1 {
			bestHand = BestHands[3]
		} else {
			bestHand = BestHands[6]
		}
	} else if threes == 0 && fours == 0 {
		if pairs == 2 {
			bestHand = BestHands[7]
		} else if pairs == 1 {
			bestHand = BestHands[8]
		}
	}
	fmt.Println(bestcards)
	return bestHand, bestcards
}

func (t *Table) EvaluateHands(hands ...*Hand) (string, []string) {
	winnerResult := "High Card"
	winner := ""
	winnerCards := []string{}
	for _, hand := range hands {
		result, cards := t.EvaluateHand(hand)
		if BestHandsOrder[winnerResult] > BestHandsOrder[result] {
			winnerResult = result
			winner = hand.PlayerName
			winnerCards = cards
		} else if BestHandsOrder[winnerResult] == BestHandsOrder[result] {
			switch result {
			case BestHands[0], BestHands[1], BestHands[4]:
				HighestCard := 0
				for _, hand := range hands {
					if HighestCard < RankInt(SortCardsDesc(hand.Cards)[0]) {
						HighestCard = RankInt(SortCardsDesc(hand.Cards)[0])
						winner = hand.PlayerName
					}
				}
			case BestHands[2], BestHands[3], BestHands[6], BestHands[7], BestHands[8], BestHands[9]:

			}
		}
	}
	return winner, winnerCards
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
	fmt.Printf("|%v   |\n", RankStr(h.Cards[0]))
	fmt.Printf("______\n")
	fmt.Printf("|   %v|\n", SuitStr(h.Cards[1]))
	fmt.Printf("|%v   |\n", RankStr(h.Cards[1]))
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
