package item

import (
	"fmt"

	"Projet-red-3/characters"
)

func AddInventory(p *characters.Character, object string) bool {
	if len(p.Inventory) >= p.MaxInventory {
		fmt.Println("Votre inventaire est plein !")
		return false
	}

	p.Inventory = append(p.Inventory, object)
	return true
}
