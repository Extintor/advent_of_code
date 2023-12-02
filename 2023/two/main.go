package main

import (
	"fmt"
	"strconv"
	"strings"

  "github.com/Extintor/advent_of_code/shared/utils"
)

type game struct {
	id    int
	draws []draw
}

type draw struct {
	red   int
	green int
	blue  int
}

func solveOne(input []game) (int, error) {
	sum := 0
	for _, v := range input {
		possible := true
		for _, d := range v.draws {
			if d.red > 12 || d.green > 13 || d.blue > 14 {
				possible = false
				break
			}
		}
		if possible {
			sum += v.id
		}
	}
	return sum, nil
}

func solveTwo(input []game) (int, error) {
	sum := 0
	for _, v := range input {
		minr, ming, minb := 0, 0, 0
		for _, d := range v.draws {
			if d.red > minr {
				minr = d.red
			}
			if d.green > ming {
				ming = d.green
			}
			if d.blue > minb {
				minb = d.blue
			}
		}
		sum += (minr * ming * minb)
	}
	return sum, nil
}

func parseDraw(input string) (draw, error) {
	var r, g, b int
	for _, d := range strings.Split(input, ", ") {
		t := strings.Split(d, " ")
		var err error
		switch t[1] {
		case "red":
			r, err = strconv.Atoi(t[0])
		case "green":
			g, err = strconv.Atoi(t[0])
		case "blue":
			b, err = strconv.Atoi(t[0])
		}
		if err != nil {
			return draw{}, err
		}
	}
	return draw{r, g, b}, nil
}

func main() {
	input := make([]game, 0)
  for _, t := range utils.ReadInput() {
		draws := make([]draw, 0)
		i := strings.Split(t, ": ")
		id, err := strconv.Atoi(strings.Split(i[0], " ")[1])
		if err != nil {
			panic(err)
		}
		for _, v := range strings.Split(i[1], "; ") {
			draw, err := parseDraw(v)
			if err != nil {
				panic(err)
			}
			draws = append(draws, draw)
		}
		input = append(input, game{id, draws})
	}
	fmt.Println(solveOne(input))
	fmt.Println(solveTwo(input))
}
