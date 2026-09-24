package menu

import (
	"Projet-red-3/characters"
	"Projet-red-3/game/exploration"
	"Projet-red-3/npc/merchant"
	"Projet-red-3/ui/showinfo"
	"Projet-red-3/utils"
	"fmt"
	"os"

	"golang.org/x/term"
)

func DetecP(p characters.Character) {
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Println("Appuyez sur P pour ouvrir le menu.")

	for {
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Erreur terminal :", err)
			return
		}

		var touche [1]byte
		os.Stdin.Read(touche[:])

		term.Restore(int(os.Stdin.Fd()), oldState)

		if touche[0] == 'p' || touche[0] == 'P' {
			utils.ClearTerminal()
			menu(&p)

			fmt.Println()
			fmt.Println("Appuyez sur P pour ouvrir le menu.")
		}
	}
}

func menu(p *characters.Character) {
	var choix [1]byte

	utils.ClearTerminal()

	fmt.Println("========== MENU ==========")
	fmt.Println("1. Afficher les informations")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Marchand")
	fmt.Println("4. Exploration")
	fmt.Println("5. Quitter le jeu")
	fmt.Print("Votre choix : ")

	os.Stdin.Read(choix[:])

	utils.ClearTerminal()

	switch choix[0] {
	case '1':
		showinfo.DisplayInfo(*p)

	case '2':
		showinfo.DisplayInventory(p)

	case '3':
		npc.Merchant(p)

	case '4':
		game.Start(p)

	case '5':
		fmt.Println("Au revoir !")
		os.Exit(0)

	default:
		fmt.Println("Choix invalide.")
	}
}
