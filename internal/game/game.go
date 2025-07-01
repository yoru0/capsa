package game

import "github.com/yoru0/capsa-custom/internal/player"

type Game struct {
	Players   player.Players
	Curr      int
	Round     int
	LastCombo Combo
}

func NewGame(numPlayers int) *Game {
	p := player.NewPlayers(numPlayers)
	first := p.PlayFirst()
	return &Game{Players: p, Curr: first}
}
