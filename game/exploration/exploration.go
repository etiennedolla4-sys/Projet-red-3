package game

import (
	"Projet-red-3/characters"
	"fmt"
)

func Start(character *characters.Character) {
	currentLayer := 1
	x := 3
	y := 3

	gameMap := NewMap(7, 7)

	for {
		if character.HP <= 0 {
			fmt.Println()
			fmt.Println("Vous n'avez plus assez de forces.")
			fmt.Println("Votre exploration prend fin.")
			return
		}

		layer := GetLayer(currentLayer)

		fmt.Println()
		fmt.Println("=======================================")
		fmt.Println("          LES PROFONDEURS")
		fmt.Println("=======================================")
		fmt.Printf("Couche : %d - %s\n", layer.Depth, layer.Name)
		fmt.Println(layer.Description)

		gameMap.Display(x, y, currentLayer)

		fmt.Println()
		fmt.Printf("PV : %d / %d\n", character.HP, character.MaxHP)
		fmt.Printf("Or : %d\n", character.Gold)

		fmt.Println()
		fmt.Println("1. Se déplacer")
		fmt.Println("2. Observer")
		fmt.Println("3. Descendre")
		fmt.Println("0. Quitter")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			Move(character, gameMap, &x, &y, currentLayer)

		case 2:
			Observe(layer, x, y)

		case 3:
			if currentLayer < 5 {
				currentLayer++

				x, y = findStartPosition(gameMap, currentLayer)

				fmt.Println()
				fmt.Printf(
					"Vous descendez vers la couche %d.\n",
					currentLayer,
				)
			} else {
				fmt.Println()
				fmt.Println(
					"Vous êtes déjà dans la couche la plus profonde connue.",
				)
			}

		case 0:
			fmt.Println()
			fmt.Println("Vous quittez les profondeurs.")
			return

		default:
			fmt.Println()
			fmt.Println("Choix invalide.")
		}
	}
}

func Move(
	character *characters.Character,
	gameMap Map,
	x *int,
	y *int,
	depth int,
) {
	fmt.Println()
	fmt.Println("========== DÉPLACEMENT ==========")
	fmt.Println("1. Aller en haut")
	fmt.Println("2. Aller en bas")
	fmt.Println("3. Aller à gauche")
	fmt.Println("4. Aller à droite")
	fmt.Println("0. Annuler")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	newX := *x
	newY := *y

	switch choice {
	case 1:
		newY++

	case 2:
		newY--

	case 3:
		newX--

	case 4:
		newX++

	case 0:
		return

	default:
		fmt.Println()
		fmt.Println("Choix invalide.")
		return
	}

	if !gameMap.IsWalkable(newX, newY, depth) {
		fmt.Println()
		fmt.Println("Vous ne pouvez pas aller dans cette direction.")
		return
	}

	*x = newX
	*y = newY

	fmt.Println()
	fmt.Printf(
		"Vous vous déplacez vers la position (%d, %d).\n",
		*x,
		*y,
	)

	HandleTile(character, gameMap, *x, *y, depth)
}

func HandleTile(
	character *characters.Character,
	gameMap Map,
	x int,
	y int,
	depth int,
) {
	tile := gameMap.Tile(x, y, depth)

	switch tile {
	case "!":
		fmt.Println()
		fmt.Println("Vous entrez dans une zone dangereuse !")
		StartEncounter(character, depth)

	case "T":
		FindChest(character)

	case "+":
		FindRest(character)

	default:
		fmt.Println("Vous avancez sans rencontrer personne.")
	}
}

func findStartPosition(gameMap Map, depth int) (int, int) {
	for y := 0; y < gameMap.Height; y++ {
		for x := 0; x < gameMap.Width; x++ {
			if gameMap.IsWalkable(x, y, depth) {
				return x, y
			}
		}
	}

	return 0, 0
}

func Observe(layer Layer, x, y int) {
	fmt.Println()
	fmt.Println("========== OBSERVATION ==========")
	fmt.Printf("Position : (%d, %d)\n", x, y)
	fmt.Printf("Zone : %s\n", layer.Name)
	fmt.Println(layer.Description)
	fmt.Println("Vous ne remarquez rien d'inhabituel.")
}
