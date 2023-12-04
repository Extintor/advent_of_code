package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/extintor/advent_of_code/shared/generics"
	"github.com/extintor/advent_of_code/shared/utils"
)

type card struct {
  id int
  winNum map[string]struct{}
  drawNum []string
}

func getCardsWins(cards []card) map[int]int {
  m := make(map[int]int)
  for _, card := range cards {
    count := 0
    for _, num := range card.drawNum {
      if _, ok := card.winNum[num]; ok == true {
        count++
      }
    }
    m[card.id] = count
  }
  return m
}

func solveOne(cards []card) int {
  sum := 0
  for _, v := range getCardsWins(cards) {
    sum += int(math.Pow(2, float64(v - 1)))
  }
  return sum
}

func solveTwo(cards []card) int {
  q := generics.NewQueue[int]()
  sum := 0
  wins := getCardsWins(cards)
  for _, card := range cards {
    q.Enqueue(card.id)
  }
  for !q.IsEmpty() { 
    id, _ := q.Dequeue()
    sum++
    count := wins[id]
    for idx := id + 1; idx <= id + count; idx++ {
      q.Enqueue(idx)
    }
  }
  return sum
}

func main() {
  input := utils.ReadInput()

	re := regexp.MustCompile(`Card\s+(\d+):\s+([\d\s]+)\|\s+([\d\s]+)`)

  cards := make([]card, 0)

  for _, v := range input {
    matches := re.FindStringSubmatch(v)
    id, err := strconv.Atoi(matches[1])
    if err != nil {
      panic(err)
    }
    cards = append(cards, card{id, utils.StringSliceToSet(strings.Fields(matches[2])), strings.Fields(matches[3])})
  }
  fmt.Println(solveOne(cards))
  fmt.Println(solveTwo(cards))
}
