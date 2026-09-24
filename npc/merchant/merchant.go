package npc

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

	var object item.BasicItem
	var price int

	switch choice {
	case 1:
		object = item.BasicItem{
			ItemName:        "Potion de vie",
			ItemDescription: "Restaure des points de vie.",
			ItemType:        "Potion",
		}
		price = 3

	case 2:
		object = item.BasicItem{
			ItemName:        "Potion de poison",
			ItemDescription: "Inflige des dégâts pendant quelques secondes.",
			ItemType:        "Potion",
		}
		price = 6

	case 3:
		object = item.BasicItem{
			ItemName:        "Livre de Sort : Boule de Feu",
			ItemDescription: "Permet d'apprendre le sort Boule de Feu.",
			ItemType:        "Livre",
		}
		price = 25

	case 4:
		object = item.BasicItem{
			ItemName:        "Fourrure de Loup",
			ItemDescription: "Une fourrure provenant d'un loup.",
			ItemType:        "Ressource",
		}
		price = 4

	case 5:
		object = item.BasicItem{
			ItemName:        "Peau de Troll",
			ItemDescription: "Une peau épaisse provenant d'un troll.",
			ItemType:        "Ressource",
		}
		price = 7

	case 6:
		object = item.BasicItem{
			ItemName:        "Cuir de Sanglier",
			ItemDescription: "Du cuir provenant d'un sanglier.",
			ItemType:        "Ressource",
		}
		price = 3

	case 7:
		object = item.BasicItem{
			ItemName:        "Plume de Corbeau",
			ItemDescription: "Une plume noire de corbeau.",
			ItemType:        "Ressource",
		}
		price = 1

	case 0:
		return

	default:
		fmt.Println("Choix invalide.")
		return
	}

	if p.Gold < price {
		fmt.Println("Vous n'avez pas assez de Gold !")
		return
	}

	if len(p.Inventory) >= p.MaxInventory {
		fmt.Println("Votre inventaire est plein !")
		return
	}

	p.Gold -= price
	p.Inventory = append(p.Inventory, object)

	fmt.Println()
	fmt.Println("Vous avez acheté :", object.Name())
	fmt.Println("L'objet a été ajouté à votre inventaire.")
	fmt.Println("Gold restant :", p.Gold)
}
