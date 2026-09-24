package item

func NewHealthPotion() Potion {
	return Potion{
		ItemName:        "Potion de vie",
		ItemDescription: "Restaure 50 PV.",
		ItemType:        "Potion",
	}
}

func NewManaPotion() Potion {
	return Potion{
		ItemName:        "Potion de mana",
		ItemDescription: "Restaure 30 MP.",
		ItemType:        "Potion",
	}
}

func NewPoisonPotion() Potion {
	return Potion{
		ItemName:        "Potion de poison",
		ItemDescription: "Inflige des dégâts.",
		ItemType:        "Potion",
	}
}
