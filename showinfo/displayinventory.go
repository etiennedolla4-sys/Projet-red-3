package showinfo

import (
	"Projet-red-3/characters"
	"fmt"
)

func DisplayInventory(p characters.Character) {
	fmt.Println("\n=== INVENTAIRE DU PERSONNAGE ===")
	for i, item := range p.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}
