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

	p.Name = name
	p.Lvl = 1
	p.HP = 50
	p.MP = 0
	p.Inventory = []string{}
	p.MaxInventory = 10
	p.Skills = []string{"Coup de poing"}

	switch classe {
	case 1:
		p.Class = "Gobelin"
		p.MaxHP = 100

	case 2:
		p.Class = "Vampire"
		p.MaxHP = 80

	}

	p.HP = p.MaxHP / 2

	return p
}
