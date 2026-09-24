package combat

import (
	"Projet-red-3/characters"
	"Projet-red-3/enemy"
	"fmt"
	"math/rand"
	"os"
)

func StartCombat(player *characters.Character, enemy *enemy.Monster) {
	fmt.Println("\n==============================")
	fmt.Println("          COMBAT")
	fmt.Println("==============================")

	fmt.Printf("\n%s VS %s\n", player.Name, enemy.Name)

	// Détermine qui commence
	playerInitiative := player.Initiative + rand.Intn(10)
	enemyInitiative := enemy.Initiative + rand.Intn(10)

	playerTurn := playerInitiative >= enemyInitiative

	for player.HP > 0 && enemy.HP > 0 {

		DisplayCombat(player, enemy)

		if playerTurn {
			fuite := PlayerTurn(player, enemy)

			if fuite {
				return
			}
		} else {
			EnemyTurn(player, enemy)
		}

		// Vérifie si quelqu'un est mort
		if player.HP <= 0 || enemy.HP <= 0 {
			break
		}

		// Change de tour
		playerTurn = !playerTurn
	}

	EndCombat(player, enemy)
}
func HealthBar(hp int, maxHP int) string {
	const size = 20

	ratio := float64(hp) / float64(maxHP)
	filled := int(ratio * size)

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
func DisplayCombat(player *characters.Character, enemy *enemy.Monster) {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║                COMBAT                ║")
	fmt.Println("╠══════════════════════════════════════╣")

	fmt.Printf("║  %-34s║\n", player.Name)
	fmt.Printf("║  PV : %3d / %-3d                    ║\n", player.HP, player.MaxHP)
	fmt.Printf("║  %s ║\n", HealthBar(player.HP, player.MaxHP))

	fmt.Println("║                                      ║")

	fmt.Printf("║  %-34s║\n", enemy.Name)
	fmt.Printf("║  PV : %3d / %-3d                    ║\n", enemy.HP, enemy.MaxHP)
	fmt.Printf("║  %s ║\n", HealthBar(enemy.HP, enemy.MaxHP))

	fmt.Println("╚══════════════════════════════════════╝")
}

func PlayerTurn(player *characters.Character, enemy *enemy.Monster) bool {
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
		fmt.Println("\nVous prenez la fuite !")
		return true

	default:
		fmt.Println("\nChoix invalide.")
	}
	return false
}

func Attack(player *characters.Character, enemy *enemy.Monster) {
	damage := player.BaseAttack - enemy.Defense

	// Ajoute un petit hasard aux dégâts
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

	// Pour l'instant, on ajoute simplement un bonus temporaire
	// qu'on gérera mieux ensuite.
}

func UseSkill(player *characters.Character, enemy *enemy.Monster) {
	switch player.Class {

	case "Gobelin":
		GobelinSkill(player, enemy)

	case "Vampire":
		VampireSkill(player, enemy)

	default:
		fmt.Println("\nCette classe n'a pas encore de compétence.")
	}
}

func EnemyTurn(player *characters.Character, enemy *enemy.Monster) {
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

func EndCombat(player *characters.Character, enemy *enemy.Monster) {
	fmt.Println("\n==============================")

	if player.HP <= 0 {
		fmt.Println("💀 Vous avez été vaincu...")
		os.Exit(0)
	} else {
		fmt.Printf("🏆 Vous avez vaincu %s !\n", enemy.Name)

		player.Gold += 20
		fmt.Println("Vous gagnez 20 gold !")

		GainXP(player, enemy.XP)
	}

	fmt.Println("==============================")
}
