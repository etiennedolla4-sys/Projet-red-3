package combat

import (
	"Projet-red-3/characters"
	"Projet-red-3/inventory"
	"Projet-red-3/utils"
	"fmt"
	"math/rand"
	"os"
)

func StartCombat(player *characters.Character, enemy *Enemy) bool {
	fmt.Println("\n==============================")
	fmt.Println("          COMBAT")
	fmt.Println("==============================")

	fmt.Printf("\n%s VS %s\n", player.Name, enemy.Name)

	playerInitiative := player.Initiative + rand.Intn(10)
	enemyInitiative := enemy.Initiative + rand.Intn(10)

	playerTurn := playerInitiative >= enemyInitiative
	defending := false

	if playerTurn {
		fmt.Println("\n⚡ Vous commencez le combat !")
	} else {
		fmt.Println("\n⚡ L'ennemi commence le combat !")
	}

	for player.HP > 0 && enemy.HP > 0 {
		DisplayCombat(player, enemy)

		if playerTurn {
			fuite, isDefending := PlayerTurn(player, enemy)
			defending = isDefending

			if fuite {
				fmt.Println("\n🏃 Vous avez fui le combat.")
				return false
			}
		} else {
			EnemyTurn(player, enemy, defending)
			defending = false
		}

		if player.HP <= 0 {
			inventory.TryRevive(player)
		}

		if player.HP <= 0 || enemy.HP <= 0 {
			break
		}

		playerTurn = !playerTurn
	}

	EndCombat(player, enemy)
	return enemy.HP <= 0
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

	fmt.Println()

	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║                COMBAT                ║")
	fmt.Println("╠══════════════════════════════════════╣")

	fmt.Printf("║  %-34s  ║\n", player.Name)
	fmt.Printf("║  PV : %3d / %-3d                      ║\n",
		player.HP,
		player.MaxHP,
	)
	fmt.Printf("║  %-34s  ║\n", HealthBar(player.HP, player.MaxHP))

	fmt.Println("║                                      ║")

	fmt.Printf("║  %-34s  ║\n", enemy.Name)
	fmt.Printf("║  PV : %3d / %-3d                      ║\n",
		enemy.HP,
		enemy.MaxHP,
	)
	fmt.Printf("║  %-34s  ║\n", HealthBar(enemy.HP, enemy.MaxHP))

	fmt.Println("╚══════════════════════════════════════╝")
}

func PlayerTurn(player *characters.Character, enemy *Enemy) (bool, bool) {
	for {
		var choice int

		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Compétence")
		fmt.Println("3. Défendre")
		fmt.Println("4. Fuir")
		fmt.Print("> ")

		fmt.Scan(&choice)
		utils.ClearTerminal()

		switch choice {
		case 1:
			Attack(player, enemy)
			return false, false
		case 2:
			UseSkill(player, enemy)
			return false, false
		case 3:
			Defend(player)
			return false, true
		case 4:
			fmt.Println("\n🏃 Vous prenez la fuite !")
			return true, false
		default:
			fmt.Println("\nChoix invalide.")
		}
	}
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
func UseMana(player *characters.Character, cost int) bool {
	if player.MP < cost {
		fmt.Printf("\nVous n'avez pas assez de mana. %d mana nécessaire.\n", cost)
		return false
	}

	player.MP -= cost
	return true
}
func UseSkill(player *characters.Character, enemy *Enemy) {
	var choice int

	switch player.Class {
	case "Gobelin":
		fmt.Println("\n=== COMPÉTENCES GOBELIN ===")
		fmt.Println("1. 🪙 Pile ou Face")
		fmt.Println("2. 🔪 Coup Vicieux")
		fmt.Println("3. 🎲 Dé Truqué")
		fmt.Print("> ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			GobelinSkill(player, enemy)
		case 2:
			GobelinCoupVicieux(player, enemy)
		case 3:
			GobelinDeTruque(player, enemy)
		default:
			fmt.Println("Choix invalide.")
		}

	case "Vampire":
		fmt.Println("\n=== COMPÉTENCES VAMPIRE ===")
		fmt.Println("1. 🩸 Drain Vampirique")
		fmt.Println("2. 🧛 Morsure")
		fmt.Println("3. 🩸 Sacrifice Sanguin")
		fmt.Print("> ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			VampireSkill(player, enemy)
		case 2:
			VampireMorsure(player, enemy)
		case 3:
			VampireSacrifice(player, enemy)
		default:
			fmt.Println("Choix invalide.")
		}

	case "Berserker":
		fmt.Println("\n=== COMPÉTENCES BERSERKER ===")
		fmt.Println("1. 💢 Rage")
		fmt.Println("2. 🪓 Exécution")
		fmt.Println("3. 🔥 Frénésie")
		fmt.Print("> ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			BerserkerSkill(player, enemy)
		case 2:
			BerserkerExecution(player, enemy)
		case 3:
			BerserkerFrenesie(player, enemy)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
func EnemyTurn(player *characters.Character, enemy *Enemy, defending bool) {
	damage := enemy.Attack - player.BaseDefense
	if defending {
		damage /= 2
	}

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
