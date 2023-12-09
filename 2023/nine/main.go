package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/extintor/advent_of_code/shared/utils"
)

func solveTwo(line []int) int {
  if line[len(line) - 1] == 0 {
    return 0
  }
  next := make([]int, 0)
  for idx := 0; idx < len(line) - 1; idx++ {
    next = append(next, line[idx+1] - line[idx])
  }
  return line[0] - solveTwo(next)
}

func solveOne(line []int) int {
  if line[len(line) - 1] == 0 {
    return 0
  }
  next := make([]int, 0)
  for idx := 0; idx < len(line) - 1; idx++ {
    next = append(next, line[idx+1] - line[idx])
  }
  return line[len(line) - 1] + solveOne(next)
}

func main() {
  input := utils.ReadInput()

  sum := 0
  sum2 := 0
  for _, line := range input {
    ns := strings.Fields(line)
    l := make([]int, 0)
    for _, n := range ns {
      num, _ := strconv.Atoi(n)
      l = append(l, num)
    }
    sum += solveOne(l)
    sum2 += solveTwo(l)
  }
  fmt.Println(sum, sum2)
}
