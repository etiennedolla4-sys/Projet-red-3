package item

type SpellBook struct {
	ItemName        string
	ItemDescription string
	ItemType        string
}

func (s SpellBook) Name() string {
	return s.ItemName
}

func (s SpellBook) Description() string {
	return s.ItemDescription
}

func (s SpellBook) Type() string {
	return s.ItemType
}
