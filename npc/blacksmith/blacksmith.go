package blacksmith

import (
	"Projet-red-3/characters"
	"Projet-red-3/ui/text"
	"fmt"
)

var blacksmithFirstTime = true

func Blacksmith(p *characters.Character) {
	message := fmt.Sprintf(
		"Une chaleur étouffante vous frappe lorsque vous entrez dans la forge. Le bruit du marteau résonne contre les murs tandis que des étincelles volent dans tous les sens. Derrière son enclume, un forgeron imposant vous observe.\n%s : « Tu veux renforcer ton équipement ? Alors montre-moi ce que tu as. Une bonne armure peut faire la différence entre rentrer vivant et finir au fond de Veyr. »",
		text.Color("Gareth", "31"),
	)

	if blacksmithFirstTime {
		text.PrintSlow(message)
		blacksmithFirstTime = false
	}

	for {
		fmt.Println("\n========== FORGERON ==========")
		fmt.Println("Gold :", p.Gold)
		fmt.Println()
		fmt.Println("1. Fabriquer un équipement")
		fmt.Println("2. Équiper un équipement")
		fmt.Println("3. Voir mes équipements")
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			craftingMenu(p)

		case 2:
			equipmentMenu(p)

		case 3:
			displayEquipment(p)

		case 0:
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func craftingMenu(p *characters.Character) {
	fmt.Println("\n===== FABRICATION =====")
	fmt.Println("1. Chapeau de l'aventurier")
	fmt.Println("   Plume de Corbeau + Cuir de Sanglier")
	fmt.Println()
	fmt.Println("2. Tunique de l'aventurier")
	fmt.Println("   2x Fourrure de Loup + Peau de Troll")
	fmt.Println()
	fmt.Println("3. Bottes de l'aventurier")
	fmt.Println("   Fourrure de Loup + Cuir de Sanglier")
	fmt.Println()
	fmt.Println("Coût : 5 Gold")
	fmt.Println("0. Retour")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	if choice != 0 {
		CraftEquipment(p, choice)
	}
}

func equipmentMenu(p *characters.Character) {
	fmt.Println("\n===== ÉQUIPEMENT =====")

	found := false
	var indexes []int

	for i, inventoryItem := range p.Inventory {
		if inventoryItem.Type() == "head" ||
			inventoryItem.Type() == "chest" ||
			inventoryItem.Type() == "feet" {

			indexes = append(indexes, i)

			fmt.Printf(
				"%d. %s\n",
				len(indexes),
				inventoryItem.Name(),
			)

			found = true
		}
	}

	if !found {
		fmt.Println("Vous n'avez aucun équipement à équiper.")
		return
	}

	fmt.Println("0. Retour")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(indexes) {
		fmt.Println("Choix invalide.")
		return
	}

	EquipEquipment(p, indexes[choice-1])
}

func displayEquipment(p *characters.Character) {
	fmt.Println("\n===== ÉQUIPEMENTS ÉQUIPÉS =====")

	if p.Equipment.Head == "" {
		fmt.Println("Tête  : Aucun")
	} else {
		fmt.Println("Tête  :", p.Equipment.Head)
	}

	if p.Equipment.Chest == "" {
		fmt.Println("Torse : Aucun")
	} else {
		fmt.Println("Torse :", p.Equipment.Chest)
	}

	if p.Equipment.Feet == "" {
		fmt.Println("Pieds : Aucun")
	} else {
		fmt.Println("Pieds :", p.Equipment.Feet)
	}

	fmt.Printf("\nPV maximum : %d\n", p.MaxHP)
}
