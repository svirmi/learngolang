// pagination through slice
package main

import (
	"fmt"
)

func main() {

	const pageSize = 4

	games := []string{
		"pacman", "mario", "tetris", "minegraft",
		"world of tanks", "civilisation", "doom", "need for speed",
		"quake", "ultima", "digger", "dota 2",
		"sudoku",
	}

	l := len(games)

	for from := 0; from < l; from += pageSize {

		to := from + pageSize
		// check to prevent go out of slice range while looping
		if to > l {
			to = l
		}
		currentPage := games[from:to]

		fmt.Printf("Page #%d [%d:%d]\n", (from/pageSize)+1, from, to)
		fmt.Println("", currentPage)

	}
}
