package game

import (
	"Projet-red-3/characters"
	"Projet-red-3/game/combat"
	"fmt"
	"math/rand"
)

func RandomEvent(character *characters.Character, depth int) {
	event := rand.Intn(4)

	switch event {
	case 0:
		FindResource(character)

	case 1:
		FindChest(character)

	case 2:
		StartEncounter(character, depth)

	case 3:
		FindRest(character)
	}
}

func StartEncounter(character *characters.Character, depth int) {
	enemy := combat.GetRandomEnemy(depth)

	fmt.Println()
	fmt.Println("========== RENCONTRE ==========")
	fmt.Printf("Un %s apparaît dans les profondeurs !\n", enemy.Name)

	combat.StartCombat(character, &enemy)
}

func FindResource(character *characters.Character) {
	gold := rand.Intn(20) + 5
	character.Gold += gold

	fmt.Println()
	fmt.Println("========== RESSOURCE ==========")
	fmt.Println("Vous trouvez une petite ressource.")
	fmt.Printf("Vous gagnez %d pièces d'or.\n", gold)
}

func FindChest(character *characters.Character) {
	gold := rand.Intn(50) + 10
	character.Gold += gold

	fmt.Println()
	fmt.Println("========== COFFRE ==========")
	fmt.Println("Vous découvrez un ancien coffre.")
	fmt.Printf("Vous trouvez %d pièces d'or.\n", gold)
}

func FindRest(character *characters.Character) {
	heal := 20

	character.HP += heal

	if character.HP > character.MaxHP {
		character.HP = character.MaxHP
	}

	fmt.Println()
	fmt.Println("========== REPOS ==========")
	fmt.Println("Vous trouvez un endroit relativement sûr.")
	fmt.Printf("Vous récupérez %d PV.\n", heal)
	fmt.Printf("PV : %d / %d\n", character.HP, character.MaxHP)
}
