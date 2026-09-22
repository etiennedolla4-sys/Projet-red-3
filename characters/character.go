package characters

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Character struct {
	Name         string
	Class        string
	Description  string
	HP           int
	MP           int
	BaseAttack   int
	BaseDefense  int
	Initiative   int
	Inventory    []string
	MaxInventory int
	MaxHP        int
	MaxMP        int
	Crit         int
	Lvl          int
	Gold         int
	Skills       []string
	Equipment    Equipment
}
