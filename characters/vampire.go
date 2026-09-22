package characters

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
			Inventory:    []string{},
			MaxInventory: 10,
			MaxHP:        200,
			MaxMP:        100,
			Crit:         0,
			Lvl:          1,
			Gold:         100,
			Skills:       []string{"Coup de poing", "Saignée"},
		},
	}
}
