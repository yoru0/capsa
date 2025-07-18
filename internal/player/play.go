package player

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/yoru0/capsa-custom/internal/deck"
)

func (p Player) Pick() ([]int, deck.Deck) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Pick your cards: ")
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)

		var picks []int
		var cards deck.Deck
		valid := true
		seen := make(map[int]bool)

		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil || num < 1 || num > len(p.Hand) {
				fmt.Println("Invalid input:", part)
				valid = false
				break
			}

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

		if valid {
			return picks, cards
		}

		fmt.Println("Please enter valid numbers between 1 and", len(p.Hand))
	}
}

func (p *Player) RemovePlayedCards(picks []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(picks)))
	for _, i := range picks {
		if i >= 0 && i < len(p.Hand) {
			p.Hand = append((p.Hand)[:i], (p.Hand)[i+1:]...)
		}
	}
}
