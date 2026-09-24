package game

import (
	"Projet-red-3/characters"
	"Projet-red-3/inventory"
	"Projet-red-3/utils"
	"fmt"
)

var gameMap = NewMap(7, 7)
var currentLayer = 1
var playerX = 3
var playerY = 3

func Start(character *characters.Character) {
	utils.ClearTerminal()

	for {
		if character.HP <= 0 {
			if inventory.TryRevive(character) {
				continue
			}

			utils.ClearTerminal()

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

		gameMap.Display(playerX, playerY, currentLayer)

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
			Move(
				character,
				gameMap,
				&playerX,
				&playerY,
				currentLayer,
			)

		case 2:
			Observe(layer, playerX, playerY)

		case 3:
			if currentLayer < 5 {
				defeated := gameMap.DefeatedEnemies(currentLayer)
				remaining := gameMap.RemainingEnemies(currentLayer)

				if defeated < 2 || remaining > 0 {
					fmt.Println()
					fmt.Println("Vous ne pouvez pas encore descendre.")
					fmt.Printf("Ennemis éliminés : %d\n", defeated)
					fmt.Printf("Ennemis restants : %d\n", remaining)
					fmt.Println("Vous devez éliminer au moins 2 ennemis et ne laisser aucun ennemi sur cette couche.")
					break
				}

				currentLayer++

				playerX, playerY = findStartPosition(
					gameMap,
					currentLayer,
				)
			} else {
				fmt.Println()
				fmt.Println(
					"Vous êtes déjà dans la couche la plus profonde connue.",
				)
			}

		case 0:
			utils.ClearTerminal()

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

	utils.ClearTerminal()

	HandleTile(
		character,
		gameMap,
		*x,
		*y,
		depth,
	)
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
	case "E":
		fmt.Println()
		fmt.Println("========== ENNEMI ==========")
		fmt.Println("Un ennemi vous bloque le passage !")

		if StartEncounter(character, depth) {
			gameMap.MarkCleared(x, y, depth)
		}

	case "T":
		fmt.Println()
		ShowChestLore()
		fmt.Println("========== TRÉSOR ==========")

		FindChest(character)
		gameMap.MarkCleared(x, y, depth)

	case "+":
		fmt.Println()
		ShowRestLore()
		fmt.Println("========== REPOS ==========")

		FindRest(character)
		gameMap.MarkCleared(x, y, depth)

	default:
		fmt.Println()
		fmt.Printf(
			"Vous avancez vers la position (%d, %d).\n",
			x,
			y,
		)
	}
}

func findStartPosition(
	gameMap Map,
	depth int,
) (int, int) {
	centerX := gameMap.Width / 2
	centerY := gameMap.Height / 2

	if gameMap.IsWalkable(centerX, centerY, depth) {
		return centerX, centerY
	}

	for y := 0; y < gameMap.Height; y++ {
		for x := 0; x < gameMap.Width; x++ {
			if gameMap.IsWalkable(x, y, depth) {
				return x, y
			}
		}
	}

	return 0, 0
}

func Observe(
	layer Layer,
	x int,
	y int,
) {
	fmt.Println()
	fmt.Println("========== OBSERVATION ==========")
	fmt.Printf("Position : (%d, %d)\n", x, y)
	fmt.Printf("Zone : %s\n", layer.Name)
	fmt.Println(layer.Description)
	fmt.Println()
	fmt.Println("Vous observez les environs.")
	fmt.Println("Les ennemis visibles apparaissent sur la carte.")
}
