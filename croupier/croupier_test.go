package croupier

import (
	"fmt"
	"testing"
)

func TestHandString(t *testing.T) {
	testHand := Hand{Cards: []string{"10D"}}

	result := testHand.String()

	if result != "10D" {
		t.Errorf("Hand String() FAILED - Expected %v, got %v\n", "10D", result)
	}
}

func TestDeal2(t *testing.T) {
	testHand := Hand{Cards: []string{}}
	table := Table{Deck: []string{
		"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "9.S", "10.S", "J.S", "Q.S", "K.S", "A.S"},
	}
	table.Deal2(&testHand)
	result := len(testHand.Cards)

	if result != 2 {
		t.Errorf("Hand Deal2(testHand) FAILED - Expected %v cards, got %v\n", 2, result)
	}
}

func TestTableString(t *testing.T) {
	testTable := Table{Deck: []string{"10D"}}

	result := testTable.String()

	if result != "10D" {
		t.Errorf("Table String(g) FAILED - Expected %v, got %v\n", "10D", result)
	}
}

func TestShuffle(t *testing.T) {
	table := Table{Deck: []string{
		"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "9.S", "10.S", "J.S", "Q.S", "K.S", "A.S"},
	}

	originalOrder := "2.S 3.S 4.S 5.S 6.S 8.S 9.S 10.S J.S Q.S K.S A.S"

	table.Shuffle()
	result := table.String()

	if result == originalOrder {
		t.Errorf("Table String(g) FAILED - Expected %v, got %v\n", "10D", result)
	}
}

func TestEvaluateHands(t *testing.T) {
	testHand1 := Hand{PlayerName: "Player1", Cards: []string{"10.S", "A.S"}}
	testHand2 := Hand{PlayerName: "Player2", Cards: []string{"J.S", "Q.S"}}
	table := Table{Deck: []string{
		"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "9.S", "Q.S", "K.S"},
	}
	result := table.EvaluateHands(&testHand1, &testHand2)
	fmt.Println(result.PlayerName)
	if result.PlayerName != testHand1.PlayerName {
		t.Errorf("Hand EvaluateHands(&testHand1, &testHand2) FAILED - Expected %v, got %v\n", testHand1.PlayerName, result.PlayerName)
	}
}
