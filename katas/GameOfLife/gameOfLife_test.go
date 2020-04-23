package katas

import (
	"fmt"
	"testing"
)

func isAlive(x int, y int) bool {
	return ((x == 4 || x == 5) && y == 2) || (x == 5 && y == 3)
}

func TestGameOfLife(t *testing.T) {
	var cells = make([]Cell, 0)
	for y := 1; y < 5; y++ {
		for x := 1; x < 9; x++ {
			var cell = Cell{x: x, y: y, alive: isAlive(x, y)}
			cells = append(cells, cell)
		}
	}

	var grid = Grid{cells: cells}
	var nextGrid = CalculateNextGeneration(grid)

	for index := 0; index < len(nextGrid.cells); index++ {
		var previousCell = grid.cells[index]
		var cell = nextGrid.cells[index]
		fmt.Printf("PreviousCell (x: %d, y: %d, alive: %t), NextCell (x: %d, y: %d, alive: %t)", previousCell.x, previousCell.y, previousCell.alive, cell.x, cell.y, cell.alive)
		fmt.Println()
	}
}
