package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/obrahc/advent_of_code/shared/utils"
)

type direction int

const (
  Right direction = iota
  Down
  Left
  Up
  Unreachable = 1
)

var moveValue map[direction]point = map[direction]point{Right: point{0, 1}, Down: point{1, 0}, Left: point{0, -1}, Up: point{-1, 0}}

type point struct {
  x, y int
}

type move struct {
  direction direction
  value int
}

func stringToDirection(r string) direction {
  switch r {
  case "R":
  return Right
  case "D":
  return Down
  case "L":
  return Left
  case "U":
  return Up
  }
  panic(Unreachable) 
}

func intStringToDirection(r string) direction {
  switch r {
  case "0":
  return Right
  case "1":
  return Down
  case "2":
  return Left
  case "3":
  return Up
  }
  panic(Unreachable) 
}

func shoelace(edge []point) int {
  area := 0
  for i := 0; i < len(edge) - 1; i++  {
    area += edge[i].x *  edge[i + 1].y - edge[i + 1].x * edge[i].y
  }
  return int(math.Abs(float64(area))) / 2
}

func solveOne(input []string) int {
  edges := make([]point, 0)
  perimeter := 0
  x := 0
  y := 0
  for _, line := range input {
    v := strings.Fields(line)
    mv, _ := strconv.Atoi(v[1])    
    m := move{stringToDirection(v[0]), mv}
    vector := moveValue[m.direction]
    edges = append(edges, point{x, y})
    perimeter += m.value
    x += vector.x * m.value
    y += vector.y * m.value
  }
  area := shoelace(edges)
  intArea := area - perimeter / 2 + 1
  return intArea + perimeter
}

func solveTwo(input []string) int {
  borders := make([]point, 0)
  perimeter := 0
  x := 0
  y := 0
  for _, line := range input {
    v := strings.Fields(line)
    hv := strings.Trim(v[len(v)-1], "(#)")
    mv, _ := strconv.ParseInt(hv[0:5], 16, 64)    
    m := move{intStringToDirection(string(hv[len(hv) - 1])), int(mv)}
    vector := moveValue[m.direction]
    borders = append(borders, point{x, y})
    perimeter += m.value
    x += vector.x * m.value
    y += vector.y * m.value
  }
  area := shoelace(borders)
  intArea := area - perimeter / 2 + 1
  return intArea + perimeter
}

func main() {
  input := utils.ReadInput()

  fmt.Println(solveOne(input))
  fmt.Println(solveTwo(input))
}
