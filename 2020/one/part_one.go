package main

import (
	"bufio"
	"os"
	"strconv"
)

func part_one() int {
	expenses := []int{}

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		new_expense, _ := strconv.Atoi(scanner.Text())
		for _, expense := range expenses {
			if new_expense+expense == 2020 {
				return new_expense * expense
			}
		}
		expenses = append(expenses, new_expense)
	}
	return 0
}
