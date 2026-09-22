package showinfo

import (
	"fmt"

	"Projet-red-3/characters"
	"Projet-red-3/item"
)

func DisplayInventory(p *characters.Character) {
	fmt.Println("\n=== INVENTAIRE DU PERSONNAGE ===")

	for i, object := range p.Inventory {
		fmt.Printf("%d. %s\n", i+1, object)
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
		}
	}
}
