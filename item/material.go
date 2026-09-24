package item

type Material struct {
	ItemName        string
	ItemDescription string
	ItemType        string
}

func (m Material) Name() string {
	return m.ItemName
}

func (m Material) Description() string {
	return m.ItemDescription
}

func (m Material) Type() string {
	return m.ItemType
}
