package main

import (
	"fmt"
	"strconv"

	"github.com/extintor/advent_of_code/shared/utils"
)

type number struct {
	value int
	start int
	end   int
}

func isSymbol(r rune) bool {
	return r != '.' && !('0' <= r && r <= '9')
}

func isGear(r rune) bool {
	return r == '*'
}

func symbolIndex(schematic []string, startX, endX, y int, check func(rune) bool) map[string]struct{} {
	symbols := make(map[string]struct{}, 0)
	for x := startX; x < endX; x++ {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				nx, ny := x+dx, y+dy
				if nx >= 0 && ny >= 0 && ny < len(schematic) && nx < len(schematic[ny]) && check(rune(schematic[ny][nx])) {
					symbols[fmt.Sprintf("%d.%d", nx, ny)] = struct{}{}
				}
			}
		}
	}
	return symbols
}

func getNumbersFromLine(line string) []number {
	result := make([]number, 0)
	for x := 0; x < len(line); {
		char := line[x]
		if '0' <= char && char <= '9' {
			startX := x
			for x < len(line) && '0' <= line[x] && line[x] <= '9' {
				x++
			}
			numStr := line[startX:x]
			num, _ := strconv.Atoi(numStr)
			result = append(result, number{num, startX, x})
		} else {
			x++
		}
	}
	return result
}

func solveOne(schematic []string) int {
	sum := 0
	for y, line := range schematic {
		for _, num := range getNumbersFromLine(line) {
			if len(symbolIndex(schematic, num.start, num.end, y, isSymbol)) != 0 {
				sum += num.value
			}
		}
	}
	return sum
}

func solveTwo(schematic []string) int {
	sum := 0

	gears := make(map[string][]int, 0)

	for y, line := range schematic {
		for _, num := range getNumbersFromLine(line) {
			for k := range symbolIndex(schematic, num.start, num.end, y, isGear) {
				if _, ok := gears[k]; ok == false {
					gears[k] = make([]int, 0)
				}
				gears[k] = append(gears[k], num.value)
			}
		}
	}

	for _, v := range gears {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	return sum
}

func main() {
	input := utils.ReadInput()
	fmt.Println(solveOne(input))
	fmt.Println(solveTwo(input))
}
