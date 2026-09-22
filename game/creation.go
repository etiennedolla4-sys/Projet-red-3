package game

import (
	"fmt"
	"strings"
	"unicode"

	"Projet-red-3/characters"
)

func CharCreation() characters.Character {
	var name string
	var classe int

	for {
		fmt.Print("Entrez votre nom : ")
		fmt.Scan(&name)

		valide := true

		for _, lettre := range name {
			if !unicode.IsLetter(lettre) {
				valide = false
			}
		}

		if valide {
			break
		}

		fmt.Println("Le nom doit contenir uniquement des lettres.")
	}

	name = strings.ToUpper(name[:1]) + strings.ToLower(name[1:])

	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Gobelin")
	fmt.Println("2. Vampire")
	fmt.Scan(&classe)

	var p characters.Character

	switch classe {
	case 1:
		gobelin := characters.NewGobelin(name)
		p = gobelin.Character

	case 2:
		vampire := characters.NewVampire(name)
		p = vampire.Character

	default:
		fmt.Println("Choix invalide")
		return CharCreation()
	}

	return p
}
