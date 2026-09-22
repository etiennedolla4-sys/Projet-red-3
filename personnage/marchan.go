package personnage

import (
	"fmt"

	"Projet-red-3/characters"
	"Projet-red-3/item"
)

func Merchant(p *characters.Character) {
	fmt.Println("\n=== MARCHAND ===")
	fmt.Println("Gold :", p.Gold)
	fmt.Println("1. Potion de vie - 3 Gold")
	fmt.Println("2. Potion de poison - 6 Gold")
	fmt.Println("3. Livre de Sort : Boule de Feu - 25 Gold")
	fmt.Println("4. Fourrure de Loup - 4 Gold")
	fmt.Println("5. Peau de Troll - 7 Gold")
	fmt.Println("6. Cuir de Sanglier - 3 Gold")
	fmt.Println("7. Plume de Corbeau - 1 Gold")
	fmt.Println("0. Retour")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	var object string
	var price int

	switch choice {
	case 1:
		object = "Potion de vie"
		price = 3

	case 2:
		object = "Potion de poison"
		price = 6

	case 3:
		object = "Livre de Sort : Boule de Feu"
		price = 25

	case 4:
		object = "Fourrure de Loup"
		price = 4

	case 5:
		object = "Peau de Troll"
		price = 7

	case 6:
		object = "Cuir de Sanglier"
		price = 3

	case 7:
		object = "Plume de Corbeau"
		price = 1

	case 0:
		return

	default:
		fmt.Println("Choix invalide")
		return
	}

	if p.Gold < price {
		fmt.Println("Vous n'avez pas assez de Gold !")
		return
	}

	if !item.AddInventory(p, object) {
		return
	}

	p.Gold -= price

	fmt.Println("Vous avez acheté :", object)
	fmt.Println("Gold restant :", p.Gold)
}
