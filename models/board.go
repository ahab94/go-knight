package models

import (
	"errors"
	"fmt"
	wraperrors "github.com/pkg/errors"
)

type board [][]int

func NewBoard(dimensions, startX, startY int) (board, error) {
	if startX >= dimensions || startY >= dimensions {
		return nil, errors.New("starting coordinates should not be greater than 10")
	}
	if startX < 0 || startY < 0 {
		return nil, errors.New("starting coordinates should not be less than 0")
	}

	brd := make(board, dimensions)
	for i := range brd {
		brd[i] = make([]int, dimensions)
	}
	if err := brd.Move(startX, startY, 0); err != nil {
		return nil, wraperrors.Wrap(err, "failed to create board")
	}

	return brd, nil
}

func (b board) AvailableMoves(x, y int) []Coord {
	possibleMoves := make([]Coord, 0)
	for _, m := range moves {
		if x+m.X < len(b) && y+m.Y < len(b) && x+m.X >= 0 && y+m.Y >= 0 && b[m.X+x][m.Y+y] == 0 {
			possibleMoves = append(possibleMoves, Coord{X: m.X + x, Y: m.Y + y})
		}

	}
	return possibleMoves
}

func (b board) LeastMovesCoord(moves []Coord) Coord {
	moveCount := len(b)
	var move Coord
	for _, m := range moves {
		if len(b.AvailableMoves(m.X, m.Y)) < moveCount {
			move = Coord{X: m.X, Y: m.Y}
			moveCount = len(b.AvailableMoves(m.X, m.Y))
		}
	}
	return move
}

func (b board) Move(x, y, visited int) error {
	if x >= len(b) || y >= len(b) || x < 0 || y < 0 {
		return errors.New(fmt.Sprintf("position {%v,%v} moves cursor out of bounds", x, y))
	}
	if b[x][y] != 0 {
		return errors.New(fmt.Sprintf("position {%v,%v} already taken", x, y))
	}

	b[x][y] = visited + 1
	b.View()
	return nil
}

func (b board) Visited() int {
	count := 0
	for xi, x := range b {
		for yi, _ := range x {
			if b[xi][yi] != 0 {
				count++
			}
		}
	}
	return count
}

func (b board) View() {
	for _, x := range b {
		for _, y := range x {
			print(y, "\t")
		}
		println()
	}

}
