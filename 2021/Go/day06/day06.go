package main

import (
	"aoc2021/util"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func main() {
	lines := util.GetLines("day06\\06.in")
	start := time.Now()
	day06a(lines)
	duration := time.Since(start)
	day06b(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func popSize(days []*big.Int) (sum *big.Int) {
	sum = big.NewInt(0)
	for _, v := range days {
		sum.Add(sum, v)
	}
	return sum
}

func increaseDay(days []*big.Int) []*big.Int {
	nextday := make([]*big.Int, 9, 9)
	for i := 0; i < 9; i++ {
		nextday[i] = big.NewInt(0)
	}
	for i, v := range days {
		if i == 0 {
			nextday[6] = v
			nextday[8] = v
		} else {
			nextday[i-1].Add(nextday[i-1], v)
		}
	}
	return nextday
}

func day06a(lines []string) {
	days := make([]*big.Int, 9, 9)
	for i := 0; i < 9; i++ {
		days[i] = big.NewInt(0)
	}
	for _, v := range strings.Split(lines[0], ",") {
		i, _ := strconv.Atoi(v)
		days[i].Add(days[i], big.NewInt(1))

	}

	for i := 0; i < 257+1; i++ {
		if i == 1 || i == 18 || i == 80 || i == 256 || i == 9999999 {
			fmt.Printf("Day: %d\tPopulation: %d\n", i, popSize(days))
		}
		days = increaseDay(days)
	}

	//fmt.Printf("Solution for part A: %d\n", days)
}

func day06b(lines []string) {

	fmt.Printf("Solution for part B: %d\n", 3)
}
