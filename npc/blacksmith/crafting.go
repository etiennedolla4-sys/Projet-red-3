package blacksmith

import (
	"Projet-red-3/characters"
	"Projet-red-3/item"
	"fmt"
)

const craftingCost = 5

func CraftEquipment(p *characters.Character, choice int) {
	var equipment item.Item
	var requirements []string

	switch choice {
	case 1:
		equipment = item.NewAdventurerHat()
		requirements = []string{
			"Plume de Corbeau",
			"Cuir de Sanglier",
		}

	case 2:
		equipment = item.NewAdventurerTunic()
		requirements = []string{
			"Fourrure de Loup",
			"Fourrure de Loup",
			"Peau de Troll",
		}

	case 3:
		equipment = item.NewAdventurerBoots()
		requirements = []string{
			"Fourrure de Loup",
			"Cuir de Sanglier",
		}

	default:
		fmt.Println("Choix invalide.")
		return
	}

	if p.Gold < craftingCost {
		fmt.Println("\n❌ Vous n'avez pas assez de Gold.")
		return
	}

	if !hasMaterials(p, requirements) {
		fmt.Println("\n❌ Vous n'avez pas les matériaux nécessaires.")
		return
	}

	newInventorySize := len(p.Inventory) - len(requirements) + 1

	if newInventorySize > p.MaxInventory {
		fmt.Println("\n❌ Votre inventaire est plein.")
		return
	}

	removeMaterials(p, requirements)

	p.Gold -= craftingCost
	p.Inventory = append(p.Inventory, equipment)

	fmt.Println("\n🔨 Fabrication terminée !")
	fmt.Println("Objet fabriqué :", equipment.Name())
	fmt.Println("Gold restant :", p.Gold)
}

func hasMaterials(
	p *characters.Character,
	requirements []string,
) bool {
	available := make(map[string]int)

	for _, inventoryItem := range p.Inventory {
		available[inventoryItem.Name()]++
	}

	for _, requirement := range requirements {
		if available[requirement] <= 0 {
			return false
		}

		available[requirement]--
	}

	return true
}

func removeMaterials(
	p *characters.Character,
	requirements []string,
) {
	for _, requirement := range requirements {
		for i, inventoryItem := range p.Inventory {
			if inventoryItem.Name() == requirement {
				p.Inventory = append(
					p.Inventory[:i],
					p.Inventory[i+1:]...,
				)
				break
			}
		}
	}
}
