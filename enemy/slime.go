package enemy

func NewSlime() Monster {
	return Monster{
		Name:       "Slime",
		HP:         50,
		MaxHP:      50,
		Attack:     10,
		Defense:    5,
		Initiative: 3,
	}
}
