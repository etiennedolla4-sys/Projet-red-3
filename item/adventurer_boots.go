package item

type AdventurerBoots struct{}

func NewAdventurerBoots() Item {
	return AdventurerBoots{}
}

func (AdventurerBoots) Name() string {
	return "Bottes de l'aventurier"
}

func (AdventurerBoots) Description() string {
	return "Des bottes solides qui augmentent les PV maximum de 15."
}

func (AdventurerBoots) Type() string {
	return "feet"
}
