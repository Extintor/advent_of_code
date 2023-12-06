package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/extintor/advent_of_code/shared/utils"
)

func solve(races [][2]int) int {
  total := 1
	for _, race := range races {
	  sum := 0
		time := race[0]
		distance := race[1]
		for speed := 0; speed <= time; speed ++ {
			remaining := time - speed
			traveled := remaining * speed
			if traveled > distance {
				sum += 1
			}
		}
    total *= sum
	}
  return total
}

func parseRace(ts, ds string) ([2]int, error) {
  t, err := strconv.Atoi(ts)
  if err != nil {
    return [2]int{}, err
  }
  d, err := strconv.Atoi(ds)
  if err != nil {
    return [2]int{}, err
  }
  return [2]int{t, d}, nil
}

func solveTwo(timeSlice, distanceSlice []string) (int, error) {
  r, err := parseRace(strings.Join(timeSlice, ""),strings.Join(distanceSlice, ""))
  if err != nil {
    return 0, err
  }
  races := [][2]int{r}
  return solve(races), nil
}

func solveOne(timeSlice, distanceSlice []string) (int, error) {
  races := make([][2]int, 0)
  for i := 0; i < len(timeSlice); i++ {
    r, err := parseRace(timeSlice[i], distanceSlice[i])
    if err != nil {
      return 0, err
    }
    races = append(races, r)
  }
  return solve(races), nil
}

func main() {
  lines := utils.ReadInput()
  timeSlice := strings.Fields(strings.TrimPrefix(lines[0], "Time:"))
  distanceSlice := strings.Fields(strings.TrimPrefix(lines[1], "Distance:"))
	fmt.Println(solveOne(timeSlice, distanceSlice))
	fmt.Println(solveTwo(timeSlice, distanceSlice))
}
