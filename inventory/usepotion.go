package inventory

import (
	"Projet-red-3/characters"
	"Projet-red-3/item"
)

func UseHealthPotion(character *characters.Character, index int) bool {
	if index < 0 || index >= len(character.Inventory) {
		return false
	}

	potion, ok := character.Inventory[index].(item.Potion)
	if !ok {
		return false
	}

	if potion.Type() != "Potion" || potion.Name() != "Potion de vie" {
		return false
	}

	character.HP += 50

	if character.HP > character.MaxHP {
		character.HP = character.MaxHP
	}

	character.Inventory = append(
		character.Inventory[:index],
		character.Inventory[index+1:]...,
	)

	return true
}
