package menu

import (
	"fmt"

	"Projet-red-3/characters"
	"Projet-red-3/game"
	"Projet-red-3/npc"
	"Projet-red-3/showinfo"
)

func PauseMenu(p *characters.Character) {
	for {
		fmt.Println()
		fmt.Println("=== MENU ===")
		fmt.Println("1. Informations du personnage")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Entraînement")
		fmt.Println("5. Quitter")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			showinfo.DisplayInfo(*p)
		case 2:
			showinfo.DisplayInventory(p)
		case 3:
			npc.Merchant(p)
		case 4:
			game.TrainingFight(p)
		case 5:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
