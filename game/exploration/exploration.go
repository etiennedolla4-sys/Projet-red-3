package game

import (
	"Projet-red-3/characters"
	"fmt"
	"math/rand"
)

func Start(character *characters.Character) {
	currentLayer := 1
	x := 2
	y := 2

	gameMap := NewMap(5, 5)

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
		fmt.Println("1. Explorer")
		fmt.Println("2. Observer")
		fmt.Println("3. Descendre")
		fmt.Println("4. Remonter")
		fmt.Println("0. Quitter")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			Explore(character, gameMap, &x, &y, currentLayer)

		case 2:
			Observe(layer, x, y)

		case 3:
			if currentLayer < 5 {
				currentLayer++

				x, y = findStartPosition(gameMap, currentLayer)

				fmt.Printf(
					"\nVous descendez vers la couche %d.\n",
					currentLayer,
				)
			} else {
				fmt.Println()
				fmt.Println(
					"Vous êtes déjà dans la couche la plus profonde connue.",
				)
			}

		case 4:
			if currentLayer > 1 {
				ApplyReturnEffect(character, currentLayer)

				if character.HP <= 0 {
					fmt.Println("Vous vous effondrez...")
					return
				}

				currentLayer--

				x, y = findStartPosition(gameMap, currentLayer)

				fmt.Printf(
					"\nVous remontez vers la couche %d.\n",
					currentLayer,
				)
			} else {
				fmt.Println()
				fmt.Println("Vous êtes déjà à la surface.")
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

func Explore(
	character *characters.Character,
	gameMap Map,
	x *int,
	y *int,
	depth int,
) {
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	direction := directions[rand.Intn(len(directions))]

	newX := *x + direction[0]
	newY := *y + direction[1]

	if !gameMap.IsWalkable(newX, newY, depth) {
		fmt.Println()
		fmt.Println("Vous ne pouvez pas passer par ici.")
		return
	}

	*x = newX
	*y = newY

	fmt.Printf(
		"\nVous avancez jusqu'à la position (%d, %d).\n",
		*x,
		*y,
	)

	tile := gameMap.Tile(*x, *y, depth)

	switch tile {
	case "!":
		fmt.Println("Vous entrez dans une zone dangereuse !")
		StartEncounter(character)

	case "T":
		FindChest(character)

	case "+":
		FindRest(character)

	default:
		if rand.Intn(100) < 30 {
			RandomEvent(character, depth)
		} else {
			fmt.Println("Vous avancez sans rencontrer personne.")
		}
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
