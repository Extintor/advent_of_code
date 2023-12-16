package main

import ( "fmt"
	"strings"

	"github.com/obrahc/advent_of_code/shared/utils"
)


func rotate90(m [][]rune) [][]rune {
	if len(m) == 0 {
		return [][]rune{} }

	rowCount := len(m)
	colCount := len(m[0])

	newMatrix := make([][]rune, colCount)
	for i := range newMatrix {
		newMatrix[i] = make([]rune, rowCount)
	}

	for i := range m {
		for j := range m[i] {
			newMatrix[j][rowCount-1-i] = m[i][j]
		}
	}

	return newMatrix
}

func newArray(n int) [][]rune {
  a := make([][]rune, n)
  for i := range a {
    if len(a[i]) == 0 {
      a[i] = make([]rune, n)
    }
    for j, c := range strings.Repeat(".", n) {
      a[i][j] = c
    }
  }
  return a
}

func solveOne(input []string) {
  maxLen := len(input)
  sum := 0
  fixedRocks := make(map[int]int)
  for i, row := range input {
    for j, rock := range row {
      lastFixed, ok := fixedRocks[j]
      if !ok {
        fixedRocks[j] = -1
        lastFixed = -1
      }
      switch rock {
      case '.':
        continue
      case '#':
        fixedRocks[j] = i 
      case 'O':
        restingPosition := (lastFixed + 1)
        sum += (maxLen - restingPosition)
        fixedRocks[j] = restingPosition 
    }
    }
  }
  fmt.Println(sum)
}

func hash(m [][]rune) string {
  r := ""
  for _, a := range m {
    for _, b := range a {
      r += string(b)
    }
  }
  return r
}

func solveTwo(input []string) {
  config := convertInput(input)
  maxLen := len(config)
  cache := make(map[string]int)
  
  cicles := 0 
  totalCicles := 1000000000
  period := 0

  newConfig := newArray(maxLen)
  for cicle := 0; cicle < totalCicles; cicle ++ {
    cicles += 1
    for rotation := 0; rotation < 4; rotation++ { 
      fixedRocks := make(map[int]int)
      newConfig = newArray(maxLen)
      for i, row := range config {
        for j, rock := range row {
          lastFixed, ok := fixedRocks[j]
          if !ok {
            fixedRocks[j] = -1
            lastFixed = -1
          }
          switch rock {
          case '.':
            continue
          case '#':
            fixedRocks[j] = i 
            newConfig[i][j] = '#'
          case 'O':
            restingPosition := (lastFixed + 1)
            fixedRocks[j] = restingPosition 
            newConfig[restingPosition][j] = 'O' 
        }
        }
      }
    config = rotate90(newConfig)
    }
    cacheKey := hash(config)
    if c, ok := cache[cacheKey]; ok  {
      period = cicle - c
      break
    }
    cache[cacheKey] = cicle 
  }
  ciclesLeft := (totalCicles - cicles) % period
  
  for cicle := 0; cicle < ciclesLeft; cicle ++ {
    for rotation := 0; rotation < 4; rotation++ { 
      fixedRocks := make(map[int]int)
      newConfig = newArray(maxLen)
      for i, row := range config {
        for j, rock := range row {
          lastFixed, ok := fixedRocks[j]
          if !ok {
            fixedRocks[j] = -1
            lastFixed = -1
          }
          switch rock {
          case '.':
            continue
          case '#':
            fixedRocks[j] = i 
            newConfig[i][j] = '#'
          case 'O':
            restingPosition := (lastFixed + 1)
            fixedRocks[j] = restingPosition 
            newConfig[restingPosition][j] = 'O' 
            
        }
        }
      }
    config = rotate90(newConfig)
    }
  }
  sum := 0
  for i, r := range config {
    for _, c := range r {
      if c == 'O' {
        sum += maxLen - i
      }
    }
  }
  fmt.Println(sum)
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

func main() {
  input := utils.ReadInput()
  solveOne(input)
  solveTwo(input)
}
