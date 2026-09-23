package npc

import (
	"fmt"

	"Projet-red-3/characters"
	"Projet-red-3/item"
)

func Blacksmith(p *characters.Character) {
	fmt.Println("\n=== FORGERON ===")
	fmt.Println("Gold :", p.Gold)
	fmt.Println("1. Chapeau de l'aventurier - 5 Gold")
	fmt.Println("2. Tunique de l'aventurier - 5 Gold")
	fmt.Println("3. Bottes de l'aventurier - 5 Gold")
	fmt.Println("0. Retour")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	var equipment string

	switch choice {
	case 1:
		equipment = "Chapeau de l'aventurier"

	case 2:
		equipment = "Tunique de l'aventurier"

	case 3:
		equipment = "Bottes de l'aventurier"

	case 0:
		return

	default:
		fmt.Println("Choix invalide")
		return
	}

	if p.Gold < 5 {
		fmt.Println("Vous n'avez pas assez de Gold !")
		return
	}

	if !item.AddInventory(p, equipment) {
		return
	}

	p.Gold -= 5

	fmt.Println("Vous avez fabriqué :", equipment)
	fmt.Println("Gold restant :", p.Gold)
}
