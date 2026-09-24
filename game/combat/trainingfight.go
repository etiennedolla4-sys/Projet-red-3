package combat

import (
	"Projet-red-3/characters"
	"Projet-red-3/monsters"
	"fmt"
)

func TrainingFight(p *characters.Character) {
	monster := monsters.NewSlime()

	turn := 1

	fmt.Println("\n=== COMBAT D'ENTRAÎNEMENT ===")

	for p.HP > 0 && monster.HP > 0 {
		fmt.Printf("\n=== Tour %d ===\n", turn)

		PlayerTurn(p, &monster)

		if monster.HP <= 0 {
			break
		}

		MonsterTurn(p, &monster)

		turn++
	}

	if p.HP > 0 {
		fmt.Println("\nVous avez gagné !")
	} else {
		fmt.Println("\nVous avez perdu !")
	}
}
