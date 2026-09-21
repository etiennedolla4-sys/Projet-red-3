package menu

import (
	"fmt"
	"os"
)

func StartMenu() {
	var choice int
	PrintSlow.PrintSlow("Bienvenue")
	PrintSlow.PrintSlow("Lore tah zebi Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur \n")
	fmt.Println("1. Start Game")
	fmt.Println("2. Load Game")
	fmt.Println("3. Exit")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		init() //init
	case 2:
		//aucune idee
	case 3:
		fmt.Println("Ciao")
		os.Exit(1)

	}
}
