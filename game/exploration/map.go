package game

import "fmt"

type Map struct {
	Width   int
	Height  int
	Cleared map[string]bool
}

func NewMap(width, height int) Map {
	return Map{
		Width:   width,
		Height:  height,
		Cleared: make(map[string]bool),
	}
}

func (m Map) tileKey(x, y, depth int) string {
	return fmt.Sprintf("%d:%d:%d", depth, x, y)
}

func (m Map) MarkCleared(x, y, depth int) {
	m.Cleared[m.tileKey(x, y, depth)] = true
}

func (m Map) RemainingEnemies(depth int) int {
	count := 0

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.baseTile(x, y, depth) == "E" && !m.Cleared[m.tileKey(x, y, depth)] {
				count++
			}
		}
	}

	return count
}

func (m Map) DefeatedEnemies(depth int) int {
	count := 0

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.baseTile(x, y, depth) == "E" && m.Cleared[m.tileKey(x, y, depth)] {
				count++
			}
		}
	}

	return count
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
	legend := []string{
		"@ = Vous",
		". = Chemin",
		"# = Mur",
		"E = Ennemi",
		"T = Trésor",
		"+ = Repos",
	}

	fmt.Println()
	fmt.Println("+-----------------+    +---------------------+")
	fmt.Printf("|   CARTE - %d     |    |       LEGENDE       |\n", depth)
	fmt.Println("+-----------------+    +---------------------+")

	for y := m.Height - 1; y >= 0; y-- {
		fmt.Print("| ")

		for x := 0; x < m.Width; x++ {
			if x == playerX && y == playerY {
				fmt.Print("@ ")
			} else {
				fmt.Printf("%s ", m.Tile(x, y, depth))
			}
		}

		fmt.Print("      ")

		legendIndex := m.Height - 1 - y

		if legendIndex < len(legend) {
			fmt.Printf("| %-19s |\n", legend[legendIndex])
		} else {
			fmt.Println("|                     |")
		}
	}

	fmt.Println("+-----------------+    +---------------------+")
}

func (m Map) Tile(x, y, depth int) string {
	if m.Cleared[m.tileKey(x, y, depth)] {
		return "."
	}

	return m.baseTile(x, y, depth)
}

func (m Map) baseTile(x, y, depth int) string {
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
	if x == 0 && y == 6 {
		return "T"
	}

	if x == 6 && y == 0 {
		return "+"
	}

	if x == 1 && y == 5 {
		return "E"
	}

	if x == 5 && y == 1 {
		return "E"
	}

	return "."
}

func ruinsTile(x, y int) string {
	if x == 1 && y == 5 {
		return "#"
	}

	if x == 5 && y == 1 {
		return "T"
	}

	if x == 6 && y == 5 {
		return "E"
	}

	if x == 1 && y == 1 {
		return "E"
	}

	return "."
}

func forestTile(x, y int) string {
	if x == 1 && y == 1 {
		return "E"
	}

	if x == 5 && y == 6 {
		return "T"
	}

	if x == 6 && y == 2 {
		return "+"
	}

	if x == 2 && y == 5 {
		return "E"
	}

	if x == 5 && y == 1 {
		return "E"
	}

	return "."
}

func caveTile(x, y int) string {
	if x == 0 || x == 6 {
		return "#"
	}

	if x == 3 && y == 6 {
		return "E"
	}

	if x == 5 && y == 1 {
		return "T"
	}

	if x == 1 && y == 3 {
		return "+"
	}

	if x == 3 && y == 2 {
		return "E"
	}

	return "."
}

func abyssTile(x, y int) string {
	if x == 0 || x == 6 || y == 0 {
		return "#"
	}

	if x == 3 && y == 5 {
		return "E"
	}

	if x == 5 && y == 3 {
		return "T"
	}

	if x == 2 && y == 2 {
		return "E"
	}

	if x == 4 && y == 4 {
		return "E"
	}

	return "."
}
