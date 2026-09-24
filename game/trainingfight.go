package game

import (
	"Projet-red-3/characters"
	"fmt"
)

func TrainingFight(p *characters.Character) {
	goblin := characters.InitGoblin()
	tour := 1

	fmt.Println("\n=== COMBAT D'ENTRAINEMENT ===")
	fmt.Println(p.Name, "affronte", goblin.Name)

	for p.HP > 0 && goblin.HP > 0 {
		fmt.Println("\nTour", tour)

		tour++

		if tour > 100 {
			break
		}
	}
}
