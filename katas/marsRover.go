package katas

type Grid struct {
	maxX int
	maxY int
}

type MarsRover struct {
	direction string
	x         int
	y         int
}

func exceededBoundary(value int, maxValue int) bool {
	return value > maxValue
}

func moveEast(rover *MarsRover, grid Grid) {
	var newPosition = rover.x + 1
	if exceededBoundary(newPosition, grid.maxX) {
		newPosition = 0
	}
	rover.x = newPosition
	rover.direction = "e"
}

func moveWest(rover *MarsRover, grid Grid) {
	var newPosition = rover.x - 1
	if newPosition < 0 {
		newPosition = grid.maxX
	}
	rover.x = newPosition
	rover.direction = "w"
}

func moveSouth(rover *MarsRover, grid Grid) {
	var newPosition = rover.y - 1
	if newPosition < 0 {
		newPosition = grid.maxY
	}
	rover.y = newPosition
	rover.direction = "s"
}

func moveNorth(rover *MarsRover, grid Grid) {
	var newPosition = rover.y + 1
	if exceededBoundary(newPosition, grid.maxY) {
		newPosition = 0
	}
	rover.y = newPosition
	rover.direction = "n"
}

func (rover *MarsRover) forward(grid Grid) {
	switch rover.direction {
	case "n":
		moveSouth(rover, grid)
	case "s":
		moveSouth(rover, grid)
	case "e":
		moveSouth(rover, grid)
	case "w":
		moveSouth(rover, grid)
	}
}

func (rover *MarsRover) backward(grid Grid) {
	switch rover.direction {
	case "n":
		moveSouth(rover, grid)
	case "s":
		moveNorth(rover, grid)
	case "e":
		moveWest(rover, grid)
	case "w":
		moveEast(rover, grid)
	}
}

func (rover *MarsRover) left(grid Grid) {
	switch rover.direction {
	case "n":
		moveSouth(rover, grid)
	case "s":
		moveSouth(rover, grid)
	case "e":
		moveSouth(rover, grid)
	case "w":
		moveSouth(rover, grid)
	}
}

func (rover *MarsRover) right(grid Grid) {
	switch rover.direction {
	case "n":
		moveSouth(rover, grid)
	case "s":
		moveSouth(rover, grid)
	case "e":
		moveSouth(rover, grid)
	case "w":
		moveSouth(rover, grid)
	}
}
