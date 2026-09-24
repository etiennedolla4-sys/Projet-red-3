package game

import (
	"Projet-red-3/characters"
	"fmt"
	"math/rand"
)

func RandomEvent(character *characters.Character) {
	event := rand.Intn(5)

	switch event {
	case 0:
		FindResource(character)

	case 1:
		FindChest(character)

	case 2:
		EncounterEnemy(character)

	case 3:
		FindRest(character)

	case 4:
		NothingHappens()
	}
}

func FindResource(character *characters.Character) {
	gold := rand.Intn(20) + 5
	character.Gold += gold

	fmt.Println("\n========== RESSOURCE ==========")
	fmt.Println("Vous trouvez une petite ressource.")
	fmt.Printf("Vous gagnez %d pièces d'or.\n", gold)
}

func FindChest(character *characters.Character) {
	gold := rand.Intn(50) + 10
	character.Gold += gold

	fmt.Println("\n========== COFFRE ==========")
	fmt.Println("Vous découvrez un ancien coffre.")
	fmt.Printf("Vous trouvez %d pièces d'or.\n", gold)
}

func EncounterEnemy(character *characters.Character) {
	damage := rand.Intn(15) + 5

	fmt.Println("\n========== RENCONTRE ==========")
	fmt.Println("Une créature des profondeurs apparaît !")
	fmt.Println("Vous parvenez à vous défendre.")

	character.HP -= damage

	if character.HP < 0 {
		character.HP = 0
	}

	fmt.Printf("Vous perdez %d PV.\n", damage)
	fmt.Printf("PV : %d / %d\n", character.HP, character.MaxHP)
}

func FindRest(character *characters.Character) {
	heal := 20

	character.HP += heal

	if character.HP > character.MaxHP {
		character.HP = character.MaxHP
	}

	fmt.Println("\n========== REPOS ==========")
	fmt.Println("Vous trouvez un endroit relativement sûr.")
	fmt.Printf("Vous récupérez %d PV.\n", heal)
	fmt.Printf("PV : %d / %d\n", character.HP, character.MaxHP)
}

func NothingHappens() {
	fmt.Println("\nVous avancez dans le silence...")
	fmt.Println("Rien ne semble vous attendre ici.")
}
