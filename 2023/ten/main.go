package main

import (
	"fmt"

	"github.com/extintor/advent_of_code/shared/utils"
)

var START_CHAR_SUB = 'J'
var START_DIRECTION = North

type pos struct {
	x int
	y int
}

type step struct {
	direction Direction
	char      rune
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func getPath(grid [][]rune, startPos pos, startDirection Direction) map[pos]step {
	currentDirection := startDirection
	currentPos := startPos
	v := make(map[pos]step)
	for {
		currentChar := grid[currentPos.y][currentPos.x]
		v[currentPos] = step{currentDirection, currentChar}
		switch currentChar {
		case '|':
			if currentDirection == South {
				currentPos = pos{currentPos.x, currentPos.y - 1}
				currentDirection = South
			} else {
				currentPos = pos{currentPos.x, currentPos.y + 1}
				currentDirection = North
			}
		case '-':
			if currentDirection == East {
				currentPos = pos{currentPos.x - 1, currentPos.y}
				currentDirection = East
			} else {
				currentPos = pos{currentPos.x + 1, currentPos.y}
				currentDirection = West
			}
		case 'L':
			if currentDirection == East {
				currentPos = pos{currentPos.x, currentPos.y - 1}
				currentDirection = South
			} else {
				currentPos = pos{currentPos.x + 1, currentPos.y}
				currentDirection = West
			}
		case 'J':
			if currentDirection == North {
				currentPos = pos{currentPos.x - 1, currentPos.y}
				currentDirection = East
			} else {
				currentPos = pos{currentPos.x, currentPos.y - 1}
				currentDirection = South
			}
		case '7':
			if currentDirection == West {
				currentPos = pos{currentPos.x, currentPos.y + 1}
				currentDirection = North
			} else {
				currentPos = pos{currentPos.x - 1, currentPos.y}
				currentDirection = East
			}
		case 'F':
			if currentDirection == East {
				currentPos = pos{currentPos.x, currentPos.y + 1}
				currentDirection = North
			} else {
				currentPos = pos{currentPos.x + 1, currentPos.y}
				currentDirection = West
			}
		}
		if currentPos.x == startPos.x && currentPos.y == startPos.y {
			break
		}
	}
	return v
}

func solveOne(grid [][]rune, start pos, direction Direction) int {
	return len(getPath(grid, start, direction)) / 2
}

func solveTwo(grid [][]rune, start pos, direction Direction) int {
	v := getPath(grid, start, direction)
	var t int
	for y, line := range grid {
		var w int
		for x := range line {
			currentPos := pos{x, y}
			currentValue, currentExists := v[currentPos]
			if currentExists {
				if currentValue.direction == South {
					w--
				} else {
					switch currentValue.char {
					case '|':
						if currentValue.direction == North {
							w++
						}
					case 'F':
						if currentValue.direction == East {
							w++
						}
					case '7':
						if currentValue.direction == West {
							w++
						}
					}
				}
			} else {
				if w != 0 {
					t++
				}
			}
		}
	}
	return t
}

func main() {
	input := utils.ReadInput()

	var start pos

	grid := make([][]rune, 0)

	for y, line := range input {
		gridLine := make([]rune, 0)
		for x, c := range line {
			if c == 'S' {
				start = pos{x, y}
				gridLine = append(gridLine, START_CHAR_SUB)
				continue
			}
			gridLine = append(gridLine, c)
		}
		grid = append(grid, gridLine)
	}

	fmt.Println(solveOne(grid, start, START_DIRECTION))
	fmt.Println(solveTwo(grid, start, START_DIRECTION))

}
