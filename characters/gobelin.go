package characters

import "Projet-red-3/item"

type Gobelin struct {
	Character
}

func NewGobelin(name string) Gobelin {
	return Gobelin{
		Character{
			Name:         name,
			Class:        "Gobelin",
			Description:  "Petit et rusé, le Gobelin mise sur sa rapidité et son agilité pour surprendre ses ennemis.",
			HP:           110,
			MP:           50,
			BaseAttack:   20,
			BaseDefense:  5,
			Initiative:   5,
			Inventory:    []item.Item{},
			MaxInventory: 10,
			HasBag:       false,
			MaxHP:        150,
			MaxMP:        150,
			Crit:         0,
			Lvl:          1,
			Gold:         120,
			Skills:       []string{},
		},
	}
}
