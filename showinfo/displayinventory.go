package showinfo

import (
	"Projet-red-3/menu"
	"fmt"
	"golang.org/x/term"
	"os"
)

func displayInventory() {
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Println("Appuyez sur P pour ouvrir le menu.")

	for {
		var touche [1]byte
		os.Stdin.Read(touche[:])

		if (touche[0] == 'i' || touche[0] == 'I') && in_menu == false {

		}
	}
}
