package menu

import (
	"Projet-red-3/characters"
	"Projet-red-3/personnage"
	"Projet-red-3/showinfo"
	"fmt"
	"os"

	"golang.org/x/term"
)

func DetecP(p characters.Character) {
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Println("Appuyez sur P pour ouvrir le menu.")

	for {
		var touche [1]byte
		os.Stdin.Read(touche[:])

		if touche[0] == 'p' || touche[0] == 'P' {
			menu(p)
			fmt.Println("\nAppuyez sur P pour ouvrir le menu.")
		}
	}
}

func menu(p characters.Character) {
	var choix [1]byte

	fmt.Println("1. Afficher les informations du personnage")
	fmt.Println("2. Accéder au contenu de l'inventaire")
	fmt.Println("3. Marchand")
	fmt.Println("4. Quitter")

	os.Stdin.Read(choix[:])

	switch choix[0] {
	case '1':
		showinfo.DisplayInfo(p)

	case '2':
		showinfo.DisplayInventory(&p)

	case '3':
		personnage.Merchant(&p)

	case '4':
		fmt.Println("Au revoir !")
		os.Exit(0)

	default:
		fmt.Println("Choix invalide")
	}
}
