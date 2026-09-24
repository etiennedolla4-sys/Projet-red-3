package menu

import (
	"Projet-red-3/game"
	"Projet-red-3/printslow"
	"fmt"
	"os"
)

func StartMenu() {
	var choice int

	printslow.PrintSlow("Bienvenue\n")

	fmt.Println("1. Start Game")
	fmt.Println("2. Exit")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		p := game.CharCreation()
		game.DetectP(&p)

	case 2:
		fmt.Println("Ciao")
		os.Exit(0)
	}
}
