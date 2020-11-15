package main

import (
	"fmt"
)

func main() {

	const (
		// board dimensions
		width  = 50
		height = 10

		emptyCell = ' '
		ballCell  = '⚽'
	)

	var cell rune

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	board[0][0] = true

	// draw the board
	for y := range board[0] {
		for x := range board {
			cell = emptyCell
			if board[x][y] {
				cell = ballCell
			}

			fmt.Print(string(cell), " ")
		}
		fmt.Println()
	}
}
