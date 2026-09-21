package characters

type Character struct {
	Name         string
	Description  string
	HP           int
	Mana         int
	BaseAttack   int
	BaseDefense  int
	Initiative   int
	Inventory    []string
	MaxInventory int
	MaxHP        int
	MaxMP        int
	Crit         int
	Lvl          int
}
