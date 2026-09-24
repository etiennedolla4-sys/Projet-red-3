package combat

import (
	"Projet-red-3/characters"
	"Projet-red-3/enemy"
	"fmt"
)

func TrainingFight(p *characters.Character) {
	monster := enemy.NewSlime()
	oldHP := p.HP

	fmt.Println("\n=== COMBAT D'ENTRAÎNEMENT ===")

	for p.HP > 0 && monster.HP > 0 {
		DisplayCombat(p, &monster)

		if fuite := PlayerTurn(p, &monster); fuite {
			p.HP = oldHP
			return
		}

		if monster.HP <= 0 {
			break
		}

		EnemyTurn(p, &monster)
	}

	p.HP = oldHP

	if p.HP > 0 {
		fmt.Println("\nVous avez gagné !")
	} else {
		fmt.Println("\nVous avez perdu !")
	}
	fmt.Println("\nRetour au menu...")
}
