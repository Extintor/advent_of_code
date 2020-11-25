package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var values []float64
	var value int
	for scanner.Scan() {
		value, _ = strconv.Atoi(scanner.Text())
		values = append(values, float64(value))
	}

	file.Close()
	var total_fuel int64 = 0
	var module_fuel int64
	for _, module_mass := range values {
		module_fuel = int64(module_mass / 3) - 2
		for module_fuel > 0 {
			total_fuel += module_fuel
			module_fuel = (module_fuel / 3) - 2
		}
	}
	fmt.Println(total_fuel)
}

