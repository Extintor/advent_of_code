package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(input []uint64) (uint64, error) {
	mask := uint64(0)
	newMask := uint64(0)
	setSize := len(input)
	co2Value := uint64(0)
	o2Value := uint64(0)
	for pos := 0; pos < 12; pos++ {
		newMask = mask + uint64(1)
		result := make([]uint64, 0)
		for _, inputLine := range input {
			if inputLine>>(12-pos-1) == mask && inputLine>>(12-pos-1) != newMask {
				result = append(result, inputLine)
			}
		}
		if (len(result) == 1) && setSize == 2 {
			co2Value = result[0]
		}
		if len(result)<<1 >= setSize {
			setSize = setSize - len(result)
			mask = newMask
		} else {
			setSize = len(result)
		}
		mask = mask << 1
	}

	mask = uint64(0)
	setSize = len(input)
	for pos := 0; pos < 12; pos++ {
		mask += uint64(1)
		result := make([]uint64, 0)
		for _, inputLine := range input {
			if inputLine>>(12-pos-1) == mask {
				result = append(result, inputLine)
			}
		}
		if (len(result) == 1) && setSize == 2 {
			o2Value = result[0]
		}
		if len(result)<<1 >= setSize {
			setSize = len(result)
		} else {
			mask -= uint64(1)
			setSize = setSize - len(result)
		}
		mask = mask << 1
	}
	return o2Value * co2Value, nil
}

func main() {
	input := make([]uint64, 0, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine, err := strconv.ParseUint(scanner.Text(), 2, 12)
		if err != nil {
			return
		}
		input = append(input, inputLine)
	}
	result, err := solve(input)
	if err != nil {
		return
	}
	fmt.Println("Result:", result)
}
