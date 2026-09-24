package game

import (
	"Projet-red-3/characters"
	"fmt"
)

func ApplyReturnEffect(character *characters.Character, depth int) {
	if depth <= 1 {
		fmt.Println("La remontée ne provoque aucun effet.")
		return
	}

	damage := depth * 5

	character.HP -= damage

	if character.HP < 0 {
		character.HP = 0
	}

	fmt.Printf(
		"La pression des profondeurs vous affecte : -%d PV.\n",
		damage,
	)
}
