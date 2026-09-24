package characters

import "Projet-red-3/item"

type berserker struct {
	Character
}

func NewBerserker(name string) berserker {
	return berserker{
		Character{
		
			Name:         name,
			Class:        "berserker",
			Description:  "Truc",
			HP:           55,
			MP:           0,
			BaseAttack:   15,
			BaseDefense:  10,
			Initiative:   6,
			Inventory:    []item.Item{},
			MaxInventory: 10,
			MaxHP:        200,
			MaxMP:        0,
			Crit:         0,
			Lvl:          1,
			Gold:         100,
			Skills:       []string{},
		},
	}
}
