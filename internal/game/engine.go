package game

import (
	"fmt"

	"github.com/yoru0/capsa-custom/internal/deck"
	"github.com/yoru0/capsa-custom/internal/player"
)

func Start() {
	p := player.NewPlayers(4)
	curr := p.PlayFirst()
	p.RemoveThree()
	fmt.Printf("%s play first\n\n", p[curr].Name)

	for len(p) > 1 {

		skip := 0

		for skip < 3 {
			if p[curr].Skip {
				fmt.Printf("%s skipped\n", p[curr].Name)
				curr = (curr + 1) % len(p)
				continue
			}

			fmt.Println("[DEBUG] curr:", curr)
			fmt.Printf("%s [card(s): %d]:\n", p[curr].Name, len(p[curr].Hand))
			p[curr].GetHand()

			var pick []int
			var cards deck.Deck
			var combo string
			var valid bool

			for {
				pick, cards = p[curr].Pick()
				fmt.Println(pick)
				combo, valid = CheckCombo(cards)
				fmt.Println("Combo:", combo)

				if valid {
					break
				}
			}

			p[curr].Remove(pick)

			if combo == "Skip" {
				p[curr].SkipTurn()
				skip++
			}

			if len(p[curr].Hand) == 0 {
				fmt.Printf("[DEBUG] %s removed\n", p[curr].Name)
				p.RemovePlayer(curr)
				if curr >= len(p) {
					curr = 0
				}
			} else {
				curr = (curr + 1) % len(p)
			}

			fmt.Println()
		}

		for i := range p {
			p[i].Skip = false
		}
	}
}

func CheckCombo(cards deck.Deck) (string, bool) {
	length := len(cards)

	switch length {
	case 0:
		return "Skip", true
	case 1:
		return "Single", true
	case 2:
		if IsPair(cards) {
			return "Pair", true
		}
	case 3:
		if IsTriple(cards) {
			return "Triple", true
		}
	case 5:
		switch {
		case IsStraightFlush(cards):
			return "Straight Flush", true
		case IsFourOfAKind(cards):
			return "Four of a Kind", true
		case IsFullHouse(cards):
			return "Full House", true
		case IsFlush(cards):
			return "Flush", true
		case IsStraight(cards):
			return "Straight", true
		default:
			fmt.Println("Invalid 5-card combo.")
		}

	default:
		fmt.Println("Invalid combo length.")
	}

	return "Invalid", false
}
