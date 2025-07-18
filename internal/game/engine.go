package game

import (
	"fmt"

	"github.com/yoru0/capsa-custom/internal/combo"
	"github.com/yoru0/capsa-custom/internal/design"
)

func StartGame() {
	g := NewGame(4)
	fmt.Printf("%s play first\n\n", g.Players[g.CurrIndex].Name)
	for len(g.Players) > 1 {
		for g.PlayerSkipped < (len(g.Players) - 1) {
			// TODO
			if g.Players[g.CurrIndex].Skip {
				fmt.Printf("%s skipped\n", g.Players[g.CurrIndex].Name)
				g.NextPlayerTurn()
			}

			fmt.Printf("%s [%d card(s)]:\n", g.Players[g.CurrIndex].Name, len(g.Players[g.CurrIndex].Hand))
			g.Players[g.CurrIndex].ShowPlayerHand()

			g.ComboPlayed = g.Players[g.CurrIndex].PickCard(g.ComboToBeat)

			if g.ComboPlayed.Type == combo.Skip {
				g.Players[g.CurrIndex].Skip = true
				g.PlayerSkipped++
			} else {
				g.ComboToBeat = g.ComboPlayed
			}
			

			g.CheckGameDetails(g.CurrIndex)

			design.PressEnterToContinue()

			g.NextPlayerTurn()

		}
		g.Round++
		g.ResetPlayerSkips()
		g.ResetLastCombo()
	}
}
