package main

import (
	"fmt"

	"github.com/fishmansky/gopoker/croupier"
)

func main() {
	table := croupier.Table{Deck: []string{
		"2.S", "3.S", "4.S", "5.S", "6.S", "8.S", "9.S", "10.S", "J.S", "Q.S", "K.S", "A.S",
		"2.H", "3.H", "4.H", "5.H", "6.H", "8.H", "9.H", "10.H", "J.H", "Q.H", "K.H", "A.H",
		"2.D", "3.D", "4.D", "5.D", "6.D", "8.D", "9.D", "10.D", "J.D", "Q.D", "K.D", "A.D",
		"2.C", "3.C", "4.C", "5.C", "6.C", "8.C", "9.C", "10.C", "J.C", "Q.C", "K.C", "A.C"},
	}
	table.Shuffle()

	myHand := croupier.Hand{}
	table.Deal2(&myHand)

	fmt.Println(myHand.String())
	myHand.Show()
}
