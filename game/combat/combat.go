package game

import (
	"Projet-red-3/characters"
	"fmt"
)

func Combat(p1, p2 characters.Character) {
	fmt.Println("Combat entre", p1.Name, "et", p2.Name)
	for p1.HP > 0 && p2.HP > 0 {
		// Tour du joueur 1
		fmt.Println(p1.Name, "attaque", p2.Name)
		damage := p1.BaseAttack - p2.BaseDefense
		if damage < 0 {
			damage = 0
		}
		p2.HP -= damage
		fmt.Println(p2.Name, "subit", damage, "points de dégâts. PV restants :", p2.HP)
	}
}
