package main

import (
	"bufio"
	"os"
	"strconv"
)

func part_two() int {
	expenses := []int{}

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		new_expense, _ := strconv.Atoi(scanner.Text())
		for _, first_expense := range expenses {
			for _, second_expense := range expenses {
				if new_expense+first_expense+second_expense == 2020 {
					return new_expense * first_expense * second_expense
				}
			}
		}
		expenses = append(expenses, new_expense)
	}
	return 0
}
