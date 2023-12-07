package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/extintor/advent_of_code/shared/utils"
)

type Strength int

const (
	HighCard Strength = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

type hand struct {
	cards string
	bid   int
}

var ranks string = "23456789TJQKA"
var jokerRanks string = "J23456789TQKA"

func cardRank(card rune, ranks string) int {
	return strings.IndexRune(ranks, card)
}

func cardRankWithJoker(card rune) int {
	ranks := "J23456789TQKA"
	return strings.IndexRune(ranks, card)
}

func maxValue(m map[rune]int) rune {
	max := math.MinInt
	var r rune
	for k, v := range m {
		if v == max {
			if cardRank(k, ranks) > cardRank(r, ranks) {
				r = k
			}
		}
		if v > max {
			max = v
			r = k
		}
	}
	return r
}

func handStrengthFromFrequency(frequency map[rune]int) Strength {
	switch len(frequency) {
	case 1:
		return FiveKind
	case 2:
		for _, count := range frequency {
			if count == 4 || count == 1 {
				return FourKind
			}
			if count == 3 || count == 2 {
				return FullHouse
			}
		}
	case 3:
		threeCount := 0
		for _, count := range frequency {
			if count == 3 {
				threeCount++
			}
		}
		if threeCount == 1 {
			return ThreeKind
		}
		return TwoPair
	case 4:
		return OnePair
	}
	return HighCard

}

func checkHandStrengthJoker(hand string) Strength {
	frequency := make(map[rune]int)
	for _, card := range hand {
		if card != 'J' {
			frequency[card]++
		}
	}

	if len(frequency) == 0 {
		return FiveKind
	}

	frequency[maxValue(frequency)] += strings.Count(hand, "J")
	return handStrengthFromFrequency(frequency)

}

func checkHandStrength(hand string) Strength {
	frequency := make(map[rune]int)
	for _, card := range hand {
		frequency[card]++
	}
	return handStrengthFromFrequency(frequency)
}

func compareHands(hand1, hand2, cardsRanks string) bool {
	for i := range hand1 {
		rank1 := cardRank(rune(hand1[i]), cardsRanks)
		rank2 := cardRank(rune(hand2[i]), cardsRanks)
		if rank1 > rank2 {
			return false
		} else if rank1 < rank2 {
			return true
		}
	}
	return false
}

func solveOne(hands []hand) int {
	sort.Slice(hands, func(i, j int) bool {
		hand1 := hands[i]
		hand2 := hands[j]
		strenght1 := checkHandStrength(hand1.cards)
		strenght2 := checkHandStrength(hand2.cards)
		if strenght1 == strenght2 {
			compareHands(hand1.cards, hand2.cards, ranks)
		}
		return strenght1 < strenght2
	})
	sum := 0
	for i, v := range hands {
		sum += v.bid * (i + 1)
	}
	return sum
}

func solveTwo(hands []hand) int {
	sort.Slice(hands, func(i, j int) bool {
		hand1 := hands[i]
		hand2 := hands[j]
		strenght1 := checkHandStrengthJoker(hand1.cards)
		strenght2 := checkHandStrengthJoker(hand2.cards)

		if strenght1 == strenght2 {
			return compareHands(hand1.cards, hand2.cards, jokerRanks)
		}
		return strenght1 < strenght2
	})
	sum := 0
	for i, v := range hands {
		sum += v.bid * (i + 1)
	}
	return sum
}

func main() {
	input := make([]hand, 0)
	for _, v := range utils.ReadInput() {
		sV := strings.Split(v, " ")
		bid, err := strconv.Atoi(sV[1])
		if err != nil {
			panic(err)
		}
		input = append(input, hand{sV[0], bid})
	}

	fmt.Println(solveOne(input))
	fmt.Println(solveTwo(input))
}
