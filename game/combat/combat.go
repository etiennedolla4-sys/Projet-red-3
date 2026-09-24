package combat

import (
	"Projet-red-3/characters"
	"fmt"
	"math/rand"
	"os"
)

func StartCombat(player *characters.Character, enemy *Enemy) {
	fmt.Println("\n==============================")
	fmt.Println("          COMBAT")
	fmt.Println("==============================")

	fmt.Printf("\n%s VS %s\n", player.Name, enemy.Name)

	playerInitiative := player.Initiative + rand.Intn(10)
	enemyInitiative := enemy.Initiative + rand.Intn(10)

	playerTurn := playerInitiative >= enemyInitiative

	if playerTurn {
		fmt.Println("\n⚡ Vous commencez le combat !")
	} else {
		fmt.Println("\n⚡ L'ennemi commence le combat !")
	}

	for player.HP > 0 && enemy.HP > 0 {
		DisplayCombat(player, enemy)

		if playerTurn {
			fuite := PlayerTurn(player, enemy)

			if fuite {
				fmt.Println("\n🏃 Vous avez fui le combat.")
				return
			}
		} else {
			EnemyTurn(player, enemy)
		}

		if player.HP <= 0 || enemy.HP <= 0 {
			break
		}

		playerTurn = !playerTurn
	}

	EndCombat(player, enemy)
}

func HealthBar(hp int, maxHP int) string {
	const size = 20

	if maxHP <= 0 {
		return "[░░░░░░░░░░░░░░░░░░░░]"
	}

	ratio := float64(hp) / float64(maxHP)
	filled := int(ratio * size)

	if filled < 0 {
		filled = 0
	}

	if filled > size {
		filled = size
	}

	bar := ""

	for i := 0; i < size; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	return "[" + bar + "]"
}

func DisplayCombat(player *characters.Character, enemy *Enemy) {
	fmt.Println()

	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║                COMBAT                ║")
	fmt.Println("╠══════════════════════════════════════╣")

	fmt.Printf("║  %-34s║\n", player.Name)
	fmt.Printf("║  PV : %3d / %-3d                    ║\n",
		player.HP,
		player.MaxHP,
	)
	fmt.Printf("║  %s ║\n", HealthBar(player.HP, player.MaxHP))

	fmt.Println("║                                      ║")

	fmt.Printf("║  %-34s║\n", enemy.Name)
	fmt.Printf("║  PV : %3d / %-3d                    ║\n",
		enemy.HP,
		enemy.MaxHP,
	)
	fmt.Printf("║  %s ║\n", HealthBar(enemy.HP, enemy.MaxHP))

	fmt.Println("╚══════════════════════════════════════╝")
}

func PlayerTurn(player *characters.Character, enemy *Enemy) bool {
	var choice int

	fmt.Println("\nQue voulez-vous faire ?")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Compétence")
	fmt.Println("3. Défendre")
	fmt.Println("4. Fuir")
	fmt.Print("> ")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		Attack(player, enemy)

	case 2:
		UseSkill(player, enemy)

	case 3:
		Defend(player)

	case 4:
		fmt.Println("\n🏃 Vous prenez la fuite !")
		return true

	default:
		fmt.Println("\nChoix invalide.")
	}

	return false
}

func Attack(player *characters.Character, enemy *Enemy) {
	damage := player.BaseAttack - enemy.Defense
	damage += rand.Intn(5)

	if damage < 1 {
		damage = 1
	}

	enemy.HP -= damage

	if enemy.HP < 0 {
		enemy.HP = 0
	}

	fmt.Printf("\n⚔️ Vous infligez %d dégâts !\n", damage)
}

func Defend(player *characters.Character) {
	fmt.Println("\n🛡 Vous vous mettez en défense !")
	fmt.Println("Votre défense est renforcée pour ce tour.")
}

func UseSkill(player *characters.Character, enemy *Enemy) {
	switch player.Class {
	case "Gobelin":
		GobelinSkill(player, enemy)

	case "Vampire":
		VampireSkill(player, enemy)

	case "Berserker":
		BerserkerSkill(player, enemy)

	default:
		fmt.Println("\nCette classe n'a pas encore de compétence.")
	}
}

func EnemyTurn(player *characters.Character, enemy *Enemy) {
	damage := enemy.Attack - player.BaseDefense

	if damage < 1 {
		damage = 1
	}

	player.HP -= damage

	if player.HP < 0 {
		player.HP = 0
	}

	fmt.Printf("\n👹 %s attaque !\n", enemy.Name)
	fmt.Printf("Vous perdez %d PV.\n", damage)
}

func EndCombat(player *characters.Character, enemy *Enemy) {
	fmt.Println("\n==============================")

	if player.HP <= 0 {
		fmt.Println("💀 Vous avez été vaincu...")
		os.Exit(0)
	} else if enemy.HP <= 0 {
		fmt.Printf("🏆 Vous avez vaincu %s !\n", enemy.Name)

		player.Gold += 20

		fmt.Println("Vous gagnez 20 Gold !")
		fmt.Printf("Gold actuel : %d\n", player.Gold)

		GainXP(player, enemy.XP)
	}

	fmt.Println("==============================")
}
