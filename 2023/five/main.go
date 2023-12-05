package main

import (
	"fmt"
	"strings"

	"github.com/extintor/advent_of_code/shared/utils"
)


type mRange struct {
  start int
  end   int
  add   int
}

func (m mRange) in(n int) bool {
  return (m.start <= n) && (n < m.end)
}

func (m mRange) transform(n int) int {
  if m.in(n) {
    return n + m.add
  }
  return n
}

func solveSeed(maps [][]mRange, seed int) int {
  currentSeed := seed
  for _, l := range maps {
    newValue := currentSeed
    for _, m := range l {
      n := m.transform(currentSeed)
      if n != newValue {
        newValue = n
        break
      }
    }
    currentSeed = newValue 
  }
  return currentSeed
}

func solveOne(maps [][]mRange, seeds []int) int {
  m := 0
  for i, s := range seeds {
    v := solveSeed(maps, s)
    if i == 0 || v < m {
      m = v
    }
  }
  return m
}

func worker(id int, maps [][]mRange, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- solveSeed(maps, j) 
    }
}

func solveTwo(maps [][]mRange, seeds []int) int {
  const numWorkers = 16
  jobs := make(chan int, numWorkers)
  results := make(chan int, numWorkers)
  for w := 1; w <= numWorkers; w++ {
    go worker(w, maps, jobs, results)
  }
  m := 1000_000_000_000
  go func() {
    for v := range results {
      if v < m {
        m = v
      }
    }
  }()
  for i := 0; i < len(seeds); i += 2 {
    for j := 0; j < seeds[i + 1]; j++ {
      s := seeds[i] + j
      jobs <- s
    }
  }
  for len(results) != 0 {

  }
  close(results)
  return m
}

func main() {
  input := utils.ReadInput()

  maps := make([][]mRange, 0)
  var m []mRange
  var s []int
  var err error

  for i, line := range input {
    if i == 0 {
      l := strings.TrimPrefix(line, "seeds: ") 
      s, err = utils.StringSliceToIntSlice(strings.Fields(l))
      if err != nil {
        panic(err)
      }
      continue
    }

    if strings.HasSuffix(line, "map:") {
      if len(m) != 0 {
        maps = append(maps, m)
      }
      m = make([]mRange, 0)
    } else if line != "" {
      numbers, err := utils.StringSliceToIntSlice(strings.Fields(line))
      if err != nil {
        panic(err)
      }
      m = append(m, mRange{numbers[1], numbers[1] + numbers[2], numbers[0] - numbers[1]})
    }
  }

  if len(m) != 0 {
    maps = append(maps, m)
  }

  fmt.Println(solveOne(maps, s))
  fmt.Println(solveTwo(maps, s))
}
