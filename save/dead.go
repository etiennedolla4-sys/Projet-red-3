package save

import (
	"fmt"

	"Projet-red-3/characters"
)

func Dead(p *characters.Character) {
	if p.HP == 0 {
		p.HP = p.MaxHP / 2

		fmt.Println("Vous êtes mort !")
		fmt.Println("Vous êtes ressuscité !")
		fmt.Printf("PV : %d / %d\n", p.HP, p.MaxHP)
	}
}
