package game

import (
	"Projet-red-3/characters"
	"Projet-red-3/ui/text"
	"fmt"
	"strings"
	"unicode"
)

func Init() characters.Character {
	var name string
	var classe int
	var character characters.Character

	text.PrintSlow("Depuis des siècles, un immense gouffre appelé Veyr s'étend au cœur du royaume. Nul ne connaît sa profondeur, ni ce qui se cache au fond. Des aventuriers descendent chaque jour dans ses entrailles à la recherche de trésors, de reliques et de réponses. Mais plus ils descendent, plus le monde devient étrange… et moins ils sont nombreux à revenir.\nAujourd'hui, tu te tiens devant l'entrée de Veyr. Personne ne t'attend au fond. Personne ne sait ce que tu trouveras.\n \033[31mUne seule chose est certaine : pour découvrir la vérité, tu devras descendre.\033[0m\n")

	for {
		fmt.Print("Entrez votre nom : ")
		fmt.Scan(&name)

		valide := true

		if name == "" {
			valide = false
		}

		for _, lettre := range name {
			if !unicode.IsLetter(lettre) {
				valide = false
				break
			}
		}

		if valide {
			break
		}

		fmt.Println("Le nom doit contenir uniquement des lettres.")
	}

	name = strings.ToUpper(name[:1]) + strings.ToLower(name[1:])

	fmt.Println("Choisissez votre classe ")
	fmt.Println("1. Gobelin")
	fmt.Println("2. Vampire")
	fmt.Scan(&classe)

	if classe < 3 && classe > 0 {
		switch classe {
		case 1:
			gobelin := characters.NewGobelin(name)
			character = gobelin.Character

		case 2:
			vampire := characters.NewVampire(name)
			character = vampire.Character
		}
	} else {
		fmt.Println("Choix invalide")
	}

	return character
}
