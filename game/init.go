package game

import (
	"Projet-red-3/characters"
	"Projet-red-3/item"
	"Projet-red-3/ui/text"
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Init() characters.Character {
	var name string
	var classe int
	var character characters.Character

	text.PrintSlow("Depuis des siècles, un immense gouffre appelé Veyr s'étend au cœur du royaume. Nul ne connaît sa profondeur, ni ce qui se cache au fond. Des aventuriers descendent chaque jour dans ses entrailles à la recherche de trésors, de reliques et de réponses. Mais plus ils descendent, plus le monde devient étrange… et moins ils sont nombreux à revenir.\nAujourd'hui, tu te tiens devant l'entrée de Veyr. Personne ne t'attend au fond. Personne ne sait ce que tu trouveras.\n \033[31mUne seule chose est certaine : pour découvrir la vérité, tu devras descendre.\033[0m\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Entrez votre nom : ")

		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		valide := name != ""

		for _, lettre := range name {
			if !unicode.IsLetter(lettre) {
				valide = false
				break
			}
		}

		if valide {
			break
		}
		fmt.Println("Le nom doit contenir uniquement des lettres et aucun espace.")
	}
	runes := []rune(strings.ToLower(name))
	runes[0] = unicode.ToUpper(runes[0])
	name = string(runes)

	for {
		fmt.Println("Choisissez votre classe ")
		fmt.Println("1. Gobelin")
		fmt.Println("2. Vampire")
		fmt.Println("3. Berserker")
		fmt.Print("> ")

		fmt.Scan(&classe)

		if classe >= 1 && classe <= 3 {
			break
		}

		fmt.Println("Choix invalide.")
	}
	switch classe {
	case 1:
		gobelin := characters.NewGobelin(name)
		character = gobelin.Character

	case 2:
		vampire := characters.NewVampire(name)
		character = vampire.Character

	case 3:
		berserker := characters.NewBerserker(name)
		character = berserker.Character
	}

	character.Inventory = append(
		character.Inventory,
		item.NewRevivalPotion(),
	)

	return character

}
