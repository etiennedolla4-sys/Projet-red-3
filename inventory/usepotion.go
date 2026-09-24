package inventory

import (
	"Projet-red-3/characters"
	"fmt"
	"time"
)

func UseHealthPotion(character *characters.Character, index int) bool {
	if index < 0 || index >= len(character.Inventory) {
		return false
	}

	object := character.Inventory[index]

	if object.Name() != "Potion de vie" {
		return false
	}

	character.HP += 50

	if character.HP > character.MaxHP {
		character.HP = character.MaxHP
	}

	removeItem(character, index)

	return true
}

func UsePoisonPotion(character *characters.Character, index int) bool {
	if index < 0 || index >= len(character.Inventory) {
		return false
	}

	object := character.Inventory[index]

	if object.Name() != "Potion de poison" {
		return false
	}

	removeItem(character, index)

	fmt.Println("\n☠️ Vous utilisez une potion de poison !")

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)

		character.HP -= 10

		if character.HP < 0 {
			character.HP = 0
		}

		fmt.Printf("PV : %d / %d\n", character.HP, character.MaxHP)

		if character.HP <= 0 {
			break
		}
	}

	return true
}

func removeItem(character *characters.Character, index int) {
	character.Inventory = append(
		character.Inventory[:index],
		character.Inventory[index+1:]...,
	)
}
