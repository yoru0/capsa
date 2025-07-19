package game

import (
	"fmt"

	"github.com/yoru0/capsa-custom/internal/combo"
	"github.com/yoru0/capsa-custom/internal/deck"
	"github.com/yoru0/capsa-custom/internal/player"
)

type Game struct {
	Players       player.Players
	CurrIndex     int
	Round         int
	PlayerSkipped int
	ComboPlayed   combo.Combo
	ComboToBeat   combo.Combo
	PlayHistory   deck.CardHistory
}

func NewGame(numPlayers int) *Game {
	p := player.NewPlayers(numPlayers)
	first := p.WhoPlaysFirst()
	p.RemoveThree()
	return &Game{
		Players:       p,
		CurrIndex:     first,
		Round:         1,
		PlayerSkipped: 0,
		ComboPlayed:   combo.Combo{Type: combo.None},
		ComboToBeat:   combo.Combo{Type: combo.None},
		PlayHistory:   nil,
	}
}

func (g *Game) ResetPlayerSkips() {
	for i := range g.Players {
		g.Players[i].Skip = false
	}
	g.PlayerSkipped = 0
}

func (g *Game) ResetLastCombo() {
	g.ComboToBeat.Type = combo.None
}

func (g *Game) NextPlayerTurn() {
	g.CurrIndex = (g.CurrIndex + 1) % len(g.Players)
}

func (g *Game) SetCurrentPlayer(name string) {
	for i := range g.Players {
		if g.Players[i].Name == name {
			g.CurrIndex = g.Players[i].Id
			return
		}
	}
}

func (g Game) CheckGameDetails(idx int) {
	fmt.Printf("{\n")
	fmt.Printf("  \"Player\": {\n")
	fmt.Printf("    \"Name\": \"%s\",\n", g.Players[idx].Name)
	fmt.Printf("    \"Hand\": %v,\n", g.Players[idx].Hand)
	fmt.Printf("    \"Skip\": %v\n", g.Players[idx].Skip)
	fmt.Printf("  },\n")

	fmt.Printf("  \"CurrIndex\": %d,\n", g.CurrIndex)
	fmt.Printf("  \"Round\": %d,\n", g.Round)
	fmt.Printf("  \"PlayerSkipped\": %d,\n", g.PlayerSkipped)

	fmt.Printf("  \"ComboPlayed\": {\n")
	fmt.Printf("    \"Type\": \"%v\",\n", g.ComboPlayed.Type)
	fmt.Printf("    \"Cards\": %v,\n", g.ComboPlayed.Cards)
	fmt.Printf("    \"Power\": %v\n", g.ComboPlayed.Power)
	fmt.Printf("  },\n")

	fmt.Printf("  \"ComboToBeat\": {\n")
	fmt.Printf("    \"Type\": \"%v\",\n", g.ComboPlayed.Type)
	fmt.Printf("    \"Cards\": %v,\n", g.ComboPlayed.Cards)
	fmt.Printf("    \"Power\": %v\n", g.ComboPlayed.Power)
	fmt.Printf("  },\n")

	fmt.Printf("  \"PlayHistory\": %v\n", g.PlayHistory)
	fmt.Printf("}\n")
}
