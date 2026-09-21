package main

import "fmt"

func main() {
	var choix int

	for {
		fmt.Println("=== MENU ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Quitter")
		fmt.Print("Votre choix : ")

		fmt.Scan(&choix)

		switch choix {
		case 1:
			fmt.Println("Informations du personnage")
		case 2:
			fmt.Println("Inventaire")
		case 3:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide")
		}

		fmt.Println()
	}
}
