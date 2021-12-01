package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part_one() int {
	var increases, lastDepth, currentDepth = -1, -1, -1

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentDepth, _ = strconv.Atoi(scanner.Text())
		if currentDepth > lastDepth {
			increases += 1
		}
		lastDepth = currentDepth
	}
	return increases
}

func main() {
	fmt.Println("Result part one:", part_one())
}


