package characters

type Gobelin struct {
	Character
}

func NewGobelin(name string) Gobelin {
	return Gobelin{
		Character{
			Name:         name,
			Class:        "Gobelin",
			Description:  "Truc",
			HP:           50,
			MP:           0,
			BaseAttack:   20,
			BaseDefense:  5,
			Initiative:   5,
			Inventory:    []string{},
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
