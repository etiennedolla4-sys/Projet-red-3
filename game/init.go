package game

import (
	"Projet-red-3/characters"
	"Projet-red-3/printslow"
	"fmt"
)

func Init() characters.Character {
	var name string
	var classe int
	printslow.PrintSlow("Depuis des siècles, un immense gouffre appelé Veyr s'étend au cœur du royaume. Nul ne connaît sa profondeur, ni ce qui se cache au fond. Des aventuriers descendent chaque jour dans ses entrailles à la recherche de trésors, de reliques et de réponses. Mais plus ils descendent, plus le monde devient étrange… et moins ils sont nombreux à revenir.\nAujourd'hui, tu te tiens devant l'entrée de Veyr. Personne ne t'attend au fond. Personne ne sait ce que tu trouveras. \033Une seule chose est certaine : pour découvrir la vérité, tu devras descendre.\033[0m")
	fmt.Print("Entrez votre pseudo : ")
	fmt.Scan(&name)

	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Gobelin")
	fmt.Println("2. Vampire")
	fmt.Scan(&classe)

	var character characters.Character

	switch classe {
	case 1:
		gobelin := characters.NewGobelin(name)
		character = gobelin.Character

	case 2:
		vampire := characters.NewVampire(name)
		character = vampire.Character
	}

	character.Name = name

	return character
}
