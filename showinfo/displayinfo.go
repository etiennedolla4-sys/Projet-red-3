package showinfo

import (
	"fmt"

	"Projet-red-3/characters"
)

func DisplayInfo(p characters.Character) {
	fmt.Println("\n=== INFORMATIONS DU PERSONNAGE ===")
	fmt.Println("Nom :", p.Name)
	fmt.Println("Description :", p.Description)
	fmt.Println("Niveau :", p.Lvl)
	fmt.Printf("PV : %d / %d\n", p.HP, p.MaxHP)
	fmt.Printf("Mana : %d / %d\n", p.MP, p.MaxMP)
	fmt.Println("Attaque :", p.BaseAttack)
	fmt.Println("Défense :", p.BaseDefense)
	fmt.Println("Initiative :", p.Initiative)
	fmt.Println("Critique :", p.Crit)
	fmt.Printf("Gold : %d\n", p.Gold)
}
