package player

import (
	"fmt"

	"github.com/yoru0/capsa-custom/internal/deck"
	"github.com/yoru0/capsa-custom/internal/design"
)

type Player struct {
	Id   int
	Name string
	Hand deck.Deck
	Skip bool
}

type Players []Player

func NewPlayers(numPlayers int) Players {
	d := deck.NewDeck()
	d.Shuffle()
	hands := d.Deal(numPlayers)
	players := make(Players, numPlayers)

	for i := range hands {
		hands[i].Sort()
	}

	for i := 0; i < numPlayers; i++ {
		players[i] = Player{
			Id:   i,
			Name: fmt.Sprintf("Player %d", i+1),
			Hand: hands[i],
			Skip: false,
		}
	}
	return players
}

func (p Player) GetHand() {
	for i, card := range p.Hand {
		design.CardColor(card, i)
	}
	fmt.Println()
}

func (p *Players) RemovePlayer(id int) {
	if id < 0 || id >= len(*p) {
		return
	}
	*p = append((*p)[:id], (*p)[id+1:]...)
}

func (p *Player) SkipTurn() {
	p.Skip = true
}