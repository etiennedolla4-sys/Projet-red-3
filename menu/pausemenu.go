package menu

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Println("Appuyez sur P pour ouvrir le menu.")

	for {
		var touche [1]byte
		os.Stdin.Read(touche[:])

		if touche[0] == 'p' || touche[0] == 'P' {
			menu()
			fmt.Println("\nAppuyez sur P pour ouvrir le menu.")
		}
	}
}

func menu() {
	var choix [1]byte

	fmt.Println("\n=== MENU  PAUSE ===")
	fmt.Println("1. Afficher les informations du personnage")
	fmt.Println("2. Accéder au contenu de l'inventaire")
	fmt.Println("3. Quitter")
	fmt.Print("Votre choix : ")

	os.Stdin.Read(choix[:])

	switch choix[0] {
	case '1':
		fmt.Println("Informations du personnage")
	case '2':
		fmt.Println("Inventaire")
	case '3':
		fmt.Println("Au revoir !")
		os.Exit(0)
	default:
		fmt.Println("Choix invalide")
	}
}
