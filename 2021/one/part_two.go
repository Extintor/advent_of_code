package one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func addWindow(window []int) int {
	result := 0
	for _, num := range window {
		result += num
	}
	return result
}

func solve(input []int) int {
	var increases, sum, lastSum = -1, -1, -1

	for i := 0; i <= len(input)-3; i++ {
		window := input[i : i+3]
		sum = addWindow(window)
		if sum > lastSum {
			increases += 1
		}
		lastSum = sum
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
		if err != nil {
			return
		}
		input = append(input, inputInt)
	}

	fmt.Println("Result:", solve(input))
}
