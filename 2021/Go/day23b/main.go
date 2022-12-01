package main

import (
	"aoc2021/util"
	"fmt"
	"time"
)

func main() {
	lines := util.GetLines(".\\day23b\\example")
	start := time.Now()
	partA(lines)
	duration := time.Since(start)
	partB(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func partA(lines []string) {

	fmt.Printf("Solution for part A: %v\n", 0)
}

func partB(lines []string) {

	fmt.Printf("Solution for part B: %v\n", 0)
}
