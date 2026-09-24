package showinfo

import (
	"Projet-red-3/characters"
	"fmt"
)

func DisplayInventory(p characters.Character) {
	fmt.Println("\n========== INVENTAIRE ==========")

	if len(p.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	for i, object := range p.Inventory {
		fmt.Printf("%d. %s\n", i+1, object.Name())
	}

	fmt.Printf("\n%d/%d objets\n", len(p.Inventory), p.MaxInventory)
}
