package main

import (
  "container/heap"
	"fmt"
	"math"
	"strconv"

	"github.com/obrahc/advent_of_code/shared/utils")

type MinHeap []NodeDistance

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(NodeDistance))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type Node struct {
  row int
  column int
  direction int
  incrementDirection int
}

type NodeDistance struct {
  Node
  distance int
}

func solve(input []string, condition func(NodeDistance, NodeDistance) bool) int {
  numRows := len(input)
  numCols := len(input[0])
  convertedInput := make([][]int, 0)
  for _, r := range input {
    t := make([]int, 0)
    for _, c := range r {
      v, _ := strconv.Atoi(string(c))
      t = append(t, v)
    }
    convertedInput = append(convertedInput, t)
  }
  minheap := &MinHeap{}
  heap.Init(minheap)
  
  visited := make(map[Node]int)
  heap.Push(minheap, NodeDistance{Node{0, 0, -1, -1}, 0})
  for minheap.Len() !=0 {
    node := heap.Pop(minheap).(NodeDistance)
    if _, ok := visited[node.Node]; ok {
        continue
    }
    visited[node.Node] = node.distance
    for i, coord := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
      newRow := node.row + coord[0]
      newCol := node.column + coord[1]
      newDir := i
      var newInDir int
      if newDir != node.direction {
        newInDir = 1
      } else {
        newInDir = node.incrementDirection + 1
      }

      isntReverse := (newDir + 2) % 4 != node.direction

      if 0 <= newRow && newRow < numRows && 0 <= newCol && newCol < numCols {
        cost := convertedInput[newRow][newCol]
        newDistance := node.distance + cost
        newNode := NodeDistance{Node{newRow, newCol, newDir, newInDir}, newDistance} 
        if condition(node, newNode) {
          if isntReverse {
            heap.Push(minheap, newNode)
          }
        }
      }
    }
  }

  min := math.MaxInt
  for n, d := range visited {
    if n.row == numRows - 1  && n.column == numCols - 1 {
      if d < min {
        min = d
      }
    }
  }
  return min
}

func solveOne(input []string) int {
  return solve(input, func(_ NodeDistance, nn NodeDistance) bool {if nn.incrementDirection > 3 { return false} else {return true} })
}

func solveTwo(input []string) int {
  return solve(input, func(cn NodeDistance, nn NodeDistance) bool {if nn.incrementDirection<=10 && (cn.direction==nn.direction || cn.incrementDirection>=4 || cn.incrementDirection==-1) { return true} else {return false} })
}

func main() {
  input := utils.ReadInput()
  fmt.Println(solveOne(input))
  fmt.Println(solveTwo(input))
}
