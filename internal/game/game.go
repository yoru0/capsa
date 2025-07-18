package game

import (
	"fmt"

	"github.com/yoru0/capsa-custom/internal/combo"
	"github.com/yoru0/capsa-custom/internal/player"
)

type Game struct {
	Players       player.Players
	CurrIndex     int
	Round         int
	PlayerSkipped int
	ComboPlayed   combo.Combo
	ComboToBeat   combo.Combo
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

func (g Game) CheckGameDetails(idx int) {
	fmt.Println("Player        : Name :", g.Players[idx].Name)
	fmt.Println("                Hand :", g.Players[idx].Hand)
	fmt.Println("                Skip :", g.Players[idx].Skip)
	fmt.Println("CurrIndex     :", g.CurrIndex)
	fmt.Println("Round         :", g.Round)
	fmt.Println("PlayerSkipped :", g.PlayerSkipped)
	fmt.Println("ComboPlayed   : Type  :", g.ComboPlayed.Type)
	fmt.Println("                Cards :", g.ComboPlayed.Cards)
	fmt.Println("                Power :", g.ComboPlayed.Power)
	fmt.Println("ComboToBeat   : Type  :", g.ComboPlayed.Type)
	fmt.Println("                Cards :", g.ComboPlayed.Cards)
	fmt.Println("                Power :", g.ComboPlayed.Power)
}
