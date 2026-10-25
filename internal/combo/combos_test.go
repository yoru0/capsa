package combo

import (
	"testing"

	"github.com/yoru0/capsa/internal/deck"
)

// c is a shorthand for building a card in test tables.
func c(r deck.Rank, s deck.Suit) deck.Card {
	return deck.Card{Rank: r, Suit: s}
}

func d(cards ...deck.Card) deck.Deck {
	return deck.Deck(cards)
}

const (
	D = deck.Diamonds
	C = deck.Clubs
	H = deck.Hearts
	S = deck.Spades
)

func TestIsPair(t *testing.T) {
	tests := []struct {
		name  string
		cards deck.Deck
		want  bool
	}{
		{"same rank", d(c(deck.Five, D), c(deck.Five, S)), true},
		{"different rank", d(c(deck.Five, D), c(deck.Six, D)), false},
		{"too few", d(c(deck.Five, D)), false},
		{"too many", d(c(deck.Five, D), c(deck.Five, C), c(deck.Five, H)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPair(tt.cards); got != tt.want {
				t.Errorf("IsPair(%v) = %v, want %v", tt.cards, got, tt.want)
			}
		})
	}
}

func TestIsTriple(t *testing.T) {
	tests := []struct {
		name  string
		cards deck.Deck
		want  bool
	}{
		{"same rank", d(c(deck.King, D), c(deck.King, C), c(deck.King, S)), true},
		{"one differs", d(c(deck.King, D), c(deck.King, C), c(deck.Ace, S)), false},
		{"pair only", d(c(deck.King, D), c(deck.King, C)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTriple(tt.cards); got != tt.want {
				t.Errorf("IsTriple(%v) = %v, want %v", tt.cards, got, tt.want)
			}
		})
	}
}

func TestIsStraight(t *testing.T) {
	tests := []struct {
		name  string
		cards deck.Deck
		want  bool
	}{
		{"in order", d(c(deck.Five, D), c(deck.Six, C), c(deck.Seven, H), c(deck.Eight, S), c(deck.Nine, D)), true},
		{"unsorted", d(c(deck.Nine, D), c(deck.Five, C), c(deck.Eight, H), c(deck.Six, S), c(deck.Seven, D)), true},
		{"ten to ace", d(c(deck.Ten, D), c(deck.Jack, C), c(deck.Queen, H), c(deck.King, S), c(deck.Ace, D)), true},
		{"gap", d(c(deck.Five, D), c(deck.Six, C), c(deck.Seven, H), c(deck.Eight, S), c(deck.Ten, D)), false},
		{"contains pair", d(c(deck.Five, D), c(deck.Five, C), c(deck.Six, H), c(deck.Seven, S), c(deck.Eight, D)), false},
		{"four cards", d(c(deck.Five, D), c(deck.Six, C), c(deck.Seven, H), c(deck.Eight, S)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStraight(tt.cards); got != tt.want {
				t.Errorf("IsStraight(%v) = %v, want %v", tt.cards, got, tt.want)
			}
		})
	}
}

func TestIsFlush(t *testing.T) {
	tests := []struct {
		name  string
		cards deck.Deck
		want  bool
	}{
		{"same suit", d(c(deck.Four, H), c(deck.Seven, H), c(deck.Nine, H), c(deck.Jack, H), c(deck.Ace, H)), true},
		{"one off suit", d(c(deck.Four, H), c(deck.Seven, H), c(deck.Nine, H), c(deck.Jack, H), c(deck.Ace, S)), false},
		{"four cards", d(c(deck.Four, H), c(deck.Seven, H), c(deck.Nine, H), c(deck.Jack, H)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFlush(tt.cards); got != tt.want {
				t.Errorf("IsFlush(%v) = %v, want %v", tt.cards, got, tt.want)
			}
		})
	}
}

func TestIsFullHouse(t *testing.T) {
	tests := []struct {
		name  string
		cards deck.Deck
		want  bool
	}{
		{"low triple", d(c(deck.Six, D), c(deck.Six, C), c(deck.Six, H), c(deck.King, S), c(deck.King, D)), true},
		{"high triple", d(c(deck.Six, D), c(deck.Six, C), c(deck.King, H), c(deck.King, S), c(deck.King, D)), true},
		{"two pair", d(c(deck.Six, D), c(deck.Six, C), c(deck.Nine, H), c(deck.King, S), c(deck.King, D)), false},
		{"four of a kind", d(c(deck.Six, D), c(deck.Six, C), c(deck.Six, H), c(deck.Six, S), c(deck.King, D)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFullHouse(tt.cards); got != tt.want {
				t.Errorf("IsFullHouse(%v) = %v, want %v", tt.cards, got, tt.want)
			}
		})
	}
}

func TestIsFourOfAKind(t *testing.T) {
	tests := []struct {
		name  string
		cards deck.Deck
		want  bool
	}{
		{"kicker high", d(c(deck.Six, D), c(deck.Six, C), c(deck.Six, H), c(deck.Six, S), c(deck.King, D)), true},
		{"kicker low", d(c(deck.Four, D), c(deck.Six, C), c(deck.Six, H), c(deck.Six, S), c(deck.Six, D)), true},
		{"full house", d(c(deck.Six, D), c(deck.Six, C), c(deck.Six, H), c(deck.King, S), c(deck.King, D)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFourOfAKind(tt.cards); got != tt.want {
				t.Errorf("IsFourOfAKind(%v) = %v, want %v", tt.cards, got, tt.want)
			}
		})
	}
}

func TestIsStraightFlush(t *testing.T) {
	sf := d(c(deck.Five, C), c(deck.Six, C), c(deck.Seven, C), c(deck.Eight, C), c(deck.Nine, C))
	if !IsStraightFlush(sf) {
		t.Errorf("IsStraightFlush(%v) = false, want true", sf)
	}
	straight := d(c(deck.Five, C), c(deck.Six, C), c(deck.Seven, C), c(deck.Eight, C), c(deck.Nine, D))
	if IsStraightFlush(straight) {
		t.Errorf("IsStraightFlush(%v) = true, want false", straight)
	}
}

func TestCheckCombo(t *testing.T) {
	tests := []struct {
		name      string
		cards     deck.Deck
		wantValid bool
		wantType  ComboType
		wantPower deck.Card
	}{
		{"skip", d(), true, Skip, deck.Card{}},
		{"single", d(c(deck.Ace, H)), true, Single, c(deck.Ace, H)},
		{"pair takes higher suit", d(c(deck.Nine, S), c(deck.Nine, D)), true, Pair, c(deck.Nine, S)},
		{"triple", d(c(deck.Jack, H), c(deck.Jack, D), c(deck.Jack, C)), true, Triple, c(deck.Jack, H)},
		{"straight", d(c(deck.Nine, D), c(deck.Five, C), c(deck.Eight, H), c(deck.Six, S), c(deck.Seven, D)), true, Straight, c(deck.Nine, D)},
		{"flush", d(c(deck.Four, H), c(deck.Ace, H), c(deck.Nine, H), c(deck.Jack, H), c(deck.Seven, H)), true, Flush, c(deck.Ace, H)},
		{"full house power is triple", d(c(deck.King, S), c(deck.Six, D), c(deck.King, D), c(deck.Six, C), c(deck.Six, H)), true, FullHouse, c(deck.Six, H)},
		{"four of a kind power is quad", d(c(deck.Six, D), c(deck.Six, C), c(deck.Six, H), c(deck.Six, S), c(deck.King, D)), true, FourOfAKind, c(deck.Six, S)},
		{"straight flush over flush", d(c(deck.Five, C), c(deck.Nine, C), c(deck.Seven, C), c(deck.Eight, C), c(deck.Six, C)), true, StraightFlush, c(deck.Nine, C)},
		{"invalid pair", d(c(deck.Five, D), c(deck.Six, D)), false, None, deck.Card{}},
		{"invalid triple", d(c(deck.Five, D), c(deck.Five, C), c(deck.Six, D)), false, None, deck.Card{}},
		{"invalid four cards", d(c(deck.Five, D), c(deck.Five, C), c(deck.Five, H), c(deck.Five, S)), false, None, deck.Card{}},
		{"invalid five cards", d(c(deck.Four, D), c(deck.Six, C), c(deck.Nine, H), c(deck.Jack, S), c(deck.Ace, D)), false, None, deck.Card{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, got := CheckCombo(tt.cards)
			if valid != tt.wantValid {
				t.Fatalf("CheckCombo(%v) valid = %v, want %v", tt.cards, valid, tt.wantValid)
			}
			if got.Type != tt.wantType {
				t.Errorf("CheckCombo(%v) type = %v, want %v", tt.cards, got.Type, tt.wantType)
			}
			if valid && got.Power != tt.wantPower {
				t.Errorf("CheckCombo(%v) power = %v, want %v", tt.cards, got.Power, tt.wantPower)
			}
		})
	}
}

// mustCombo builds a combo from cards and fails the test if they are not a valid combo.
func mustCombo(t *testing.T, cards deck.Deck) Combo {
	t.Helper()
	valid, cb := CheckCombo(cards)
	if !valid {
		t.Fatalf("%v is not a valid combo", cards)
	}
	return cb
}

func TestCheckStrongerCombo(t *testing.T) {
	straightLow := d(c(deck.Four, D), c(deck.Five, C), c(deck.Six, H), c(deck.Seven, S), c(deck.Eight, D))
	straightHigh := d(c(deck.Five, D), c(deck.Six, C), c(deck.Seven, H), c(deck.Eight, S), c(deck.Nine, D))
	flushLow := d(c(deck.Four, H), c(deck.Seven, H), c(deck.Nine, H), c(deck.Jack, H), c(deck.King, H))
	flushHigh := d(c(deck.Four, C), c(deck.Seven, C), c(deck.Nine, C), c(deck.Jack, C), c(deck.Ace, C))
	fullHouseLow := d(c(deck.Eight, D), c(deck.Eight, C), c(deck.Eight, H), c(deck.King, S), c(deck.King, D))
	fullHouseHigh := d(c(deck.Nine, D), c(deck.Nine, C), c(deck.Nine, H), c(deck.Four, S), c(deck.Four, D))
	quadLow := d(c(deck.Four, D), c(deck.Four, C), c(deck.Four, H), c(deck.Four, S), c(deck.King, D))
	quadHigh := d(c(deck.Five, D), c(deck.Five, C), c(deck.Five, H), c(deck.Five, S), c(deck.Four, C))
	straightFlush := d(c(deck.Five, C), c(deck.Six, C), c(deck.Seven, C), c(deck.Eight, C), c(deck.Nine, C))

	tests := []struct {
		name   string
		toBeat deck.Deck
		played deck.Deck
		want   bool
	}{
		{"single higher rank", d(c(deck.King, S)), d(c(deck.Ace, D)), true},
		{"single lower rank", d(c(deck.Ace, D)), d(c(deck.King, S)), false},
		{"single same rank higher suit", d(c(deck.Ten, H)), d(c(deck.Ten, S)), true},
		{"single same rank lower suit", d(c(deck.Ten, H)), d(c(deck.Ten, C)), false},
		{"two beats ace", d(c(deck.Ace, S)), d(c(deck.Two, D)), true},
		{"pair higher rank", d(c(deck.Nine, S), c(deck.Nine, H)), d(c(deck.Ten, D), c(deck.Ten, C)), true},
		{"pair same rank decided by top suit", d(c(deck.Nine, D), c(deck.Nine, H)), d(c(deck.Nine, C), c(deck.Nine, S)), true},
		{"triple lower rank", d(c(deck.Jack, D), c(deck.Jack, C), c(deck.Jack, H)), d(c(deck.Ten, D), c(deck.Ten, C), c(deck.Ten, H)), false},
		{"length mismatch", d(c(deck.Four, D)), d(c(deck.Two, D), c(deck.Two, S)), false},
		{"straight higher straight", straightLow, straightHigh, true},
		{"straight lower straight", straightHigh, straightLow, false},
		{"flush beats straight", straightHigh, flushLow, true},
		{"straight loses to flush", flushLow, straightHigh, false},
		{"flush higher top card", flushLow, flushHigh, true},
		{"full house beats flush", flushHigh, fullHouseLow, true},
		{"full house compared by triple", fullHouseLow, fullHouseHigh, true},
		{"four of a kind beats full house", fullHouseHigh, quadLow, true},
		{"four of a kind compared by quad", quadLow, quadHigh, true},
		{"straight flush beats four of a kind", quadHigh, straightFlush, true},
		{"four of a kind loses to straight flush", straightFlush, quadHigh, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toBeat := mustCombo(t, tt.toBeat)
			played := mustCombo(t, tt.played)
			if got := CheckStrongerCombo(toBeat, played); got != tt.want {
				t.Errorf("CheckStrongerCombo(%v, %v) = %v, want %v", tt.toBeat, tt.played, got, tt.want)
			}
		})
	}
}
