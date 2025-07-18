package main

import (
	"fmt"

	"github.com/yoru0/capsa-custom/internal/deck"
)

func main() {

	d := deck.NewDeck()
	for i := range d {
		fmt.Printf("%s ", d[i].Rank.String()+d[i].Suit.String())
	}

}
