package deck

import (
	"math/rand"
	"sort"
)

type Deck []Card

func NewDeck() Deck {
	var d Deck
	for i := Diamonds; i <= Spades; i++ {
		for j := Three; j <= Two; j++ {
			d = append(d, Card{Suit: i, Rank: j})
		}
	}
	return d
}

func (d Deck) ShuffleDeck() {
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func (d Deck) SortDeck() {
	sort.Slice(d, func(i, j int) bool { return d[i].Less(d[j]) })
}

func (d *Deck) DealDeck(players int) []Deck {
	per := len(*d) / players
	hands := make([]Deck, players)
	for i := 0; i < players; i++ {
		start := i * per
		end := start + per
		hands[i] = (*d)[start:end]
	}
	total := players * per
	*d = (*d)[total:]
	return hands
}

func (d *Deck) Remove() {
	*d = (*d)[1:]
}
