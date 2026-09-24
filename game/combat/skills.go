package combat

import (
	"Projet-red-3/characters"
	"fmt"
	"math/rand"
)

func GobelinSkill(player *characters.Character, enemy *Enemy) {
	fmt.Println("\n🪙 Vous lancez la pièce...")

	result := rand.Intn(2)

	if result == 0 {
		fmt.Println("PILE !")
		fmt.Println("💀 Coup mortel !")
		enemy.HP /= 2
	} else {
		fmt.Println("FACE !")
		fmt.Println("💥 Malchance !")

		player.HP /= 2

		if player.HP < 0 {
			player.HP = 0
		}

		fmt.Println("Vous perdez la moitié de vos PV.")
		fmt.Printf("PV : %d / %d\n", player.HP, player.MaxHP)
	}
}
func GobelinCoupVicieux(player *characters.Character, enemy *Enemy) {
	damage := player.BaseAttack + 10

	if rand.Intn(100) < 30 {
		damage *= 2
		fmt.Println("\n🍀 Coup critique chanceux !")
	}

	enemy.HP -= damage

	if enemy.HP < 0 {
		enemy.HP = 0
	}

	fmt.Printf("🔪 Vous infligez %d dégâts.\n", damage)
}

func GobelinDeTruque(player *characters.Character, enemy *Enemy) {
	damage := rand.Intn(41) + 10

	enemy.HP -= damage

	if enemy.HP < 0 {
		enemy.HP = 0
	}

	fmt.Println("\n🎲 Vous lancez votre dé truqué...")
	fmt.Printf("Le dé inflige %d dégâts !\n", damage)
}

func VampireSkill(player *characters.Character, enemy *Enemy) {
	damage := 25

	enemy.HP -= damage

	if enemy.HP < 0 {
		enemy.HP = 0
	}

	heal := damage / 2

	player.HP += heal

	if player.HP > player.MaxHP {
		player.HP = player.MaxHP
	}

	fmt.Println("\n🩸 DRAIN VAMPIRIQUE !")
	fmt.Printf("Vous infligez %d dégâts.\n", damage)
	fmt.Printf("Vous récupérez %d PV.\n", heal)
	fmt.Printf("PV : %d / %d\n", player.HP, player.MaxHP)
}

func VampireMorsure(player *characters.Character, enemy *Enemy) {
	damage := 35
	heal := 10

	enemy.HP -= damage
	player.HP += heal

	if enemy.HP < 0 {
		enemy.HP = 0
	}

	if player.HP > player.MaxHP {
		player.HP = player.MaxHP
	}

	fmt.Println("\n🧛 MORSURE !")
	fmt.Printf("Vous infligez %d dégâts.\n", damage)
	fmt.Printf("Vous récupérez %d PV.\n", heal)
}

func VampireSacrifice(player *characters.Character, enemy *Enemy) {
	cost := 15
	damage := 50

	if player.HP <= cost {
		fmt.Println("\nVous n'avez pas assez de PV pour utiliser cette compétence.")
		return
	}

	player.HP -= cost
	enemy.HP -= damage

	if enemy.HP < 0 {
		enemy.HP = 0
	}

	fmt.Println("\n🩸 SACRIFICE SANGUIN !")
	fmt.Printf("Vous perdez %d PV.\n", cost)
	fmt.Printf("Vous infligez %d dégâts.\n", damage)
}

func BerserkerSkill(player *characters.Character, enemy *Enemy) {
	fmt.Println("\n💢 RAGE DU BERSERKER !")

	damage := player.BaseAttack + 20

	enemy.HP -= damage

	if enemy.HP < 0 {
		enemy.HP = 0
	}

	recoil := 10

	player.HP -= recoil

	if player.HP < 0 {
		player.HP = 0
	}

	fmt.Printf("💥 Vous infligez %d dégâts !\n", damage)
	fmt.Printf("Vous perdez %d PV.\n", recoil)
}

func BerserkerExecution(player *characters.Character, enemy *Enemy) {
	damage := player.BaseAttack + 10

	if enemy.HP <= enemy.MaxHP/2 {
		damage *= 2
		fmt.Println("\n☠️ L'ennemi est affaibli !")
	}

	enemy.HP -= damage

	if enemy.HP < 0 {
		enemy.HP = 0
	}

	fmt.Println("🪓 EXÉCUTION !")
	fmt.Printf("Vous infligez %d dégâts.\n", damage)
}

func BerserkerFrenesie(player *characters.Character, enemy *Enemy) {
	fmt.Println("\n🔥 FRÉNÉSIE !")

	for i := 0; i < 2; i++ {
		damage := player.BaseAttack + rand.Intn(6)

		enemy.HP -= damage

		if enemy.HP < 0 {
			enemy.HP = 0
		}

		fmt.Printf("Coup %d : %d dégâts !\n", i+1, damage)

		if enemy.HP == 0 {
			break
		}
	}
}
