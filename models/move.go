package models

type Coord struct {
	X int
	Y int
}

const (
	North     = "N"
	South     = "S"
	East      = "E"
	West      = "W"
	NorthEast = "NE"
	NorthWest = "NW"
	SouthEast = "SE"
	SouthWest = "SW"
)

var (
	moves = map[string]Coord{
		North:     {3, 0},
		South:     {-3, 0},
		East:      {0, 3},
		West:      {0, -3},
		NorthEast: {2, 2},
		NorthWest: {2, -2},
		SouthEast: {-2, 2},
		SouthWest: {-2, -2},
	}
)
