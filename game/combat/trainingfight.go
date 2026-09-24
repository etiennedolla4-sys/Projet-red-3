package combat

import (
	"Projet-red-3/characters"
	"fmt"
)

func TrainingFight(player *characters.Character) {
	monster := NewSlime()
	startingHP := player.HP

	fmt.Println("\n==============================")
	fmt.Println("     CAMP D'ENTRAÎNEMENT")
	fmt.Println("==============================")

	fmt.Printf("\nVous affrontez un %s !\n", monster.Name)

	playerTurn := player.Initiative >= monster.Initiative
	defending := false

	for player.HP > 0 && monster.HP > 0 {
		DisplayCombat(player, &monster)

		if playerTurn {
			fuite, isDefending := PlayerTurn(player, &monster)
			defending = isDefending

			if fuite {
				player.HP = startingHP
				fmt.Println("\n🏃 Vous quittez l'entraînement.")
				fmt.Printf("Vos PV reviennent à %d / %d.\n", player.HP, player.MaxHP)
				return
			}
		} else {
			EnemyTurn(player, &monster, defending)
			defending = false
		}

		if player.HP <= 0 || monster.HP <= 0 {
			break
		}

		playerTurn = !playerTurn
	}

	if player.HP > 0 {
		fmt.Println("\n🏆 Vous avez remporté l'entraînement !")
	} else {
		fmt.Println("\n💀 Vous avez perdu l'entraînement...")
	}

	player.HP = startingHP
	fmt.Printf("Vos PV reviennent à %d / %d.\n", player.HP, player.MaxHP)
}
