package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(input []uint64) (uint64, error) {
	result := uint64(0)
	for i:=0; i<12; i++ {
		count := uint64(0)
		mask := uint64(1 << (12 - i -1))
		for _, inputLine := range input {
			count += inputLine & mask
		}
		result += count / uint64(500) & mask
	}
	invertedResult := result ^ ((1 << 12) - 1)
	return result*invertedResult, nil
}

func main() {
	input := make([]uint64, 0, 0)

	file, err := os.Open("input.txt")
	if err != nil{
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine, err := strconv.ParseUint(scanner.Text(), 2, 12)
		if err != nil{
			return
		}
		input = append(input, inputLine)
	}
	result, err := solve(input)
	if err != nil{
		return
	}
	fmt.Println("Result:", result)
}
