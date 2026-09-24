package item

type AdventurerTunic struct{}

func NewAdventurerTunic() Item {
	return AdventurerTunic{}
}

func (AdventurerTunic) Name() string {
	return "Tunique de l'aventurier"
}

func (AdventurerTunic) Description() string {
	return "Une tunique renforcée qui augmente les PV maximum de 25."
}

func (AdventurerTunic) Type() string {
	return "chest"
}
