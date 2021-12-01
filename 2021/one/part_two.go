package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part_two() int {
	var increases, sum, lastSum = -1, 0, -1
	var depthOne, depthTwo, depthThree int = -1, -1, -1

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depthOne = depthTwo
		depthTwo = depthThree
		depthThree, _ = strconv.Atoi(scanner.Text())
		if (depthOne < 0){
			continue
		}
		sum = depthOne + depthTwo + depthThree
		fmt.Printf("%d %d %d\n", depthOne, depthTwo, depthThree)
		if sum > lastSum{
			increases += 1
		}
		lastSum = sum
	}
	return increases
}

func main() {
	fmt.Println("Result part two:", part_two())
}


