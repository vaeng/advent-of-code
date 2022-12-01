package main

import (
	"aoc2021/util"
	"fmt"
	_ "strconv"
	_ "strings"
	"time"
)

func main() {
	lines := util.GetLines("day14\\14.in")
	start := time.Now()
	day14a(lines)
	duration := time.Since(start)
	day14b(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func day14a(lines []string) {

	seed := lines[0]
	// set up initial counter
	counter := map[string]int{}
	for _, r := range seed {
		counter[string(r)]++
	}

	// set up initial parts list
	parts := map[string]int{}
	for i := 0; i < len(seed)-1; i++ {
		parts[seed[i:i+2]]++
	}

	// set up rules
	rules := map[string]string{}
	for _, rule := range lines[2:] {
		rules[rule[0:2]] = rule[6:7]
	}

	//run algorithm
	goal := 10
	for i := 0; i < goal; i++ {
		// make new empty parts map
		newParts := map[string]int{}
		// use old part map as iteration base
		for part, count := range parts {
			if rules[part] != "" {
				insertion := rules[part]
				counter[insertion] += count
				leftPair := string(part[0]) + insertion
				rightPair := insertion + string(part[1])
				newParts[leftPair] += count
				newParts[rightPair] += count
			} else {
				newParts[part] += count
			}
		}
		parts = newParts
	}
	// check max and min
	min := 999999999999999999
	max := -1
	for _, c := range counter {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	fmt.Printf("Solution for part A: %v\n", max-min)
}

func day14b(lines []string) {
	seed := lines[0]
	// set up initial counter
	counter := map[string]uint64{}
	for _, r := range seed {
		counter[string(r)]++
	}

	// set up initial parts list
	parts := map[string]uint64{}
	for i := 0; i < len(seed)-1; i++ {
		parts[seed[i:i+2]]++
	}

	// set up rules
	rules := map[string]string{}
	for _, rule := range lines[2:] {
		rules[rule[0:2]] = rule[6:7]
	}

	//run algorithm
	goal := 400
	for i := 0; i < goal; i++ {
		// make new empty parts map
		newParts := map[string]uint64{}
		// use old part map as iteration base
		for part, count := range parts {
			if rules[part] != "" {
				insertion := rules[part]
				counter[insertion] += count
				leftPair := string(part[0]) + insertion
				rightPair := insertion + string(part[1])
				newParts[leftPair] += count
				newParts[rightPair] += count
			} else {
				newParts[part] += count
			}
		}
		parts = newParts
	}
	// check max and min
	min := uint64(9999999999999999999)
	max := uint64(0)
	for _, c := range counter {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	fmt.Printf("Solution for part B: %v\n", max-min)
}
