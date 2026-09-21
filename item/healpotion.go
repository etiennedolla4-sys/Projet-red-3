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

	fmt.Println("Vous utilisez une potion de soin !")
	fmt.Printf("PV : %d / %d\n", c.HP, c.MaxHP)
}
