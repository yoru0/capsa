package winner

import (
	"fmt"
	"strconv"
	"strings"
)

type Winner []string

func (w *Winner) AppendWinner(name string) {
	*w = append(*w, name)
}

func (w Winner) ShowWinners(players []string) {
	fmt.Println("──────────────────────────────────────────────────────────────────────────────────────────────────")
	fmt.Println("Winners:")
	for i, name := range w {
		fmt.Printf("%d. %s\n", i+1, name)
	}
	fmt.Printf("%d. %s\n", len(w)+1, w.lastPlace(players))
	fmt.Printf("──────────────────────────────────────────────────────────────────────────────────────────────────\n\n")
}

func (w Winner) lastPlace(players []string) string {
	total := 10

	for _, name := range w {
		parts := strings.Split(name, " ")
		if len(parts) != 2 {
			panic("Invalid name format, expected 'Player X'")
		}
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		total -= num
	}

	return "Player " + strconv.Itoa(total)
}
