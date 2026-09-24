package combat

import (
	"Projet-red-3/characters"
	"fmt"
	"math/rand"
)

type Enemy struct {
	Name       string
	HP         int
	MaxHP      int
	Attack     int
	Defense    int
	Initiative int
}

func NewSlime() Enemy {
	return Enemy{
		Name:       "Slime",
		HP:         80,
		MaxHP:      80,
		Attack:     15,
		Defense:    3,
		Initiative: 3,
	}
}

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
			if PlayerTurn(player, enemy) {
				fmt.Println("\nVous avez fui le combat.")
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

func DisplayCombat(player *characters.Character, enemy *Enemy) {
	fmt.Println("\n------------------------------")

	fmt.Printf(
		"%s : %d/%d PV\n",
		player.Name,
		player.HP,
		player.MaxHP,
	)

	fmt.Printf(
		"%s : %d/%d PV\n",
		enemy.Name,
		enemy.HP,
		enemy.MaxHP,
	)

	fmt.Println("------------------------------")
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
	} else if enemy.HP <= 0 {
		fmt.Printf("🏆 Vous avez vaincu %s !\n", enemy.Name)

		player.Gold += 20

		fmt.Println("Vous gagnez 20 Gold !")
		fmt.Printf("Gold actuel : %d\n", player.Gold)
	}

	fmt.Println("==============================")
}
