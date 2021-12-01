package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(input []int) int {
	var increases, lastDepth = -1, -1
	for _, currentDepth := range input {
		if currentDepth > lastDepth {
			increases += 1
		}
		lastDepth = currentDepth
	}
	return increases
}

func main() {

	var input []int

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputInt, err := strconv.Atoi(scanner.Text())
		if err != nil{
			return
		}
		input = append(input, inputInt)
	}

	fmt.Println("Result:", solve(input))
}
