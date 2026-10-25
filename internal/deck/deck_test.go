package deck

import "testing"

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	if len(d) != 52 {
		t.Fatalf("len(NewDeck()) = %d, want 52", len(d))
	}
	seen := make(map[Card]bool)
	for _, card := range d {
		if seen[card] {
			t.Errorf("duplicate card %v", card)
		}
		seen[card] = true
	}
}

func TestDealDeck(t *testing.T) {
	d := NewDeck()
	hands := d.DealDeck(4)
	if len(hands) != 4 {
		t.Fatalf("got %d hands, want 4", len(hands))
	}
	seen := make(map[Card]bool)
	for i, hand := range hands {
		if len(hand) != 13 {
			t.Errorf("hand %d has %d cards, want 13", i, len(hand))
		}
		for _, card := range hand {
			if seen[card] {
				t.Errorf("card %v dealt twice", card)
			}
			seen[card] = true
		}
	}
	if len(d) != 0 {
		t.Errorf("%d cards left in deck after dealing, want 0", len(d))
	}
}

func TestSortDeckByRankAscending(t *testing.T) {
	d := Deck{
		{Rank: Two, Suit: Diamonds},
		{Rank: Five, Suit: Spades},
		{Rank: Five, Suit: Diamonds},
		{Rank: Ace, Suit: Hearts},
	}
	d.SortDeckByRankAscending()
	want := Deck{
		{Rank: Five, Suit: Diamonds},
		{Rank: Five, Suit: Spades},
		{Rank: Ace, Suit: Hearts},
		{Rank: Two, Suit: Diamonds},
	}
	for i := range want {
		if d[i] != want[i] {
			t.Fatalf("sorted = %v, want %v", d, want)
		}
	}
}

func TestStrings(t *testing.T) {
	tests := []struct {
		card Card
		want string
	}{
		{Card{Rank: Three, Suit: Spades}, "3♠"},
		{Card{Rank: Ten, Suit: Hearts}, "10♥"},
		{Card{Rank: Jack, Suit: Clubs}, "J♣"},
		{Card{Rank: Two, Suit: Diamonds}, "2♦"},
	}
	for _, tt := range tests {
		if got := tt.card.Rank.String() + tt.card.Suit.String(); got != tt.want {
			t.Errorf("card %v = %q, want %q", tt.card, got, tt.want)
		}
	}
}
