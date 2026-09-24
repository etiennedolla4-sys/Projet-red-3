package menu

import (
	"Projet-red-3/characters"
	"Projet-red-3/game/combat"
	exploration "Projet-red-3/game/exploration"
	blacksmith "Projet-red-3/npc/blacksmith"
	merchant "Projet-red-3/npc/merchant"
	"Projet-red-3/ui/showinfo"
	"Projet-red-3/ui/text"
	"Projet-red-3/utils"
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

	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║                                        ║")
	fmt.Println("║       V E Y R  -  A V E N T U R E      ║")
	fmt.Println("║                                        ║")
	fmt.Println("║       Appuyez sur P pour ouvrir        ║")
	fmt.Println("║              le menu                   ║")
	fmt.Println("║                                        ║")
	fmt.Println("╚════════════════════════════════════════╝")

	for {
		b, err := text.Stdin.ReadByte()
		if err != nil {
			return
		}

		if b == 'p' || b == 'P' {
			term.Restore(int(os.Stdin.Fd()), oldState)

			utils.ClearTerminal()
			menu(p)

			oldState, err = term.MakeRaw(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("Erreur terminal :", err)
				return
			}

			fmt.Println()
			fmt.Println("╔════════════════════════════════════════╗")
			fmt.Println("║                                        ║")
			fmt.Println("║       Appuyez sur P pour ouvrir        ║")
			fmt.Println("║              le menu                   ║")
			fmt.Println("║                                        ║")
			fmt.Println("╚════════════════════════════════════════╝")
		}
	}
}

func menu(p *characters.Character) {
	var choice [1]byte

	utils.ClearTerminal()

	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║                                        ║")
	fmt.Println("║                M E N U                 ║")
	fmt.Println("║                                        ║")
	fmt.Println("╠════════════════════════════════════════╣")
	fmt.Println("║                                        ║")
	fmt.Println("║  1. Afficher les informations          ║")
	fmt.Println("║  2. Inventaire                         ║")
	fmt.Println("║  3. Marchand                           ║")
	fmt.Println("║  4. Forgeron                           ║")
	fmt.Println("║  5. Camp d'entrainement                ║")
	fmt.Println("║  6. Exploration                        ║")
	fmt.Println("║  7. Quitter le jeux                    ║")
	fmt.Println("║                                        ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Print("\nVotre choix : ")

	os.Stdin.Read(choice[:])

	utils.ClearTerminal()

	switch choice[0] {
	case '1':
		showinfo.DisplayInfo(*p)

	case '2':
		showinfo.DisplayInventory(p)

	case '3':
		merchant.Merchant(p)

	case '4':
		blacksmith.Blacksmith(p)

	case '5':
		combat.TrainingFight(p)

	case '6':
		exploration.Start(p)

	case '7':
		fmt.Println("Au revoir !")
		os.Exit(0)

	default:
		fmt.Println("Choix invalide.")
	}
}
