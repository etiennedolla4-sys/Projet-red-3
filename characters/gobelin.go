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
			HP:           50,
			MP:           0,
			BaseAttack:   20,
			BaseDefense:  5,
			Initiative:   5,
			Inventory:    []item.Item{},
			MaxInventory: 10,
			MaxHP:        150,
			MaxMP:        0,
			Crit:         0,
			Lvl:          1,
			Gold:         120,
			Skills:       []string{},
		},
	}
}
