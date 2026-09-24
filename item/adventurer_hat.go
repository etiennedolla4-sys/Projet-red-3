package item

type AdventurerHat struct{}

func NewAdventurerHat() Item {
	return AdventurerHat{}
}

func (AdventurerHat) Name() string {
	return "Chapeau de l'aventurier"
}

func (AdventurerHat) Description() string {
	return "Un chapeau solide qui augmente les PV maximum de 10."
}

func (AdventurerHat) Type() string {
	return "head"
}
