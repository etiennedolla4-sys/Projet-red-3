package blacksmith

import (
	"Projet-red-3/characters"
	"fmt"
)

func EquipEquipment(p *characters.Character, index int) {
	if index < 0 || index >= len(p.Inventory) {
		fmt.Println("\n❌ Choix invalide.")
		return
	}

	object := p.Inventory[index]

	switch object.Type() {
	case "head":
		equipHead(p, index)

	case "chest":
		equipChest(p, index)

	case "feet":
		equipFeet(p, index)

	default:
		fmt.Println("\n❌ Cet objet n'est pas un équipement.")
	}
}

func equipHead(p *characters.Character, index int) {
	object := p.Inventory[index]

	if p.Equipment.Head != "" {
		fmt.Println("\n⚠️ Vous avez déjà un équipement de tête.")
		return
	}

	p.Equipment.Head = object.Name()
	p.MaxHP += 10

	removeInventoryItem(p, index)

	fmt.Println("\n🪖", object.Name(), "équipé !")
	fmt.Println("+10 PV maximum.")
}

func equipChest(p *characters.Character, index int) {
	object := p.Inventory[index]

	if p.Equipment.Chest != "" {
		fmt.Println("\n⚠️ Vous avez déjà un équipement de torse.")
		return
	}

	p.Equipment.Chest = object.Name()
	p.MaxHP += 25

	removeInventoryItem(p, index)

	fmt.Println("\n🛡", object.Name(), "équipée !")
	fmt.Println("+25 PV maximum.")
}

func equipFeet(p *characters.Character, index int) {
	object := p.Inventory[index]

	if p.Equipment.Feet != "" {
		fmt.Println("\n⚠️ Vous avez déjà un équipement aux pieds.")
		return
	}

	p.Equipment.Feet = object.Name()
	p.MaxHP += 15

	removeInventoryItem(p, index)

	fmt.Println("\n🥾", object.Name(), "équipées !")
	fmt.Println("+15 PV maximum.")
}

func removeInventoryItem(
	p *characters.Character,
	index int,
) {
	p.Inventory = append(
		p.Inventory[:index],
		p.Inventory[index+1:]...,
	)
}
