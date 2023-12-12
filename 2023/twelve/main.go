package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/extintor/advent_of_code/shared/utils"
)

type key struct {
	s          string
	inGroup    bool
	groupCount int
	remain     string
}

func newKey(s string, inGroup bool, groupCount int, remain []int) key {
	r := ""
	for _, v := range remain {
		r += "/" + strconv.Itoa(v)
	}
	return key{s, inGroup, groupCount, r}
}

var cache map[key]int = make(map[key]int)

func solve(s string, inGroup bool, groupCount int, remain []int) int {
	cacheKey := newKey(s, inGroup, groupCount, remain)
	if v, ok := cache[cacheKey]; ok {
		return v
	}

	// If len is 0 we only have 2 possible valid configurations
	if len(s) == 0 {
		//We just closed the group with the last characted
		if len(remain) == 1 && inGroup && groupCount == remain[0] {
			return 1
		}

		//nothing else to process and not in group
		if !inGroup && len(remain) == 0 {
			return 1
		}

		// Everything else is invalid
		return 0
	}

	if inGroup && len(remain) == 0 {
		return 0
	}

	switch s[0] {
	case '#':
		if inGroup {
			// Char is inside the group
			d := solve(s[1:], true, groupCount+1, remain)
			cache[cacheKey] = d
			return d
		}
		// Starting a Group
		d := solve(s[1:], true, 1, remain)
		cache[cacheKey] = d
		return d
	case '?':
		if inGroup {
			if groupCount == remain[0] {
				// Group Closed
				d := solve(s[1:], false, 0, remain[1:])
				cache[cacheKey] = d
				return d
			}
			// Char is inside the group
			d := solve(s[1:], true, groupCount+1, remain)
			cache[cacheKey] = d
			return d
		}
		// We need to sum the cases were is a start of the group and the cases were it isn't
		d := solve(s[1:], true, 1, remain) + solve(s[1:], false, 0, remain)
		cache[cacheKey] = d
		return d
	case '.':
		if inGroup {
			if groupCount != remain[0] {
				// Invalid Config
				return 0
			}

			// Group closed
			d := solve(s[1:], false, 0, remain[1:])
			cache[cacheKey] = d
			return d
		}

		// Not in group so we can continue
		d := solve(s[1:], false, 0, remain)
		cache[cacheKey] = d
		return d
	}

	return 0
}

func solveOne(input []string) int {
	s := 0
	for _, l := range input {
		line := strings.Split(l, " ")
		springs := line[0]
		damaged, _ := utils.StringSliceToIntSlice(strings.Split(line[1], ","))
		s += solve(springs, false, 0, damaged)
	}
	return s
}

func solveTwo(input []string) int {
	s := 0
	for _, l := range input {
		springs := make([]string, 0)
		damaged := make([]int, 0)
		for i := 0; i < 5; i++ {
			line := strings.Split(l, " ")
			springs = append(springs, line[0])
			d, _ := utils.StringSliceToIntSlice(strings.Split(line[1], ","))
			damaged = append(damaged, d...)
		}
		s += solve(strings.Join(springs, "?"), false, 0, damaged)
	}
	return s
}

func main() {
	fmt.Println(solveOne(utils.ReadInput()))
	fmt.Println(solveTwo(utils.ReadInput()))
}
