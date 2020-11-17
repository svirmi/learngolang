package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {

	const (
		// board dimensions
		width  = 50
		height = 10

		maxFrames = 1200
		speed     = time.Second / 20

		emptyCell = ' '
		ballCell  = '⚽'

		// drawing buffer length
		// *2 for extra spaces
		// +1 for newlines
		bufLen = (width*2 + 1) * height
	)

	var (
		px, py int
		vx, vy = 1, 1 // velocity
		cell   rune
	)

	// create the board
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	buf := make([]rune, 0, bufLen)

	screen.Clear()

	for i := 0; i < maxFrames; i++ {

		px += vx
		py += vy

		if px <= 0 || px >= width-1 { // ball hits left or right wall
			vx *= -1
		}

		if py <= 0 || py >= height-1 { // ball hits left or right wall
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[px][py] = true

		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]

		// draw the board into the buffer
		for y := range board[0] {
			for x := range board {
				cell = emptyCell
				if board[x][y] {
					cell = ballCell
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		fmt.Println(string(buf))

		time.Sleep(speed)
	}
}
