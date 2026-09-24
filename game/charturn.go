package game

import (
	"Projet-red-3/characters"
	"Projet-red-3/item"
	"fmt"
)

func CharTurn(p *characters.Character, goblin *characters.Monster, tour int) {
	fmt.Println("\n=== TOUR DU JOUEUR ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		damage := 5
		goblin.HP -= damage

		if goblin.HP < 0 {
			goblin.HP = 0
		}

		fmt.Printf("%s inflige %d dégâts à %s\n",
			p.Name,
			damage,
			goblin.Name,
		)

		fmt.Printf("PV : %d / %d\n", goblin.HP, goblin.MaxHP)

	case 2:
		UseCombatItem(p)

	default:
		fmt.Println("Choix invalide")
	}
}

func UseCombatItem(p *characters.Character) {
	if len(p.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	fmt.Println("\n=== INVENTAIRE ===")

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

	if choice < 1 || choice > len(p.Inventory) {
		fmt.Println("Choix invalide")
		return
	}

	switch p.Inventory[choice-1] {
	case "Potion de vie":
		item.TakePotion(p)

	case "Potion de mana":
		item.TakeManaPotion(p)

	case "Potion de poison":
		item.TakePoisonPotion(p)

	default:
		fmt.Println("Cet objet ne peut pas être utilisé pendant le combat.")
	}
}
