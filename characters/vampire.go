package characters

type Vampire struct {
	Character
}

func NewVampire() Vampire {
	return Vampire{
		Character{
			Name:         "Vampire",
			HP:           60,
			Mana:         10,
			BaseAttack:   20,
			BaseDefense:  10,
			Initiative:   3,
			Inventory:    []string{},
			MaxInventory: 10,
			MaxHP:        200,
			MaxMP:        100,
			Crit:         0,
			Lvl:          1,
		},
	}
}
