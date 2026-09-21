package game

import (
	"Projet-red-3/characters"
	"fmt"
)

func Init() characters.Character {
	var name string
	var classe int

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
