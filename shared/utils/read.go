package utils

import (
	"bufio"
	"os"
)

func ReadInput() []string {
  input := make([]string, 0)
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    input = append(input, scanner.Text())
  }

  return input
}
