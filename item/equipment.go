package item

import (
	"Projet-red-3/characters"
	"fmt"
)

func EquipItem(p *characters.Character, object string) {
	switch object {
	case "Chapeau de l'aventurier":
		p.Equipment.Head = object
		p.MaxHP += 10

	case "Tunique de l'aventurier":
		p.Equipment.Torso = object
		p.MaxHP += 25

	case "Bottes de l'aventurier":
		p.Equipment.Feet = object
		p.MaxHP += 15

	default:
		return
	}

	for i, item := range p.Inventory {
		if item == object {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			break
		}
	}

	fmt.Println("Vous équipez :", object)
	fmt.Printf("PV maximum : %d\n", p.MaxHP)
}
