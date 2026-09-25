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
			Description:  "Le Vampire est une créature mystérieuse qui se nourrit de ses adversaires pour récupérer ses forces.",
			HP:           130,
			MP:           80,
			BaseAttack:   20,
			BaseDefense:  10,
			Initiative:   3,
			Inventory:    []item.Item{},
			MaxInventory: 10,
			HasBag:       false,
			MaxHP:        200,
			MaxMP:        100,
			Crit:         0,
			Lvl:          1,
			Gold:         100,
		},
	}
}
