package combat

import (
	"Projet-red-3/characters"
	"Projet-red-3/enemy"
	"fmt"
	"math/rand"
)

func GobelinSkill(player *characters.Character, enemy *enemy.Monster) {
	fmt.Println("\n🪙 Vous lancez la pièce...")

	result := rand.Intn(2)

	if result == 0 {
		fmt.Println("PILE !")
		fmt.Println("💀 Coup de malade !")
		enemy.HP /= 2
	} else {
		fmt.Println("FACE !")
		fmt.Println("💥 Malchance !")
		player.HP /= 2
	}
}

func VampireSkill(player *characters.Character, enemy *enemy.Monster) {
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

	fmt.Println("\n🩸 Drain vampirique !")
	fmt.Printf("Vous infligez %d dégâts.\n", damage)
	fmt.Printf("Vous récupérez %d PV.\n", heal)
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
	fmt.Printf("La puissance de l'attaque vous fait perdre %d PV.\n", recoil)
	fmt.Printf("PV : %d / %d\n", player.HP, player.MaxHP)
}
