package main

import (
	"fmt"

	"github.com/obrahc/advent_of_code/shared/utils"
)

func verticalReflection(pattern []string, original int) int {
    width := len(pattern[0])
    for col := 0; col < width-1; col++ {
        isMirror := true
        for i := 0; i <= col; i++ {
            if col+i+1 >= width {
                break
            }
            for row := range pattern {
                if pattern[row][col-i] != pattern[row][col+i+1] {
                    isMirror = false
                    break
                }
            }
            if !isMirror {
                break
            }
        }
        if isMirror {
          if col + 1 != original {
            return col + 1 
          }
        }
    }
    return -1
}

func horizontalReflection(pattern []string, original int) int {
    height := len(pattern)
    for row := 0; row < height-1; row++ {
        isMirror := true
        for i := 0; i <= row; i++ {
            if row+i+1 >= height {
                break
            }
            if pattern[row-i] != pattern[row+i+1] {
                isMirror = false
                break
            }
        }
        if isMirror {
          if row + 1 != original {
            return row + 1
          }
        }
    }
    return -1
}

func flipCharacter(c rune) rune {
    if c == '#' {
        return '.'
    }
    return '#'
}

func generateFlippedPatterns(pattern []string) [][]string {
    var flippedPatterns [][]string
    for i, row := range pattern {
        for j, char := range row {
            newPattern := make([]string, len(pattern))
            copy(newPattern, pattern)

            newRow := []rune(newPattern[i])
            newRow[j] = flipCharacter(rune(char))
            newPattern[i] = string(newRow)

            flippedPatterns = append(flippedPatterns, newPattern)
        }
    }
    return flippedPatterns
}

func solveOne(patterns [][]string) int {
  total := 0
  for _, pattern := range patterns {
    vReflection := verticalReflection(pattern, -2)
    hReflection := horizontalReflection(pattern, -2)

    if vReflection != -1 {
      total += vReflection
      continue
    }
    if hReflection != -1 {
      total += 100 * hReflection
      continue
    }
  }
  return total
}

func solveTwo(patterns [][]string) int {
  total := 0
  for _, pattern := range patterns {
    originalVReflection := verticalReflection(pattern, -2)
    originalHReflection := horizontalReflection(pattern, -2)
    flippedPatterns := generateFlippedPatterns(pattern)
    for _, flippedPattern := range flippedPatterns {
      vReflection := verticalReflection(flippedPattern, originalVReflection)
      hReflection := horizontalReflection(flippedPattern, originalHReflection)

      isNewVReflection := vReflection != -1 && vReflection != originalVReflection
      isNewHReflection := hReflection != -1 &&  hReflection != originalHReflection

      if isNewVReflection {
        total += vReflection
        break
      } else if isNewHReflection {
        total += 100 * hReflection
        break
      }
    }
  }
  return total
}

func main() {
  var patterns [][]string
  var currentPattern []string

  for _, line := range utils.ReadInput() {
    if line == "" {
      patterns = append(patterns, currentPattern)
      currentPattern = nil
    } else {
      currentPattern = append(currentPattern, line)
    }
  }
  patterns = append(patterns, currentPattern)

  fmt.Println(solveOne(patterns))
  fmt.Println(solveTwo(patterns))
}

