package sort

import (
	"fmt"

	"Projet-red-3/characters"
)

func SpellBook(p *characters.Character) {
	for _, skill := range p.Skills {
		if skill == "Boule de feu" {
			fmt.Println("Vous connaissez déjà le sort Boule de feu !")
			return
		}
	}

	p.Skills = append(p.Skills, "Boule de feu")
	fmt.Println("Vous avez appris le sort : Boule de feu !")
}
