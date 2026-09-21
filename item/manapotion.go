package item

import (
	"fmt"

	"Projet-red-3/characters"
)

func TakeManaPotion(c *characters.Character) {
	c.MP += 30

	if c.MP > c.MaxMP {
		c.MP = c.MaxMP
	}

	for i, object := range c.Inventory {
		if object == "Potion de mana" {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			break
		}
	}

	fmt.Println("Vous utilisez une potion de mana !")
	fmt.Printf("PM : %d / %d\n", c.MP, c.MaxMP)
}
