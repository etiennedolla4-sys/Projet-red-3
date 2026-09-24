package characters

type Monster struct {
	Name       string
	MaxHP      int
	HP         int
	BaseAttack int
}

func InitGoblin() Monster {
	return Monster{
		Name:       "Gobelin d'entrainement",
		MaxHP:      40,
		HP:         40,
		BaseAttack: 5,
	}
}
