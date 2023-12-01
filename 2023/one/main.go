package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type value string

func (v value) replace(o, n string) value {
	return value(strings.ReplaceAll(string(v), o, n))
}

func (v value) toString() string {
	return string(v)
}

func solveOne(input []string) (int, error) {
	sum := 0
	for _, v := range input {
		r, _ := regexp.Compile("(\\d)")
		m := r.FindAllString(v, -1)
		n, err := strconv.Atoi(m[0] + m[len(m)-1])
		if err != nil {
			return 0, err
		}
		sum += n
	}
	return sum, nil
}

func solveTwo(input []string) (int, error) {
	modifiedInput := make([]string, 0)
	for _, v := range input {
		s := value(v).
			replace("one", "o1e").
			replace("two", "t2o").
			replace("three", "t3e").
			replace("four", "f4r").
			replace("five", "f5e").
			replace("six", "s6x").
			replace("seven", "s7n").
			replace("eight", "e8t").
			replace("nine", "n9e").toString()
		modifiedInput = append(modifiedInput, s)
	}
	return solveOne(modifiedInput)
}

func main() {
	input := make([]string, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		input = append(input, t)
	}
	fmt.Println(solveOne(input))
	fmt.Println(solveTwo(input))

}
