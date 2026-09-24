package combat

import (
	"Projet-red-3/characters"
	"Projet-red-3/utils"
	"fmt"
)

func TrainingFight(player *characters.Character) {
	monster := NewSlime()

	fmt.Println("\n==============================")
	fmt.Println("     CAMP D'ENTRAÎNEMENT")
	fmt.Println("==============================")

	fmt.Printf("\nVous affrontez un %s !\n", monster.Name)

	playerTurn := player.Initiative >= monster.Initiative

	for player.HP > 0 && monster.HP > 0 {
		utils.ClearTerminal()

		DisplayCombat(player, &monster)

		if playerTurn {
			fuite := PlayerTurn(player, &monster)

			if fuite {
				utils.ClearTerminal()
				fmt.Println("\n🏃 Vous quittez l'entraînement.")
				return
			}
		} else {
			EnemyTurn(player, &monster)
		}

		if player.HP <= 0 || monster.HP <= 0 {
			break
		}

		playerTurn = !playerTurn
	}

	utils.ClearTerminal()

	if player.HP > 0 {
		fmt.Println("\n🏆 Vous avez remporté l'entraînement !")
	} else {
		fmt.Println("\n💀 Vous avez perdu l'entraînement...")
	}
}
