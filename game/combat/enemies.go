package combat

import "math/rand"

func GetEnemies(depth int) []Enemy {
	switch depth {
	case 1:
		return []Enemy{
			{
				Name:       "Slime",
				HP:         80,
				MaxHP:      80,
				Attack:     15,
				Defense:    3,
				Initiative: 3,
			},
			{
				Name:       "Petit Gobelin",
				HP:         60,
				MaxHP:      60,
				Attack:     18,
				Defense:    2,
				Initiative: 5,
			},
		}

	case 2:
		return []Enemy{
			{
				Name:       "Gobelin des ruines",
				HP:         90,
				MaxHP:      90,
				Attack:     20,
				Defense:    5,
				Initiative: 6,
			},
			{
				Name:       "Araignée ancienne",
				HP:         70,
				MaxHP:      70,
				Attack:     25,
				Defense:    4,
				Initiative: 8,
			},
		}

	case 3:
		return []Enemy{
			{
				Name:       "Loup sauvage",
				HP:         100,
				MaxHP:      100,
				Attack:     25,
				Defense:    7,
				Initiative: 10,
			},
			{
				Name:       "Prédateur de la forêt",
				HP:         120,
				MaxHP:      120,
				Attack:     30,
				Defense:    8,
				Initiative: 9,
			},
		}

	case 4:
		return []Enemy{
			{
				Name:       "Golem des cavernes",
				HP:         150,
				MaxHP:      150,
				Attack:     30,
				Defense:    12,
				Initiative: 4,
			},
			{
				Name:       "Chauve-souris géante",
				HP:         110,
				MaxHP:      110,
				Attack:     35,
				Defense:    8,
				Initiative: 15,
			},
		}

	case 5:
		return []Enemy{
			{
				Name:       "Monstre des profondeurs",
				HP:         180,
				MaxHP:      180,
				Attack:     40,
				Defense:    15,
				Initiative: 12,
			},
			{
				Name:       "Gardien abyssal",
				HP:         220,
				MaxHP:      220,
				Attack:     45,
				Defense:    18,
				Initiative: 10,
			},
		}

	default:
		return []Enemy{
			NewSlime(),
			NewSlime(),
		}
	}
}

func GetRandomEnemy(depth int) Enemy {
	enemies := GetEnemies(depth)

	return enemies[rand.Intn(len(enemies))]
}
