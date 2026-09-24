package characters

import "Projet-red-3/item"

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
	MaxHP        int
	MaxMP        int
	Crit         int
	Lvl          int
	XP           int
	Gold         int
	Skills       []string
}
