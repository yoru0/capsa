package player

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/yoru0/capsa-custom/internal/combo"
	"github.com/yoru0/capsa-custom/internal/deck"
	"github.com/yoru0/capsa-custom/internal/design"
)

func (p *Player) PickCard(lastCombo combo.Combo) combo.Combo {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Pick your cards: ")
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		var combos combo.Combo

		picks, cards, valid := p.checkPickedCardIsValid(parts)

		// Check if the cards have a valid combo.
		if valid {
			valid, combos = combo.CheckCombo(cards)
			if !valid {
				fmt.Println("Not a playable combo.")
			}
		}

		// First turn can't skip.
		if valid && combos.Type == combo.Skip && lastCombo.Type.String() == "None" {
			valid = false
			fmt.Println("You can't skip this")
		} else if combos.Type == combo.Skip {
			return combos
		}

		// Check if the combo played is stronger than last combo.
		if valid && lastCombo.Type == combo.None {
			valid = true
		} else if valid {
			valid = combo.CheckStrongerCombo(lastCombo, combos)
			if !valid {
				errorNeedStrongerCards(lastCombo)
			}
		}

		if valid {
			if combos.Type == combo.Skip {
				return lastCombo
			}
			p.RemovePlayedCards(picks)
			return combos
		}

		fmt.Println("Please enter valid numbers between 1 and", len(p.Hand))
	}
}

func (p *Player) checkPickedCardIsValid(parts []string) ([]int, deck.Deck, bool) {
	var picks []int
	var cards deck.Deck
	valid := true
	seen := make(map[int]bool)

	for _, part := range parts {
		num, err := strconv.Atoi(part)

		// Check out of bound case.
		if err != nil || num < 1 || num > len(p.Hand) {
			fmt.Println("Invalid input:", part)
			valid = false
			break
		}

		// Check duplicates.
		index := num - 1
		if seen[index] {
			fmt.Printf("Duplicate pick: %d\n", num)
			valid = false
			break
		}
		seen[index] = true

		picks = append(picks, index)
		cards = append(cards, p.Hand[index])
	}
	return picks, cards, valid
}

func (p *Player) RemovePlayedCards(picks []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(picks)))
	for _, i := range picks {
		if i >= 0 && i < len(p.Hand) {
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
		}
	}
}

func errorNeedStrongerCards(lastCombo combo.Combo) {
	fmt.Printf("Need a `%s` with power higher than ", lastCombo.Type)
	for i := range lastCombo.Cards {
		design.PrintIndividualCardWithColor(lastCombo.Cards[i])
	}
	fmt.Println()
}
