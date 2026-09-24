package combat

import (
	"Projet-red-3/characters"
	"fmt"
)

const xpToLevelUp = 100

func GainXP(player *characters.Character, amount int) {
	fmt.Printf("\n✨ Vous gagnez %d XP !\n", amount)

	player.XP += amount

	for player.XP >= xpToLevelUp {
		player.XP -= xpToLevelUp
		player.Lvl++
		LevelUp(player)
	}
}

func LevelUp(player *characters.Character) {
	fmt.Printf("\n🌟 Niveau supérieur ! Vous êtes maintenant niveau %d !\n", player.Lvl)

	player.MaxHP += 10
	player.HP = player.MaxHP
	player.BaseAttack += 2
	player.BaseDefense += 1
}
