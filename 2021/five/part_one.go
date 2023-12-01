package five

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type point struct {
	x int
	y int
}

func solveOne(input []uint64) (uint64, error) {
	result := uint64(0)
	for i := 0; i < 12; i++ {
		count := uint64(0)
		mask := uint64(1 << (12 - i - 1))
		for _, inputLine := range input {
			count += inputLine & mask
		}
		result += count / uint64(500) & mask
	}
	invertedResult := result ^ ((1 << 12) - 1)
	return result * invertedResult, nil
}

func mainOne() {
	// input := make([]uint64, 0, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	r, _ := regexp.Compile("(/d),(/d) -> (/d),(/d)")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		points := r.FindAllString(scanner.Text(), -1)
		if err != nil {
			return
		}
		fmt.Println(points)
	}
	/*
		 	result, err := solve(input)
			if err != nil {
				return
			}
			fmt.Println("Result:", result)
	*/
}
