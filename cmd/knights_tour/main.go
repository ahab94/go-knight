package main

import (
	"log"
	"os"

	"github.com/ahab94/go-knight/config"
	"github.com/ahab94/go-knight/models"
	"github.com/spf13/viper"
)

func main() {
	dimension := 10
	startX := viper.GetInt(config.StartX)
	startY := viper.GetInt(config.StartY)

	board, err := models.NewBoard(dimension, startX, startY)
	if err != nil {
		panic(err)
	}

	moves := board.AvailableMoves(startX, startY)
	for i := 0; i < dimension*dimension; i++ {
		visited := board.Visited()
		log.Println("points visited: ", visited)
		if visited == dimension*dimension {
			os.Exit(0)
		}

		coord := board.LeastMovesCoord(moves)
		if err := board.Move(coord.X, coord.Y); err != nil {
			panic(err)
		}
		moves = board.AvailableMoves(coord.X, coord.Y)
	}
}
