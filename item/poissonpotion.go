package item

import (
	"fmt"
	"time"

	"Projet-red-3/characters"
)

func TakePoisonPotion(c *characters.Character) {
	for i, object := range c.Inventory {
		if object == "Potion de poison" {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			break
		}
	}

	fmt.Println("Vous utilisez une potion de poison !")

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		c.HP -= 10

		if c.HP < 0 {
			c.HP = 0
		}

		fmt.Printf("PV : %d / %d\n", c.HP, c.MaxHP)
	}
}
