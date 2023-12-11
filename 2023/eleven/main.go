package main

import (
	"fmt"

	"github.com/extintor/advent_of_code/shared/utils"
)

type position struct {
	x int
	y int
}

func distance(a, b position, rows, columns map[int]struct{}, multiplier int) int {
	distance := 0
	x1, x2 := a.x, b.x
	if a.x > b.x {
		x2, x1 = a.x, b.x
	}
	for x := x1; x < x2; x++ {
		d := 1
		if _, ok := columns[x]; !ok {
			d *= multiplier
		}
		distance += d
	}
	for y := a.y; y < b.y; y++ {
		d := 1
		if _, ok := rows[y]; !ok {
			d *= multiplier
		}
		distance += d
	}
	return distance
}

func solveOne(stars []position, rows, columns map[int]struct{}) int {
	sum := 0
	for i := 0; i < len(stars); i++ {
		for j := i + 1; j < len(stars); j++ {
			d := distance(stars[i], stars[j], rows, columns, 2)
			sum += d
		}
	}
	return sum
}

func solveTwo(stars []position, rows, columns map[int]struct{}) int {
	sum := 0
	for i := 0; i < len(stars); i++ {
		for j := i + 1; j < len(stars); j++ {
			d := distance(stars[i], stars[j], rows, columns, 1_000_000)
			sum += d
		}
	}
	return sum
}

func main() {
	rows := make(map[int]struct{}, 0)
	columns := make(map[int]struct{}, 0)
	stars := make([]position, 0)
	for i, line := range utils.ReadInput() {
		for j, char := range line {
			if char == '#' {
				rows[i] = struct{}{}
				columns[j] = struct{}{}
				stars = append(stars, position{j, i})
			}
		}
	}
	fmt.Println(solveOne(stars, rows, columns))
	fmt.Println(solveTwo(stars, rows, columns))
}
