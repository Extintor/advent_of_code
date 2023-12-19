package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"github.com/obrahc/advent_of_code/shared/utils"
)

type Condition struct {
	Name     string
	Operator string
	Value    int
  Destination string
}

type Rule struct {
	Conditions []Condition
}

type Range struct {
  data map[string][2]int
}

type Key struct {
  data map[string]int
}

func parseData(input string) (Key, error) {
	var data Key

	re := regexp.MustCompile(`(\w)=(\d+)`)

	matches := re.FindAllStringSubmatch(input, -1)

  data.data = make(map[string]int)

	for _, match := range matches {
		key := match[1]
		valueStr := match[2]

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return Key{}, err
		}

		data.data[key] = value
	}

	return data, nil
}

func parseInput(input string) (string, Rule) {

	re := regexp.MustCompile(`(\w+)\{([^}]+)\}`)

	matches := re.FindAllStringSubmatch(input, -1)

  match := matches[0]
  name := match[1]
 	conditionsStr := match[2]

	conditions := parseConditions(conditionsStr)

	return name, Rule{
			Conditions: conditions,
	  }
}

func parseConditions(conditionsStr string) []Condition {
	conditions := make([]Condition, 0)

	conditionStrs := strings.Split(conditionsStr, ",")

	for _, conditionStr := range conditionStrs {
    if !strings.Contains(conditionStr, ":") {
			condition := Condition{
				Name:     "",
				Operator: "",
				Value:    0,
        Destination: conditionStr,
			}
			conditions = append(conditions, condition)
      continue
    }
    re := regexp.MustCompile(`(\w+)(<|>)(\d+):(\w+)`)

    matches := re.FindAllStringSubmatch(conditionStr, -1)
    match := matches[0]

    value := parseValue(match[3])

    condition := Condition{
      Name:     match[1],
      Operator: match[2],
      Value:    value,
      Destination: match[4],
    }

    conditions = append(conditions, condition)
	}

	return conditions
}

func parseValue(valueStr string) int {
	var value int
	fmt.Sscanf(valueStr, "%d", &value)
	return value
}

func copyRange (r Range) Range {
  nr := make(map[string][2]int)
  for k, v := range r.data {
    nr[k] = [2]int{v[0], v[1]}
  }
  return Range{nr}
}

func sortRange(rules map[string]Rule, currentRule Rule, r Range) int {
  for _, c := range currentRule.Conditions {
    next := ""
    nextRange := copyRange(r)
    nextRangeMultiple := copyRange(r)
    match := false
    multiple := false
    if c.Operator == "" {
      match = true
      next = c.Destination
    }
    if c.Operator == ">" {
      if r.data[c.Name][0] < c.Value && r.data[c.Name][1] > c.Value {
        match = true
        multiple = true
        nextRange.data[c.Name] = [2]int{c.Value + 1, r.data[c.Name][1]}
        nextRangeMultiple.data[c.Name] = [2]int{r.data[c.Name][0], c.Value }
        next = c.Destination 
      }
      if r.data[c.Name][0] > c.Value {
        match = true
        next = c.Destination 
      }
    } else {
      if r.data[c.Name][0] < c.Value && r.data[c.Name][1] > c.Value {
        match = true
        multiple = true
        nextRange.data[c.Name] = [2]int{r.data[c.Name][0], c.Value - 1}
        nextRangeMultiple.data[c.Name] = [2]int{c.Value, r.data[c.Name][1]}
        next = c.Destination 
      }
      if r.data[c.Name][1] < c.Value {
        match = true
        next = c.Destination 
      }
    }
    if match {
      if next == "A" {
        s := 1
        for _, t := range nextRange.data {
          s *= t[1] - t[0] + 1
        }
        if multiple {
          return s + sortRange(rules, currentRule, nextRangeMultiple)
        }
        return s 
      } else if next == "R" {
        if multiple {
          return sortRange(rules, currentRule, nextRangeMultiple)
        }
        return 0
      } else {
        if multiple {
          return sortRange(rules, rules[next], nextRange) + sortRange(rules, currentRule, nextRangeMultiple)
        }
        return sortRange(rules, rules[next], nextRange)
      }
    }
  }
  return 0
}

func sortKey(rules map[string]Rule, k Key) bool {
  currentRule := rules["in"]
  for {
    for _, c := range currentRule.Conditions {
      next := ""
      match := false
      if c.Operator == "" {
        match = true
        next = c.Destination
      }
      if c.Operator == ">" {
        if k.data[c.Name] > c.Value {
          match = true
          next = c.Destination 
        }
      } else {
        if k.data[c.Name] < c.Value {
          match = true
          next = c.Destination 
        }
      }
      if match {
        if next == "A" {
          return true
        }
        if next == "R" {
          return false
        }
        currentRule = rules[next]
        break
      }
    }
  }
}

func main() {
  input := utils.ReadInput()

  rules := make(map[string]Rule)
  keys := make([]Key, 0)

  changed := false
  for _, line := range input {
    if len(line) == 0{
      changed = true
      continue
    }
    if !changed {
      name, rule := parseInput(line)
      rules[name] = rule
      continue
    }
    key, _ := parseData(line)
    keys = append(keys, key)
  }
  sum := 0
  for _, k := range keys {
    if sortKey(rules, k) {
      for _, v := range k.data {
        sum += v
      }
    }
  }
  fmt.Println(sum)
  r := Range{map[string][2]int{"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000}}}
  fmt.Println(sortRange(rules, rules["in"], r))
}
