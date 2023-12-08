package main

import (
	"fmt"
	"strings"
  
  "github.com/extintor/advent_of_code/shared/utils"
)

type Node struct {
	Value string
	Left  string
	Right string
}

func GCD(a, b int) int {
  for b != 0 {
    t := b
    b = a % b
    a = t
  }
  return a
}

func LCM(integers ...int) int {
  result := integers[0] * integers[1] / GCD(integers[0], integers[1])

  for i := 2; i < len(integers); i++ {
    result = LCM(result, integers[i])
  }

  return result
}

func main() {
	lines := utils.ReadInput()

	// Decode the first line into a list of characters
	charList := []rune(lines[0])
	fmt.Println("List of Characters:", string(charList))

	// Parse the tree structure
	nodes := make(map[string]*Node)
	for _, line := range lines[1:] {
    if line == "" { continue}
		parts := strings.Split(line, " = ")
		value := parts[0]
		children := strings.Trim(parts[1], "()")
		childParts := strings.Split(children, ", ")

		node, exists := nodes[value]
		if !exists {
			node = &Node{Value: value, Left: childParts[0], Right: childParts[1]}
			nodes[value] = node
		}
	}

  nextNode := "AAA"
  count := 0
  found := false
  for found == false {
    for _, move := range charList {
      currentNode := nodes[nextNode]
      if currentNode.Value == "ZZZ" {
        found = true
        break
      }
      if move == 'L' {
        nextNode = currentNode.Left
      } else {
        nextNode = currentNode.Right
      }
      count ++
    }
  }
  fmt.Println(count)

  startNode := make([]string, 0) 
  for move := range nodes {
    if move[len(move) - 1] == 'A' {
      startNode = append(startNode, move)
  }
  }

  result := make([]int, 0)
  
  for _, node := range startNode {
    nextNode := node 
    count := 0
    found := false
    for found == false {
      for _, move := range charList {
        currentNode := nodes[nextNode]
        if currentNode.Value[len(currentNode.Value) - 1] == 'Z' {
          found = true
          break
        }
        if move == 'L' {
          nextNode = currentNode.Left
        } else {
          nextNode = currentNode.Right
        }
        count ++
      }
    }
    result = append(result, count)
  }

  fmt.Println(LCM(result...))
}
