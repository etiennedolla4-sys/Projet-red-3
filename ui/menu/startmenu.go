package menu

import (
	"Projet-red-3/game"
	exploration "Projet-red-3/game/exploration"
	"Projet-red-3/utils"
	"fmt"
	"os"
)

func StartMenu() {
	var choice int

	for {
		utils.ClearTerminal()

		fmt.Println(`
╔══════════════════════════════════════════════════════════════════╗
║                                                                  ║
║  ██╗   ██╗███████╗██╗   ██╗██████╗                               ║
║  ██║   ██║██╔════╝╚██╗ ██╔╝██╔══██╗                              ║
║  ██║   ██║█████╗   ╚████╔╝ ██████╔╝                              ║
║  ╚██╗ ██╔╝██╔══╝    ╚██╔╝  ██╔══██╗                              ║
║   ╚████╔╝ ███████╗   ██║   ██║  ██║                              ║
║    ╚═══╝  ╚══════╝   ╚═╝   ╚═╝  ╚═╝                              ║
║                                                                  ║
║              ─── Un monde oublié t'attend ───                    ║
║                                                                  ║
╠══════════════════════════════════════════════════════════════════╣
║                                                                  ║
║                     1. Start Game                                ║
║                     2. Exit                                      ║
║                                                                  ║
╚══════════════════════════════════════════════════════════════════╝`)

		fmt.Scan(&choice)

		switch choice {
		case 1:
			utils.ClearTerminal()

			p := game.Init()
			exploration.Start(&p)
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
