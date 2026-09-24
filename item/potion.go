package item

type Potion struct {
	ItemName        string
	ItemDescription string
	ItemType        string
}

func (p Potion) Name() string {
	return p.ItemName
}

func (p Potion) Description() string {
	return p.ItemDescription
}

func (p Potion) Type() string {
	return p.ItemType
}
