package main

import (
	"bufio"
	"os"
	"strconv"
)

func part_one() int {
	expenses := make(map[int]int)

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		new_expense, _ := strconv.Atoi(scanner.Text())
		if expenses[2020-new_expense] != 0 {
			return new_expense * expenses[2020-new_expense]
		}
		expenses[new_expense] = new_expense
	}
	return 0
}
