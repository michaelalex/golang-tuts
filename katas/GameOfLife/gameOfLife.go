package katas

import (
	"fmt"
)

type Cell struct {
	x     int
	y     int
	alive bool
}

type Grid struct {
	cells []Cell
}

type Neighbours struct {
	numberAlive int
	cells       []Cell
}

func (grid *Grid) findNeighbour(cell Cell, xDiff int, yDiff int) (Cell, bool) {
	var x = cell.x + xDiff
	var y = cell.y + yDiff

	for index := 0; index < len(grid.cells); index++ {
		if grid.cells[index].x == x && grid.cells[index].y == y {
			return grid.cells[index], true
		}
	}

	return Cell{}, false
}

func (cell *Cell) getNeighbours(grid Grid) Neighbours {
	var neighbours = Neighbours{cells: make([]Cell, len(grid.cells)), numberAlive: 0}
	var cellsToCheck = []int{-1, 0, 1}

	for x := 0; x < len(cellsToCheck); x++ {
		var xValue = cellsToCheck[x]
		for y := 0; y < len(cellsToCheck); y++ {
			var yValue = cellsToCheck[y]
			if yValue == 0 && xValue == 0 {
				continue
			}

			nextCell, found := grid.findNeighbour(*cell, x, y)
			if found == true {
				neighbours.cells = append(neighbours.cells, nextCell)
			}
		}
	}

	fmt.Println(neighbours)
	return neighbours
}

func CalculateNextGeneration(grid Grid) Grid {
	var nextGrid = grid

	for index := 0; index < len(grid.cells); index++ {
		var neighbours = grid.cells[index].getNeighbours(grid)
		if grid.cells[index].alive {
			nextGrid.cells[index].alive = neighbours.numberAlive == 2 || neighbours.numberAlive == 3
		} else {
			nextGrid.cells[index].alive = neighbours.numberAlive == 3
		}
	}
	return nextGrid
}
