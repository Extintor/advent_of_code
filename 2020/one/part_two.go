package main

import (
	"bufio"
	"os"
	"strconv"
)

func part_two() int {
	expenses := make(map[int]int)
	processed_expenses := []int{}

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		new_expense, _ := strconv.Atoi(scanner.Text())
		for _, expense := range processed_expenses {
			expenses[new_expense+expense] = new_expense * expense
		}

		if expenses[2020-new_expense] != 0 {
			return new_expense * expenses[2020-new_expense]
		}
		processed_expenses = append(processed_expenses, new_expense)
	}
	return 0
}
