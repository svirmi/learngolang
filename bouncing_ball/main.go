package main

func main() {

	const (
		// board dimensions
		width  = 50
		height = 10
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	// draw the board
}
