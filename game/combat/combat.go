package combat

import (
	"Projet-red-3/characters"
	"Projet-red-3/inventory"
	"Projet-red-3/ui/text"
	"Projet-red-3/utils"
	"fmt"
	"math/rand"
	"os"
	"time"
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
		utils.ClearTerminal()
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
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║                COMBAT                ║")
	fmt.Println("╠══════════════════════════════════════╣")

	fmt.Printf("║  %-36s║\n", player.Name)
	fmt.Printf("║  %-36s║\n",
		fmt.Sprintf("PV : %d / %d", player.HP, player.MaxHP),
	)
	fmt.Printf("║  %-36s║\n",
		fmt.Sprintf("Mana : %d / %d", player.MP, player.MaxMP),
	)
	fmt.Printf("║  %-36s║\n", HealthBar(player.HP, player.MaxHP))

	fmt.Println("║                                      ║")

	fmt.Printf("║  %-36s║\n", enemy.Name)
	fmt.Printf("║  %-36s║\n",
		fmt.Sprintf("PV : %d / %d", enemy.HP, enemy.MaxHP),
	)
	fmt.Printf("║  %-36s║\n", HealthBar(enemy.HP, enemy.MaxHP))

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
			time.Sleep(800 * time.Millisecond)
			return false, false
		case 2:
			if !UseSkill(player, enemy) {
				continue
			}
			time.Sleep(800 * time.Millisecond)
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
func UseSkill(player *characters.Character, enemy *Enemy) bool {
	var choice int

	switch player.Class {
	case "Gobelin":
		fmt.Println("\n=== COMPÉTENCES GOBELIN ===")
		fmt.Println("1. 🪙 Pile ou Face - 25 Mana")
		fmt.Println("2. 🔪 Coup Vicieux - 30 Mana")
		fmt.Println("3. 🎲 Dé Truqué - 30 Mana")
		fmt.Println("0. Retour")
		fmt.Printf("Mana : %d / %d\n", player.MP, player.MaxMP)
		fmt.Print("> ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			return GobelinSkill(player, enemy)
		case 2:
			return GobelinCoupVicieux(player, enemy)
		case 3:
			return GobelinDeTruque(player, enemy)
		case 0:
			return false
		default:
			fmt.Println("Choix invalide.")
			return false
		}

	case "Vampire":
		fmt.Println("\n=== COMPÉTENCES VAMPIRE ===")
		fmt.Println("1. 🩸 Drain Vampirique - 40 Mana")
		fmt.Println("2. 🧛 Morsure - 30 Mana")
		fmt.Println("3. 🩸 Sacrifice Sanguin - 35 Mana")
		fmt.Println("0. Retour")
		fmt.Printf("Mana : %d / %d\n", player.MP, player.MaxMP)
		fmt.Print("> ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			return VampireSkill(player, enemy)
		case 2:
			return VampireMorsure(player, enemy)
		case 3:
			return VampireSacrifice(player, enemy)
		case 0:
			return false
		default:
			fmt.Println("Choix invalide.")
			return false
		}

	case "Berserker":
		fmt.Println("\n=== COMPÉTENCES BERSERKER ===")
		fmt.Println("1. 💢 Rage - 25 Mana")
		fmt.Println("2. 🪓 Exécution - 30 Mana")
		fmt.Println("3. 🔥 Frénésie - 35 Mana")
		fmt.Println("0. Retour")
		fmt.Printf("Mana : %d / %d\n", player.MP, player.MaxMP)
		fmt.Print("> ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			return BerserkerSkill(player, enemy)
		case 2:
			return BerserkerExecution(player, enemy)
		case 3:
			return BerserkerFrenesie(player, enemy)
		case 0:
			return false
		default:
			fmt.Println("Choix invalide.")
			return false
		}
	}

	return false
}
func ShowEnemyLore(enemyName string) {
	switch enemyName {
	case "Slime":
		text.PrintSlow("Le slime tremble encore quelques secondes avant de se dissoudre.\n")

	case "Petit Gobelin":
		text.PrintSlow("Le petit gobelin serre une vieille pièce contre lui avant de tomber.\n")

	case "Gobelin des ruines":
		text.PrintSlow("Son équipement porte les mêmes symboles que ceux gravés sur les ruines.\n")

	case "Golem des cavernes":
		text.PrintSlow("Les pierres qui composent son corps cessent lentement de vibrer.\n")
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
	time.Sleep(800 * time.Millisecond)

}

func EndCombat(player *characters.Character, enemy *Enemy) {
	fmt.Println("\n==============================")

	if player.HP <= 0 {
		fmt.Println("💀 Vous avez été vaincu...")
		os.Exit(0)
	} else if enemy.HP <= 0 {
		ShowEnemyLore(enemy.Name)
		fmt.Printf("🏆 Vous avez vaincu %s !\n", enemy.Name)

		player.Gold += 20

		fmt.Println("Vous gagnez 20 Gold !")
		fmt.Printf("Gold actuel : %d\n", player.Gold)

		GainXP(player, enemy.XP)
	}

	fmt.Println("==============================")
}
