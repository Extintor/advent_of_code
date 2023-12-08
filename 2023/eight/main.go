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

func calculateNode(startNode, charList string, nodes map[string]*Node, checkLast func(*Node) bool) int {
	nextNode := startNode
	count := 0
	found := false
	for found == false {
		for _, move := range charList {
			currentNode := nodes[nextNode]
			if checkLast(currentNode) {
				found = true
				break
			}
			if move == 'L' {
				nextNode = currentNode.Left
			} else {
				nextNode = currentNode.Right
			}
			count++
		}
	}
	return count
}

func solveOne(charList string, nodes map[string]*Node) int {
	return calculateNode("AAA", charList, nodes, func(n *Node) bool {
		if n.Value == "ZZZ" {
			return true
		}
		return false
	},
	)
}

func solveTwo(charList string, nodes map[string]*Node) int {
	startNode := make([]string, 0)
	for move := range nodes {
		if move[len(move)-1] == 'A' {
			startNode = append(startNode, move)
		}
	}

	result := make([]int, 0)

	for _, node := range startNode {
		result = append(result, calculateNode(node, charList, nodes, func(n *Node) bool {
			if n.Value[len(n.Value)-1] == 'Z' {
				return true
			}
			return false
		},
		))
	}

	return LCM(result...)
}

func main() {
	lines := utils.ReadInput()

	charList := lines[0]

	nodes := make(map[string]*Node)
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " = ")
		value := parts[0]
		children := strings.Trim(parts[1], "()")
		childParts := strings.Split(children, ", ")

		node := nodes[value]
		node = &Node{Value: value, Left: childParts[0], Right: childParts[1]}
		nodes[value] = node
	}

	fmt.Println(solveOne(charList, nodes))
	fmt.Println(solveTwo(charList, nodes))
}
