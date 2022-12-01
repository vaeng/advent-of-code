package main

import (
	"aoc2021/util"
	"fmt"
	"sort"
	"time"
)

func main() {
	lines := util.GetLines("day10\\10.in")
	start := time.Now()
	day10a(lines)
	duration := time.Since(start)
	day10b(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func checkLine(line string) int {
	stack := make([]rune, len(line))
	stackpointer := 0
	for _, r := range line {
		switch rune(r) {
		case '(':
			stack[stackpointer] = ')'
			stackpointer++
		case '<':
			stack[stackpointer] = '>'
			stackpointer++
		case '{':
			stack[stackpointer] = '}'
			stackpointer++
		case '[':
			stack[stackpointer] = ']'
			stackpointer++
		case ')', '>', '}', ']':
			if stack[stackpointer-1] == rune(r) {
				stackpointer--
			} else {
				return score(r)
			}
		}
	}
	return 0
}

func checkLine2(line string) int {
	stack := make([]rune, len(line))
	stackpointer := 0
	for _, r := range line {
		switch rune(r) {
		case '(':
			stack[stackpointer] = ')'
			stackpointer++
		case '<':
			stack[stackpointer] = '>'
			stackpointer++
		case '{':
			stack[stackpointer] = '}'
			stackpointer++
		case '[':
			stack[stackpointer] = ']'
			stackpointer++
		case ')', '>', '}', ']':
			if stack[stackpointer-1] == rune(r) {
				stackpointer--
			} else {
				return -1
			}
		}
	}
	score := 0
	for i := stackpointer - 1; i >= 0; i-- {
		score *= 5
		switch stack[i] {
		case ')':
			score++
		case '>':
			score += 4
		case '}':
			score += 3
		case ']':
			score += 2
		}
	}
	return score
}

func score(r rune) int {
	switch r {
	case ')':
		return 3
	case '>':
		return 25137
	case '}':
		return 1197
	case ']':
		return 57
	default:
		return 0
	}

}

func day10a(lines []string) {
	sum := 0
	for _, line := range lines {
		sum += checkLine(line)
	}
	fmt.Printf("Solution for part A: %d\n", sum)
}

func day10b(lines []string) {
	results := []int{}
	for _, line := range lines {
		result := checkLine2(line)
		if result != -1 {
			results = append(results, result)
		}
	}
	sort.Ints(results)
	result := results[len(results)/2]
	fmt.Printf("Solution for part B: %d\n", result)
}
