package characters

import "Projet-red-3/item"

type Berserker struct {
	Character
}

func NewBerserker(name string) Berserker {
	return Berserker{
		Character{

			Name:         name,
			Class:        "Berserker",
			Description:  "Le Berserker est un combattant puissant qui privilégie la force brute et les attaques dévastatrices.",
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
