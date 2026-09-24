package menu

import (
	"Projet-red-3/game"
	"fmt"
	"os"
)

func StartMenu() {
	var choice int

	fmt.Println("1. Start Game")
	fmt.Println("2. Exit")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		p := game.Init()
		DetecP(&p)

	case 2:
		fmt.Println("Ciao")
		os.Exit(0)
	}

}
