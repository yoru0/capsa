package menu

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yoru0/capsa-custom/internal/game"
)

func MainMenu() {
	var input int

	for {
		fmt.Println("十三")
		fmt.Println("1. New Game")
		fmt.Println("2. Exit")
		fmt.Print(">> ")

		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')

		switch input {
		case 1:
			fmt.Println("\nStarting a new game...")
			game.StartGame()
			continue
		case 2:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid input. Please enter 1 or 2.")
		}

		fmt.Println('\n')
	}
}
