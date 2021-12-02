package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(input []string) (int, error) {
	horizontal, vertical, aim := 0, 0, 0
	for _, inputLine := range input {
		splitted := strings.Split(inputLine, " ")
		k := splitted[0]
		v, err := strconv.Atoi(splitted[1])
		if err != nil{
			return -1, err
		}

		switch direction := k; direction {
		case "forward":
			horizontal += v
			vertical += aim * v
		case "up":
			aim -= v
		case "down":
			aim += v
		}
	}
	return horizontal * vertical, nil
}

func main() {

	input := make([]string, 0, 0)

	file, err := os.Open("input.txt")
	if err != nil{
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		input = append(input, inputLine)
	}
	result, err := solve(input)
	if err != nil{
		return
	}
	fmt.Println("Result:", result)
}
