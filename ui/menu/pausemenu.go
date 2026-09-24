package menu

import (
	"Projet-red-3/characters"
	"Projet-red-3/game/combat"
	npc "Projet-red-3/npc/merchant"
	"Projet-red-3/ui/showinfo"
	"Projet-red-3/ui/text"
	"fmt"
	"os"

	"golang.org/x/term"
)

func DetecP(p *characters.Character) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Erreur terminal :", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Println("Appuyez sur P pour ouvrir le menu.")

	for {
		b, err := text.Stdin.ReadByte()
		if err != nil {
			return
		}

		if b == 'p' || b == 'P' {
			term.Restore(int(os.Stdin.Fd()), oldState)

			menu(p)

			// Back to raw mode to keep listening for 'P'.
			oldState, err = term.MakeRaw(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("Erreur terminal :", err)
				return
			}

			fmt.Println("\nAppuyez sur P pour ouvrir le menu.")
		}
	}
}

func menu(p *characters.Character) {
	var choix [1]byte

	fmt.Println("\n=== MENU ===")
	fmt.Println("1. Afficher les informations du personnage")
	fmt.Println("2. Accéder au contenu de l'inventaire")
	fmt.Println("3. Marchand")
	fmt.Println("4. Camp d'entrainement")
	fmt.Println("5. Quitter")
	fmt.Print("Votre choix : ")

	os.Stdin.Read(choix[:])

	switch choix[0] {
	case '1':
		showinfo.DisplayInfo(*p)
	case '2':
		showinfo.DisplayInventory(p)
	case '3':
		npc.Merchant(p)
	case '4':
		combat.TrainingFight(p)
	case '5':
		fmt.Println("Au revoir !")
		os.Exit(0)

	default:
		fmt.Println("Choix invalide")
	}
}
