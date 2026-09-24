package game

import (
	"Projet-red-3/characters"
	"fmt"
)

func GoblinPattern(goblin characters.Monster, player *characters.Character, tour int) {
	damage := goblin.BaseAttack

	if tour%3 == 0 {
		damage *= 2
	}

	player.HP -= damage

	if player.HP < 0 {
		player.HP = 0
	}

	fmt.Printf("%s inflige à %s %d dégâts\n",
		goblin.Name,
		player.Name,
		damage,
	)

	fmt.Printf("PV : %d / %d\n", player.HP, player.MaxHP)
}
