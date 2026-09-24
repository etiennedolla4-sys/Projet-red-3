package game

import "fmt"

type Map struct {
	Width  int
	Height int
}

func NewMap(width, height int) Map {
	return Map{
		Width:  width,
		Height: height,
	}
}

func (m Map) IsInside(x, y int) bool {
	return x >= 0 && x < m.Width && y >= 0 && y < m.Height
}

func (m Map) IsWalkable(x, y, depth int) bool {
	if !m.IsInside(x, y) {
		return false
	}

	return m.Tile(x, y, depth) != "#"
}

func (m Map) Display(playerX, playerY, depth int) {
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Printf("          CARTE - COUCHE %d\n", depth)
	fmt.Println("=======================================")

	for y := m.Height - 1; y >= 0; y-- {
		for x := 0; x < m.Width; x++ {
			if x == playerX && y == playerY {
				fmt.Print("@ ")
			} else {
				fmt.Printf("%s ", m.Tile(x, y, depth))
			}
		}

		fmt.Println()
	}

	fmt.Println("=======================================")
	fmt.Println("@ = Vous")
	fmt.Println(". = Chemin")
	fmt.Println("# = Mur")
	fmt.Println("T = Trésor")
	fmt.Println("! = Danger")
	fmt.Println("+ = Repos")
	fmt.Println("=======================================")
}

func (m Map) Tile(x, y, depth int) string {
	switch depth {
	case 1:
		return surfaceTile(x, y)
	case 2:
		return ruinsTile(x, y)
	case 3:
		return forestTile(x, y)
	case 4:
		return caveTile(x, y)
	case 5:
		return abyssTile(x, y)
	default:
		return "."
	}
}

func surfaceTile(x, y int) string {
	if x == 0 && y == 4 {
		return "T"
	}

	if x == 4 && y == 0 {
		return "+"
	}

	return "."
}

func ruinsTile(x, y int) string {
	if x == 1 && y == 3 {
		return "#"
	}

	if x == 3 && y == 1 {
		return "T"
	}

	if x == 4 && y == 4 {
		return "!"
	}

	return "."
}

func forestTile(x, y int) string {
	if x == 1 && y == 1 {
		return "!"
	}

	if x == 3 && y == 4 {
		return "T"
	}

	if x == 4 && y == 2 {
		return "+"
	}

	return "."
}

func caveTile(x, y int) string {
	if x == 0 || x == 4 {
		return "#"
	}

	if x == 2 && y == 4 {
		return "!"
	}

	if x == 3 && y == 1 {
		return "T"
	}

	if x == 1 && y == 2 {
		return "+"
	}

	return "."
}

func abyssTile(x, y int) string {
	if x == 0 || x == 4 || y == 0 {
		return "#"
	}

	if x == 2 && y == 3 {
		return "!"
	}

	if x == 3 && y == 2 {
		return "T"
	}

	return "."
}
