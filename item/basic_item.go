package item

type BasicItem struct {
	ItemName        string
	ItemDescription string
	ItemType        string
	AttackBonus     int
	DefenseBonus    int
	InitiativeBonus int
	HPBonus         int
}

func (i BasicItem) Name() string {
	return i.ItemName
}

func (i BasicItem) Description() string {
	return i.ItemDescription
}

func (i BasicItem) Type() string {
	return i.ItemType
}
