package game

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
