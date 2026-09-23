package item

type HealPotion struct {
	ItemName        string
	ItemDescription string
	HealAmount      int
}

type ManaPotion struct {
	ItemName        string
	ItemDescription string
	ManaAmount      int
}

func (p HealPotion) Name() string {
	return p.ItemName
}

func (p HealPotion) Description() string {
	return p.ItemDescription
}

func (p HealPotion) Type() string {
	return "Potion de vie"
}

func (p ManaPotion) Name() string {
	return p.ItemName
}

func (p ManaPotion) Description() string {
	return p.ItemDescription
}

func (p ManaPotion) Type() string {
	return "Potion de mana"
}
