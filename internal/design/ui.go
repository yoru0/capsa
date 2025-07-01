package design

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yoru0/capsa-custom/internal/deck"
)

func CardColor(card deck.Card, i int) {
	colorCode := ""

	switch card.Suit {
	case 0, 2:
		colorCode = "\033[31m" // red
	case 1, 3:
		colorCode = "\033[36m" // cyan
	}
	reset := "\033[0m" // reset

	fmt.Printf("%d: %s%-3s%s  ", i+1, colorCode, card.Rank.String()+card.Suit.String(), reset)
}

func WaitForEnter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to continue...")
	_, _ = reader.ReadString('\n')
}
