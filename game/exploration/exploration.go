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
			fmt.Println("\nVous n'avez plus assez de forces.")
			fmt.Println("Votre exploration prend fin.")
			return
		}

		layer := GetLayer(currentLayer)

		fmt.Println("\n================================")
		fmt.Println("       LES PROFONDEURS")
		fmt.Println("================================")
		fmt.Printf("Couche : %d - %s\n", layer.Depth, layer.Name)
		fmt.Println(layer.Description)
		fmt.Printf("Position : (%d, %d)\n", x, y)
		fmt.Printf("PV : %d / %d\n", character.HP, character.MaxHP)
		fmt.Printf("Or : %d\n", character.Gold)

		fmt.Println("\n1. Explorer")
		fmt.Println("2. Observer")
		fmt.Println("3. Descendre")
		fmt.Println("4. Remonter")
		fmt.Println("0. Quitter")
		fmt.Print("\nVotre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			Explore(character, gameMap, &x, &y)

		case 2:
			Observe(layer, x, y)

		case 3:
			if currentLayer < 5 {
				currentLayer++
				x = rand.Intn(gameMap.Width)
				y = rand.Intn(gameMap.Height)

				fmt.Printf(
					"\nVous descendez vers la couche %d.\n",
					currentLayer,
				)
			} else {
				fmt.Println("\nVous êtes déjà dans la couche la plus profonde connue.")
			}

		case 4:
			if currentLayer > 1 {
				ApplyReturnEffect(character, currentLayer)
				currentLayer--

				fmt.Printf(
					"\nVous remontez vers la couche %d.\n",
					currentLayer,
				)
			} else {
				fmt.Println("\nVous êtes déjà à la surface.")
			}

		case 0:
			fmt.Println("\nVous quittez les profondeurs.")
			return

		default:
			fmt.Println("\nChoix invalide.")
		}
	}
}

func Explore(
	character *characters.Character,
	gameMap Map,
	x *int,
	y *int,
) {
	direction := rand.Intn(4)

	switch direction {
	case 0:
		if gameMap.IsInside(*x, *y+1) {
			*y = *y + 1
		}

	case 1:
		if gameMap.IsInside(*x, *y-1) {
			*y = *y - 1
		}

	case 2:
		if gameMap.IsInside(*x+1, *y) {
			*x = *x + 1
		}

	case 3:
		if gameMap.IsInside(*x-1, *y) {
			*x = *x - 1
		}
	}

	fmt.Printf(
		"\nVous avancez jusqu'à la position (%d, %d).\n",
		*x,
		*y,
	)

	RandomEvent(character)
}

func Observe(layer Layer, x, y int) {
	fmt.Println("\n========== OBSERVATION ==========")
	fmt.Printf("Position : (%d, %d)\n", x, y)
	fmt.Printf("Zone : %s\n", layer.Name)
	fmt.Println(layer.Description)
	fmt.Println("Vous ne remarquez rien d'inhabituel.")
}
