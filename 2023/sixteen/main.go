package main

import (
	"fmt"

	"github.com/obrahc/advent_of_code/shared/utils"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type pos struct {
	x int
	y int
}

type Tile struct {
	traversedDirections []Direction
}

type Vector struct {
	position  pos
	direction Direction
}

func toRuneList(s string) []rune {
	return []rune(s)
}

func convertInput(s []string) [][]rune {
	r := make([][]rune, len(s))
	for i, l := range s {
		r[i] = toRuneList(l)
	}
	return r
}

var cache map[Vector]struct{} = make(map[Vector]struct{})
var resultCache map[pos]struct{} = make(map[pos]struct{})

func processStep(grid [][]rune, position pos, direction Direction) {
	if position.x < 0 || position.x >= len(grid[0]) || position.y < 0 || position.y >= len(grid) {
		return
	}
	char := grid[position.y][position.x]
	v := Vector{position, direction}
	if _, ok := cache[v]; ok {
		return
	}
	cache[v] = struct{}{}
	resultCache[position] = struct{}{}

	switch char {
	case '.':
	case '/':
		switch direction {
		case North:
			direction = East
		case East:
			direction = North
		case South:
			direction = West
		case West:
			direction = South
		}
	case '\\':
		switch direction {
		case North:
			direction = West
		case East:
			direction = South
		case South:
			direction = East
		case West:
			direction = North
		}
	case '-':
		if direction == North || direction == South {
			processStep(grid, pos{position.x + 1, position.y}, East)
			processStep(grid, pos{position.x - 1, position.y}, West)
			return
		}
	case '|':
		if direction == East || direction == West {
			processStep(grid, pos{position.x, position.y + 1}, South)
			processStep(grid, pos{position.x, position.y - 1}, North)
			return
		}
	}

	switch direction {
	case North:
		position.y--
	case East:
		position.x++
	case South:
		position.y++
	case West:
		position.x--
	}

	processStep(grid, position, direction)
}

func solve(input [][]rune, position pos, direction Direction) int {
  cache = make(map[Vector]struct{})
  resultCache = make(map[pos]struct{})
  processStep(input, position, direction)
  return len(resultCache)
}

func solveOne(input [][]rune) int {
	processStep(input, pos{0, 0}, East)
	return len(resultCache)
}

func solveTwo(input [][]rune) int {
	max := 0
	for i, r := range input {
		for j := range r {
			if i == 0 {
        n := solve(input, pos{j, i}, South)
				if n > max {
					max = n
				}
			}
			if j == 0 {
        n := solve(input, pos{j, i}, East)
				if n > max {
					max = n
				}
			}
			if i == len(input)-1 {
        n := solve(input, pos{j, i}, North)
				if n > max {
					max = n
				}
			}
			if j == len(r)-1 {
        n := solve(input, pos{j, i}, West)
				if n > max {
					max = n
				}
			}
		}
	}
  return max
}

func main() {
	input := convertInput(utils.ReadInput())
	fmt.Println(solveOne(input))
	fmt.Println(solveTwo(input))
}
