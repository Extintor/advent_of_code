package main

import "testing"

func TestOne(t *testing.T) {

	input := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	have, err := solveOne(input)
	if err != nil {
		t.Fatal(err)
	}
	want := 142

	if want != have {
		t.Fatalf("Want %d but have %d", want, have)
	}

}

func TestTwo(t *testing.T) {

	input := []string{"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen"}

	want, err := solveTwo(input)
	if err != nil {
		t.Fatal(err)
	}
	have := 281

	if have != want {
		t.Fatalf("Want %d but have %d", want, have)
	}

}
