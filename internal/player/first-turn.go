package player

import (
	"fmt"

	"github.com/yoru0/capsa-custom/internal/deck"
	"github.com/yoru0/capsa-custom/internal/design"
)

func (p *Players) WhoPlaysFirst() int {
	for i := range *p {
		for _, card := range (*p)[i].Hand {
			if card.Rank == deck.Three && card.Suit == deck.Spades {
				return i
			}
		}
	}
	return -1
}

func (p *Players) RemoveThree() {
	for i := range *p {
		var picks []int
		fmt.Printf("%s: ", (*p)[i].Name)
		for j, card := range (*p)[i].Hand {
			if card.Rank == deck.Three {
				design.PrintIndividualCardWithColor(card)
				picks = append(picks, j)
			}
		}
		(*p)[i].RemovePlayedCards(picks)
		fmt.Println()
	}
}