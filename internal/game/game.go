package game

import "github.com/yoru0/capsa-custom/internal/player"

type Game struct {
	Players       player.Players
	CurrIndex     int
	Round         int
	PlayerSkipped int
	LastCombo     Combo
}

func NewGame(numPlayers int) *Game {
	p := player.NewPlayers(numPlayers)
	first := p.WhoPlaysFirst()
	return &Game{
		Players: p, 
		CurrIndex: first,
		Round: 1,
		PlayerSkipped: 0,
	}
}

func (g *Game) ResetPlayerSkips() {
	for i := range g.Players {
		g.Players[i].Skip = false
	}
	g.PlayerSkipped = 0
}
