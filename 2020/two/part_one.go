package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part_one() int {
	result := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	r := regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		clean := r.FindAllStringSubmatch(scanner.Text(), -1)[0][1:]
		lower_bound, _ := strconv.Atoi(clean[0])
		higher_bound, _ := strconv.Atoi(clean[1])
		if lower_bound <= strings.Count(clean[3], clean[2]) && strings.Count(clean[3], clean[2]) <= higher_bound {
			result += 1
		}
	}
	return result
}
