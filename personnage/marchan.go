package personnage

import (
	"fmt"

	"Projet-red-3/characters"
)

func Merchant(p *characters.Character) {
	fmt.Println("\n=== MARCHAND ===")
	fmt.Println("1. Potion de vie - Gratuit")
	fmt.Println("2. Potion de poison - Gratuit")
	fmt.Println("0. Retour")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if len(p.Inventory) >= p.MaxInventory {
			fmt.Println("Votre inventaire est plein !")
			return
		}

		p.Inventory = append(p.Inventory, "Potion de vie")
		fmt.Println("Vous avez acheté : Potion de vie")

	case 2:
		if len(p.Inventory) >= p.MaxInventory {
			fmt.Println("Votre inventaire est plein !")
			return
		}

		p.Inventory = append(p.Inventory, "Potion de poison")
		fmt.Println("Vous avez acheté : Potion de poison")

	case 0:
		return

	default:
		fmt.Println("Choix invalide")
	}
}
