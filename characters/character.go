package characters

import "Projet-red-3/item"

type Equipment struct {
	Head  string
	Chest string
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
	Inventory    []item.Item
	MaxInventory int
	HasBag       bool
	MaxHP        int
	MaxMP        int
	Crit         int
	Lvl          int
	XP           int
	Gold         int
	Skills       []string
	Equipment    Equipment
	Revived      bool
}
