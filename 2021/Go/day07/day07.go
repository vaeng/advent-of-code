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
	lines := util.GetLines("day07\\07.in")
	start := time.Now()
	day07a(lines)
	duration := time.Since(start)
	day07b(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func day07a(lines []string) {
	heights := map[int]int{}
	total := 0
	min := 99999999999
	max := 0
	for _, v := range strings.Split(lines[0], ",") {
		i, _ := strconv.Atoi(v)
		heights[i]++
		total += i
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	mostCommon := 0
	highestCounter := 0
	for k, v := range heights {
		if v > highestCounter {
			mostCommon = k
			highestCounter = v
		}
	}
	fmt.Println(mostCommon)
	minfuel := total

	for i := min; i < max; i++ {
		fuel := 0
		for k, v := range heights {
			if i < k {
				fuel += v * (k - i)
			} else {
				fuel += v * (i - k)
			}
		}
		if fuel < minfuel {
			minfuel = fuel
		}
	}

	fmt.Printf("Solution for part A: %d\n", minfuel)
}
func sliceSum(slice []*big.Int) *big.Int {
	total := big.NewInt(0)
	for _, v := range slice {
		total.Add(total, v)
	}
	return total
}

func fillFuelCost(max int) []*big.Int {
	fuelcost := make([]*big.Int, max+1)
	fuelcost[0] = big.NewInt(0)
	fuelcost[1] = big.NewInt(1)
	for i := 2; i < max+1; i++ {
		fuelcost[i] = big.NewInt(0).Add(fuelcost[i-1], big.NewInt(int64(i)))
	}
	return fuelcost
}

func day07b(lines []string) {
	heights := map[int]int{}
	total := 0
	min := 99999999999
	max := 0
	for _, v := range strings.Split(lines[0], ",") {
		i, _ := strconv.Atoi(v)
		heights[i]++
		total += i
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	fuelcost := fillFuelCost(max - min)
	minfuel := big.NewInt(0)
	minfuel.Mul(fuelcost[max-min], big.NewInt(int64(total)))
	chosenPos := 0

	for i := min; i < max; i++ {
		fuel := big.NewInt(0)
		for k, v := range heights {
			if i < k {
				fuel.Add(fuel, big.NewInt(0).Mul(fuelcost[k-i], big.NewInt(int64(v))))
			} else {
				fuel.Add(fuel, big.NewInt(0).Mul(fuelcost[i-k], big.NewInt(int64(v))))
			}
			if fuel.Cmp(minfuel) == 1 {
				break
			}
		}
		if fuel.Cmp(minfuel) == -1 {
			minfuel = fuel
			chosenPos = i
		}
	}

	fmt.Printf("Solution for part B: Fuel: %d, Position: %d\n", minfuel, chosenPos)
}
