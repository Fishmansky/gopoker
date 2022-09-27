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
	case "TestCardValues":
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
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "J.S", "Q.S", "K.S"}}
	case "TestEvaluateHands":
		suite.testHand1 = croupier.Hand{PlayerName: "Player1", Cards: []string{"10.S", "A.S"}}
		suite.testHand2 = croupier.Hand{PlayerName: "Player1", Cards: []string{"8.S", "9.S"}}
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "7.S", "J.S", "Q.S", "K.S"}}
	case "TestEvaluateHand_Pair":
		suite.testHand = croupier.Hand{PlayerName: "Player1", Cards: []string{"K.D", "A.S"}}
		suite.testTable = croupier.Table{Deck: []string{"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "J.S", "Q.S", "K.S"}, CommunityCards: []string{"Q.S", "K.S"}}
	case "TestEvaluateHand_Order":
		suite.testHand1 = croupier.Hand{PlayerName: "Player1", Cards: []string{"10.S", "A.S"}}
		suite.testTable = croupier.Table{CommunityCards: []string{"4.S", "5.S", "6.S", "J.S", "Q.S", "K.S"}}
	}

}

func (suite *CroupierTestSuite) TestRank() {
	result := croupier.Rank(suite.testHand.Cards[0])
	suite.Equal("10", result)
}

func (suite *CroupierTestSuite) TestSuit() {
	result := croupier.Suit(suite.testHand.Cards[0])
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

func (suite *CroupierTestSuite) TestCardValues() {
	result := croupier.CardValues(suite.testTable.Deck)
	suite.Equal([]int{2, 3}, result)
}

func (suite *CroupierTestSuite) TestTableString() {
	suite.Equal("10.D", suite.testTable.String())
}

func (suite *CroupierTestSuite) TestShuffle() {
	suite.testTable.Shuffle()
	suite.NotEqual("2.S 3.S 4.S 5.S 6.S 8.S 9.S 10.S J.S Q.S K.S A.S", suite.testTable.String())
}

func (suite *CroupierTestSuite) TestEvaluateHand() {
	result := suite.testTable.EvaluateHand(&suite.testHand)
	suite.Equal("Royal Flush", result)
}

func (suite *CroupierTestSuite) TestEvaluateHand_Pair() {
	result := suite.testTable.EvaluateHand(&suite.testHand)
	suite.Equal("Pair", result)
}
func (suite *CroupierTestSuite) TestEvaluateHand_Order() {
	result := suite.testTable.EvaluateHand(&suite.testHand1)
	suite.Equal("Royal Flush", result)
}

func (suite *CroupierTestSuite) TestEvaluateHands() {
	result := suite.testTable.EvaluateHands(&suite.testHand1, &suite.testHand2)
	suite.Equal(suite.testHand1.PlayerName, result.PlayerName)
}

func TestCroupierTestSuite(t *testing.T) {
	suite.Run(t, &CroupierTestSuite{})
}
