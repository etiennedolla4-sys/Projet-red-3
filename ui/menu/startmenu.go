package menu

import (
	"Projet-red-3/game"
	"Projet-red-3/utils"
	"fmt"
	"os"
)

func StartMenu() {
	var choice int

	for {
		utils.ClearTerminal()

		fmt.Println("╔════════════════════════════════════════╗")
		fmt.Println("║                                        ║")
		fmt.Println("║                V E Y R                 ║")
		fmt.Println("║                                        ║")
		fmt.Println("║          RPG CLI ADVENTURE             ║")
		fmt.Println("║                                        ║")
		fmt.Println("╠════════════════════════════════════════╣")
		fmt.Println("║                                        ║")
		fmt.Println("║            1. Start Game               ║")
		fmt.Println("║            2. Exit                     ║")
		fmt.Println("║                                        ║")
		fmt.Println("╚════════════════════════════════════════╝")
		fmt.Print("\n> ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			utils.ClearTerminal()

			p := game.Init()
			DetecP(&p)

			return

		case 2:
			utils.ClearTerminal()
			fmt.Println("\nMerci d'avoir joué à VEYR !")
			os.Exit(0)

		default:
			fmt.Println("\n❌ Choix invalide.")
			fmt.Println("Appuyez sur Entrée pour continuer...")
			fmt.Scanln()
			fmt.Scanln()
		}
	}
}
