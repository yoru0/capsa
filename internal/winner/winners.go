package winner

import (
	"fmt"
	"sort"
	"strconv"
)

type Winner struct {
	PlayerName string
	Position   int
}

func AddWinner(winnerList []Winner, playerName string, position int) []Winner {
	return append(winnerList, Winner{PlayerName: playerName, Position: position})
}

func ShowWinners(winners []Winner) {
	fmt.Println("Winners:")
	sort.Slice(winners, func(i, j int) bool {
		return winners[i].Position < winners[j].Position
	})

	for _, w := range winners {
		fmt.Printf("%d. %s\n", w.Position, w.PlayerName)
	}
}

func LastPlace(winners []Winner) string {
	total := 10
	for i := 0; i < len(winners); i++ {
		str := winners[i].PlayerName[len(winners[i].PlayerName)-1:]
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		total -= i
	}
	lastPlaceName := "Player " + strconv.Itoa(total)
	return lastPlaceName
}
