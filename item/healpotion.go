package item

import (
	"fmt"

	"Projet-red-3/characters"
)

func TakePotion(c *characters.Character) {
	c.HP += 50

	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}

	for i, object := range c.Inventory {
		if object == "Potion de vie" {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			break
		}
	}

	fmt.Println("Vous utilisez une potion de vie !")
	fmt.Printf("PV : %d / %d\n", c.HP, c.MaxHP)
}
