package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func part_two() int {
	result := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	r := regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		clean := r.FindAllStringSubmatch(scanner.Text(), -1)[0][1:]
		first_position, _ := strconv.Atoi(clean[0])
		second_position, _ := strconv.Atoi(clean[1])
		count := 0
		if len(clean[3]) > first_position-1 && len(clean[3]) > second_position-1 {
			if string(clean[3][first_position-1]) == clean[2] {
				count += 1
			}
			if string(clean[3][second_position-1]) == clean[2] {
				count += 1
			}
			if count == 1 {
				result += 1
			}
		}
	}
	return result
}
