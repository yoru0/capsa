package player

import (
	"testing"

	"github.com/yoru0/capsa/internal/deck"
)

func TestNewPlayers(t *testing.T) {
	p := NewPlayers(4)
	if len(p) != 4 {
		t.Fatalf("got %d players, want 4", len(p))
	}
	for i, pl := range p {
		if pl.Id != i {
			t.Errorf("player %d has Id %d", i, pl.Id)
		}
		if len(pl.Hand) != 13 {
			t.Errorf("%s has %d cards, want 13", pl.Name, len(pl.Hand))
		}
	}
}

func TestWhoPlaysFirst(t *testing.T) {
	p := Players{
		{Name: "Player 1", Hand: deck.Deck{{Rank: deck.Three, Suit: deck.Hearts}}},
		{Name: "Player 2", Hand: deck.Deck{{Rank: deck.Four, Suit: deck.Spades}}},
		{Name: "Player 3", Hand: deck.Deck{{Rank: deck.Three, Suit: deck.Spades}}},
	}
	if got := p.WhoPlaysFirst(); got != 2 {
		t.Errorf("WhoPlaysFirst() = %d, want 2", got)
	}
}

func TestRemoveThree(t *testing.T) {
	p := Players{
		{Name: "Player 1", Hand: deck.Deck{
			{Rank: deck.Three, Suit: deck.Diamonds},
			{Rank: deck.Three, Suit: deck.Spades},
			{Rank: deck.Nine, Suit: deck.Clubs},
		}},
		{Name: "Player 2", Hand: deck.Deck{{Rank: deck.King, Suit: deck.Hearts}}},
	}
	p.RemoveThree()
	for _, pl := range p {
		for _, card := range pl.Hand {
			if card.Rank == deck.Three {
				t.Errorf("%s still holds %v", pl.Name, card)
			}
		}
	}
	if len(p[0].Hand) != 1 || len(p[1].Hand) != 1 {
		t.Errorf("hand sizes = %d, %d; want 1, 1", len(p[0].Hand), len(p[1].Hand))
	}
}

func TestRemovePlayedCards(t *testing.T) {
	p := Player{Hand: deck.Deck{
		{Rank: deck.Four, Suit: deck.Diamonds},
		{Rank: deck.Five, Suit: deck.Diamonds},
		{Rank: deck.Six, Suit: deck.Diamonds},
		{Rank: deck.Seven, Suit: deck.Diamonds},
	}}
	p.RemovePlayedCards([]int{0, 2})
	want := deck.Deck{
		{Rank: deck.Five, Suit: deck.Diamonds},
		{Rank: deck.Seven, Suit: deck.Diamonds},
	}
	if len(p.Hand) != len(want) || p.Hand[0] != want[0] || p.Hand[1] != want[1] {
		t.Errorf("hand = %v, want %v", p.Hand, want)
	}
}

func TestCheckPickedCardIsValid(t *testing.T) {
	p := Player{Hand: deck.Deck{
		{Rank: deck.Four, Suit: deck.Diamonds},
		{Rank: deck.Five, Suit: deck.Diamonds},
		{Rank: deck.Six, Suit: deck.Diamonds},
	}}
	tests := []struct {
		name      string
		parts     []string
		wantValid bool
		wantPicks []int
	}{
		{"empty is skip", nil, true, nil},
		{"one card", []string{"2"}, true, []int{1}},
		{"several cards", []string{"3", "1"}, true, []int{2, 0}},
		{"zero", []string{"0"}, false, nil},
		{"past end", []string{"4"}, false, nil},
		{"not a number", []string{"a"}, false, nil},
		{"duplicate", []string{"1", "1"}, false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			picks, cards, valid := p.checkPickedCardIsValid(tt.parts)
			if valid != tt.wantValid {
				t.Fatalf("valid = %v, want %v", valid, tt.wantValid)
			}
			if !valid {
				return
			}
			if len(picks) != len(tt.wantPicks) {
				t.Fatalf("picks = %v, want %v", picks, tt.wantPicks)
			}
			for i := range picks {
				if picks[i] != tt.wantPicks[i] || cards[i] != p.Hand[picks[i]] {
					t.Errorf("picks = %v cards = %v, want picks %v", picks, cards, tt.wantPicks)
				}
			}
		})
	}
}
