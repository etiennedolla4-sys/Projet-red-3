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

func NewRevivalPotion() Potion {
	return Potion{
		ItemName:        "Potion de réanimation",
		ItemDescription: "Vous réanime automatiquement une seule fois avec 50% de vos PV maximum.",
		ItemType:        "Potion",
	}
}
func NewFireballBook() SpellBook {
	return SpellBook{
		ItemName:        "Livre de Sort : Boule de Feu",
		ItemDescription: "Un livre contenant le sort Boule de Feu.",
		ItemType:        "Livre de Sort",
	}
}

func NewWolfFur() Material {
	return Material{
		ItemName:        "Fourrure de Loup",
		ItemDescription: "Une fourrure épaisse provenant d'un loup.",
		ItemType:        "Matériau",
	}
}

func NewTrollSkin() Material {
	return Material{
		ItemName:        "Peau de Troll",
		ItemDescription: "Une peau résistante provenant d'un troll.",
		ItemType:        "Matériau",
	}
}

func NewBoarLeather() Material {
	return Material{
		ItemName:        "Cuir de Sanglier",
		ItemDescription: "Un cuir robuste provenant d'un sanglier.",
		ItemType:        "Matériau",
	}
}

func NewCrowFeather() Material {
	return Material{
		ItemName:        "Plume de Corbeau",
		ItemDescription: "Une plume noire de corbeau.",
		ItemType:        "Matériau",
	}
}
