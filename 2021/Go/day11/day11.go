package main

import (
	"aoc2021/util"
	"fmt"
	"time"
)

func main() {
	lines := util.GetLines("day11\\11.in")
	start := time.Now()
	day11a(lines)
	duration := time.Since(start)
	day11b(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func day11a(lines []string) {
	flashes := 0
	// prepare grid
	grid := make([][]int, 10)
	for i := 0; i < 10; i++ {
		grid[i] = make([]int, 10)
	}
	// fil grid
	for i := range lines {
		for j := range lines[0] {
			grid[i][j] = int(lines[i][j] - '0')
		}
	}

	steps := 100
	for i := 0; i < steps; i++ {
		flashedOcti, grid := makeStep(grid)
		flashes += flashedOcti
		if i == steps {
			printlines(grid)
		}
	}

	fmt.Printf("Solution for part A: %d\n", flashes)
}

func printlines(grid [][]int) {
	fmt.Println("===ITERATION===")
	for i := range grid {
		for j := range grid[0] {
			fmt.Print(grid[i][j])
		}
		fmt.Print("\n")
	}
}

func makeStep(grid [][]int) (int, [][]int) {
	flashHappened := false
	// increase everything by 1
	for i := range grid {
		for j := range grid[0] {
			grid[i][j]++
			if grid[i][j] > 9 {
				flashHappened = true
			}
		}
	}
	// loop till everything has flashed
	// mark flashes with negativ number
	for flashHappened {
		flashHappened = false
		for i := range grid {
			for j := range grid[0] {
				if grid[i][j] > 9 {
					flashHappened = true
					grid[i][j] = -1
					for k := -1; k <= 1; k++ {
						for l := -1; l <= 1; l++ {
							x := i + k
							y := j + l
							if x < 10 && x >= 0 && y < 10 && y >= 0 {
								if grid[x][y] >= 0 {
									grid[x][y]++
								}
							}
						}
					}
				}
			}
		}
	}
	// count flashes and reset energy levels
	flashes := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] < 0 {
				flashes++
				grid[i][j] = 0
			}
		}
	}

	return flashes, grid
}

func day11b(lines []string) {
	// prepare grid
	grid := make([][]int, 10)
	for i := 0; i < 10; i++ {
		grid[i] = make([]int, 10)
	}
	// fil grid
	for i := range lines {
		for j := range lines[0] {
			grid[i][j] = int(lines[i][j] - '0')
		}
	}

	syncedAt := 0
	for i := 0; ; i++ {
		_, grid := makeStep(grid)
		if allFlash(grid) {
			syncedAt = i + 1
			break
		}
	}
	fmt.Printf("Solution for part B: %d\n", syncedAt)
}

func allFlash(grid [][]int) bool {
	sum := 0
	for i := range grid {
		for j := range grid[0] {
			sum += grid[i][j]
		}
	}
	return sum == 0
}
