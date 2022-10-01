package main

import (
	"testing"

	"github.com/fishmansky/gopoker/croupier"
	"github.com/stretchr/testify/suite"
)

type CroupierTestSuite struct {
	suite.Suite
	testHand  croupier.Hand
	testHand1 croupier.Hand
	testHand2 croupier.Hand
	testTable croupier.Table
}

func (suite *CroupierTestSuite) BeforeTest(suiteName, testName string) {
	switch testName {
	case "TestRank", "TestSuit", "TestHandString":
		suite.testHand = croupier.Hand{Cards: []string{"10.D"}}
	case "TestCardsIntValues":
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S"}}
	case "TestDeal2":
		suite.testHand = croupier.Hand{PlayerName: "Player1", Cards: []string{}}
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "9.S", "10.S", "J.S", "Q.S", "K.S", "A.S"}}
	case "TestTableString":
		suite.testTable = croupier.Table{Deck: []string{"10.D"}}
	case "TestShuffle":
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "9.S", "10.S", "J.S", "Q.S", "K.S", "A.S"}}
	case "TestEvaluateHand":
		suite.testHand = croupier.Hand{PlayerName: "Player1", Cards: []string{"10.S", "A.S"}}
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "J.S", "Q.S", "K.S"}, CommunityCards: []string{"Q.S", "K.S", "J.S"}}
	case "TestEvaluateHand_RoyalFlush":
		suite.testHand = croupier.Hand{PlayerName: "Player1", Cards: []string{"10.S", "A.S"}}
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "J.S", "Q.S", "K.S"}, CommunityCards: []string{"Q.S", "K.S", "J.S"}}
	case "TestEvaluateHand_StraightFlush":
		suite.testHand = croupier.Hand{PlayerName: "Player1", Cards: []string{"10.D", "A.S"}}
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "J.S", "Q.S", "K.S"}, CommunityCards: []string{"Q.H", "K.S", "J.S"}}
	case "TestEvaluateHand_FourOfKind":
		suite.testHand = croupier.Hand{PlayerName: "Player1", Cards: []string{"A.D", "J.S"}}
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "J.S", "Q.S", "K.S"}, CommunityCards: []string{"A.H", "A.C", "A.S"}}
	case "TestEvaluateHands":
		suite.testHand1 = croupier.Hand{PlayerName: "Player1", Cards: []string{"10.S", "A.S"}}
		suite.testHand2 = croupier.Hand{PlayerName: "Player1", Cards: []string{"8.S", "9.S"}}
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "7.S", "J.S", "Q.S", "K.S"}, CommunityCards: []string{"J.S", "Q.S", "K.S"}}
	case "TestFindSameKind":
		suite.testHand = croupier.Hand{PlayerName: "Player1", Cards: []string{"K.D", "A.S"}}
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "J.S", "Q.S", "K.S"}, CommunityCards: []string{"Q.S", "K.S"}}
	case "TestFindOrder":
		suite.testHand = croupier.Hand{PlayerName: "Player1", Cards: []string{"10.S", "A.S"}}
		suite.testTable = croupier.Table{CommunityCards: []string{"4.S", "5.S", "6.S", "J.S", "Q.S", "K.S"}}
	case "TestFindSameSuit":
		suite.testHand = croupier.Hand{PlayerName: "Player1", Cards: []string{"J.S", "A.S"}}
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "J.S", "Q.S", "K.S"}, CommunityCards: []string{"8.S", "6.S", "3.S"}}
	case "TestSortCardsDesc":
		suite.testHand = croupier.Hand{PlayerName: "Player1", Cards: []string{"8.S", "2.S", "K.S", "3.S", "4.S", "J.S", "Q.S", "6.S", "5.S"}}
	}

}

func (suite *CroupierTestSuite) AfterTest(suiteName, testName string) {
	suite.testHand = croupier.Hand{}
	suite.testHand1 = croupier.Hand{}
	suite.testHand2 = croupier.Hand{}
	suite.testTable = croupier.Table{}
}

func (suite *CroupierTestSuite) TestRank() {
	result := croupier.RankStr(suite.testHand.Cards[0])
	suite.Equal("10", result)
}

func (suite *CroupierTestSuite) TestSuit() {
	result := croupier.SuitStr(suite.testHand.Cards[0])
	suite.Equal("D", result)
}

func (suite *CroupierTestSuite) TestHandString() {
	result := suite.testHand.String()
	suite.Equal("10.D", result)
}

func (suite *CroupierTestSuite) TestDeal2() {
	suite.testTable.Deal2(&suite.testHand)
	suite.Equal(2, len(suite.testHand.Cards))
}

func (suite *CroupierTestSuite) TestCardsIntValues() {
	result := croupier.CardsIntValues(suite.testTable.Deck)
	suite.Equal([]int{2, 3}, result)
}

func (suite *CroupierTestSuite) TestTableString() {
	suite.Equal("10.D", suite.testTable.String())
}

func (suite *CroupierTestSuite) TestShuffle() {
	suite.testTable.Shuffle()
	suite.NotEqual("2.S 3.S 4.S 5.S 6.S 8.S 9.S 10.S J.S Q.S K.S A.S", suite.testTable.String())
}

func (suite *CroupierTestSuite) TestFindSameSuit() {
	result := croupier.FindSameSuit(suite.testHand.Cards, suite.testTable.CommunityCards)
	suite.Equal(map[string][]string{"S": {"J.S", "A.S", "8.S", "6.S", "3.S"}}, result)
}

func (suite *CroupierTestSuite) TestFindSameKind() {
	result := croupier.FindSameKind(suite.testHand.Cards, suite.testTable.CommunityCards)
	suite.Equal(map[string][]string{"K": {"K.S", "K.D"}}, result)
}
func (suite *CroupierTestSuite) TestFindOrder() {
	result := croupier.FindOrder(suite.testHand.Cards, suite.testTable.CommunityCards)
	suite.Equal(map[string][]string(map[string][]string{"A.S": {"A.S", "K.S", "Q.S", "J.S", "10.S"}}), result)
}

func (suite *CroupierTestSuite) TestEvaluateHand() {
	result, cards := suite.testTable.EvaluateHand(&suite.testHand)
	suite.Equal("Royal Flush", result)
	suite.Equal([]string([]string{"A.S", "K.S", "Q.S", "J.S", "10.S"}), cards)
}

func (suite *CroupierTestSuite) TestEvaluateHand_RoyalFlush() {
	result, cards := suite.testTable.EvaluateHand(&suite.testHand)
	suite.Equal("Royal Flush", result)
	suite.Equal([]string([]string{"A.S", "K.S", "Q.S", "J.S", "10.S"}), cards)
}

func (suite *CroupierTestSuite) TestEvaluateHand_StraightFlush() {
	result, cards := suite.testTable.EvaluateHand(&suite.testHand)
	suite.Equal("Straight Flush", result)
	suite.Equal([]string([]string{"A.S", "K.S", "Q.H", "J.S", "10.D"}), cards)
}

func (suite *CroupierTestSuite) TestEvaluateHand_FourOfKind() {
	result, cards := suite.testTable.EvaluateHand(&suite.testHand)
	suite.Equal("Four of kind", result)
	suite.Equal([]string{"A.H", "A.C", "A.S", "A.D", "J.S"}, cards)
}

func (suite *CroupierTestSuite) TestEvaluateHands() {
	result, cards := suite.testTable.EvaluateHands(&suite.testHand1, &suite.testHand2)
	suite.Equal(suite.testHand1.PlayerName, result)
	suite.Equal([]string([]string{"A.S", "K.S", "Q.S", "J.S", "10.S"}), cards)
}

func (suite *CroupierTestSuite) TestSortCardsDesc() {
	sorted := []string{}
	result := croupier.SortCardsDesc(suite.testHand.Cards, sorted)
	suite.Equal([]string{"K.S", "Q.S", "J.S", "8.S", "6.S", "5.S", "4.S", "3.S", "2.S"}, result)
}

func TestCroupierTestSuite(t *testing.T) {
	suite.Run(t, &CroupierTestSuite{})
}
