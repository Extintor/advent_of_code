package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/obrahc/advent_of_code/shared/utils"
)

type rel struct {
  label string
  lens int
}

type hashmap struct {
  boxes map[int][]rel
}

func newHashMap() *hashmap {
  boxes := make(map[int][]rel)
  return &hashmap{boxes}
}


func (hm *hashmap) add(index int, r rel ) {
  var box []rel
  var ok bool
  if box, ok = hm.boxes[index]; !ok {
    box = make([]rel, 0)
  }
  found := false
  for i, c := range box {
    if c.label == r.label {
      box[i] = r
      found = true
      break
    }
  }

  if !found {
    box = append(box, r)
  }
  hm.boxes[index] = box
}

func (hm *hashmap) remove(index int, label string ) {
  box := hm.boxes[index]
  found := false
  foundIdx := 0
  for i, c := range box {
    if c.label == label {
      found = true
      foundIdx = i
    }
  }
  if found {
    box = append(box[:foundIdx], box[foundIdx+1:]...)
  }
  hm.boxes[index] = box
}

func hash(seq string) int {
  value := 0
  for _, c := range seq {
    value +=int(c)
    value *= 17
    value = value % 256
  }
  return value
}

func solveTwo(){
  input := utils.ReadInput()
  line := input[0]
  hm := newHashMap()
  for _, seq := range strings.Split(line, ",") {
    if seq[len(seq) - 1] == '-' {
      label := seq[:len(seq) - 1]
      box := hash(label)
      fmt.Println(seq, box, label)
      hm.remove(box, label)
    } else {
      label := seq[:len(seq) - 2]
      box := hash(label)
      lens, _ := strconv.Atoi(string(seq[len(seq) - 1]))
      fmt.Println(seq, box, label, lens)
      hm.add(box, rel{label, lens})
    }
  }
  fmt.Println(hm.boxes)

  sum := 0
  for k, v := range hm.boxes {
    for i, b := range v {
      value := k + 1
      value *= i + 1
      value *= b.lens
      fmt.Println(b.label, value)
      sum += value
    }
  }
  fmt.Println(sum)
}

func main() {
  input := utils.ReadInput()
  line := input[0]
  sum := 0
  for _, seq := range strings.Split(line, ",") {
    sum += hash(seq)
  }
    fmt.Println(sum)
  solveTwo()
}
