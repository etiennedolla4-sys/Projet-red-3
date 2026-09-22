package showinfo

import (
	"Projet-red-3/characters"
	"Projet-red-3/item"
	"Projet-red-3/sort"
	"fmt"
)

func DisplayInventory(p *characters.Character) {
	fmt.Println("\n=== INVENTAIRE DU PERSONNAGE ===")

	if len(p.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
	} else {
		for i, object := range p.Inventory {
			fmt.Printf("%d. %s\n", i+1, object)
		}
	}

	fmt.Println("0. Retour")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	if choice > 0 && choice <= len(p.Inventory) {
		switch p.Inventory[choice-1] {

		case "Potion de vie":
			item.TakePotion(p)

		case "Potion de mana":
			item.TakeManaPotion(p)

		case "Potion de poison":
			item.TakePoisonPotion(p)

		case "Livre de Sort : Boule de Feu":
			sort.SpellBook(p)

			for i, object := range p.Inventory {
				if object == "Livre de Sort : Boule de Feu" {
					p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
					break
				}
			}

		case "Chapeau de l'aventurier":
			item.EquipItem(p, p.Inventory[choice-1])

		case "Tunique de l'aventurier":
			item.EquipItem(p, p.Inventory[choice-1])

		case "Bottes de l'aventurier":
			item.EquipItem(p, p.Inventory[choice-1])
		}
	}
}
