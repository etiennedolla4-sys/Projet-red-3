package showinfo

import (
	"Projet-red-3/characters"
	"Projet-red-3/inventory"
	"fmt"
)

func DisplayInventory(p *characters.Character) {
	for {
		fmt.Println("\n========== INVENTAIRE ==========")

		if len(p.Inventory) == 0 {
			fmt.Println("Votre inventaire est vide.")
			fmt.Println("0. Retour")
			fmt.Print("Votre choix : ")

			var choice int
			fmt.Scan(&choice)

			return
		}

		for i, object := range p.Inventory {
			fmt.Printf("%d. %s\n", i+1, object.Name())
		}

		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			return
		}

		index := choice - 1

		if index < 0 || index >= len(p.Inventory) {
			fmt.Println("Choix invalide.")
			continue
		}

		object := p.Inventory[index]

		switch object.Name() {
		case "Potion de vie":
			if inventory.UseHealthPotion(p, index) {
				fmt.Println("\nVous utilisez une Potion de vie.")
				fmt.Printf("PV : %d / %d\n", p.HP, p.MaxHP)
			} else {
				fmt.Println("Impossible d'utiliser cette potion.")
			}

		default:
			fmt.Println("Cet objet ne peut pas encore être utilisé.")
		}
	}
}
