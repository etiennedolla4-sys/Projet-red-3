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

		enemy.HP = 0
	} else {
		fmt.Println("FACE !")
		fmt.Println("💥 Malchance !")

		player.HP /= 2

		if player.HP < 0 {
			player.HP = 0
		}

		fmt.Printf("Vous perdez la moitié de vos PV.\n")
		fmt.Printf("PV : %d / %d\n", player.HP, player.MaxHP)
	}
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
	fmt.Printf(
		"Vous perdez %d PV à cause de la puissance de l'attaque.\n",
		recoil,
	)
	fmt.Printf("PV : %d / %d\n", player.HP, player.MaxHP)
}
