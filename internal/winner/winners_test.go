package winner

import (
	"fmt"
	"testing"
)

func TestLastPlace(t *testing.T) {
	tests := []struct {
		name    string
		players []string
		winners Winner
		want    string
	}{
		{
			name:    "four players",
			players: []string{"Player 1", "Player 2", "Player 3", "Player 4"},
			winners: Winner{"Player 2", "Player 4", "Player 1"},
			want:    "Player 3",
		},
		{
			name:    "three players",
			players: []string{"Player 1", "Player 2", "Player 3"},
			winners: Winner{"Player 1", "Player 2"},
			want:    "Player 3",
		},
		{
			name:    "custom names",
			players: []string{"Alice", "Bob", "Carol", "Dave"},
			winners: Winner{"Alice", "Carol", "Dave"},
			want:    "Bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safeLastPlace(tt.winners, tt.players)
			if err != nil {
				t.Fatalf("lastPlace(%v) panicked: %v", tt.players, err)
			}
			if got != tt.want {
				t.Errorf("lastPlace(%v) = %q, want %q", tt.players, got, tt.want)
			}
		})
	}
}

// safeLastPlace turns a panic into an error so one bad case doesn't stop the other tests.
func safeLastPlace(w Winner, players []string) (got string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return w.lastPlace(players), nil
}
