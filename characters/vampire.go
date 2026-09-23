package characters

import "Projet-red-3/item"

type Vampire struct {
	Character
}

func NewVampire(name string) Vampire {
	return Vampire{
		Character{
			Name:         name,
			Class:        "Vampire",
			Description:  "Truc",
			HP:           60,
			MP:           10,
			BaseAttack:   20,
			BaseDefense:  10,
			Initiative:   3,
			Inventory:    []item.Item{},
			MaxInventory: 10,
			MaxHP:        200,
			MaxMP:        100,
			Crit:         0,
			Lvl:          1,
			Gold:         100,
		},
	}
}
