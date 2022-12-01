package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	lines := util.GetLines("day13\\13.example2")
	start := time.Now()
	day13a(lines)
	duration := time.Since(start)
	day13b(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

type point struct{ x, y int }
type fold struct {
	pos   int
	xFold bool
}

func day13a(lines []string) {
	folds := []fold{}
	positions := map[point]bool{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		if line[0] == 'f' {
			pos, _ := strconv.Atoi(line[13:])
			xFold := 'x' == line[11]
			folds = append(folds, fold{pos, xFold})
			break
		}
		xySlice := strings.Split(line, ",")
		x, _ := strconv.Atoi(xySlice[0])
		y, _ := strconv.Atoi(xySlice[1])
		positions[point{x, y}] = true
	}

	counter := 0
	pos := folds[0].pos
	xFold := folds[0].xFold

	for p := range positions {
		if xFold && p.x > pos {
			newX := 2*pos - p.x
			if positions[point{newX, p.y}] == false {
				counter++
			}
		} else if !xFold && p.y > pos {
			newY := 2*pos - p.y
			if positions[point{p.x, newY}] == false {
				counter++
			}
		} else {
			counter++
		}
	}

	fmt.Printf("Solution for part A: %d\n", counter)
}

func day13b(lines []string) {
	folds := []fold{}
	positions := map[point]bool{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		if line[0] == 'f' {
			pos, _ := strconv.Atoi(line[13:])
			xFold := 'x' == line[11]
			folds = append(folds, fold{pos, xFold})
			continue
		}
		xySlice := strings.Split(line, ",")
		x, _ := strconv.Atoi(xySlice[0])
		y, _ := strconv.Atoi(xySlice[1])
		positions[point{x, y}] = true
	}

	height := 0
	width := 0
	for _, f := range folds {
		pos := f.pos
		xFold := f.xFold
		currentHeight := 0
		currentWidth := 0
		newPositions := map[point]bool{}
		for p := range positions {
			newP := p
			if xFold && p.x > pos {
				newX := 2*pos - p.x
				newP = point{newX, p.y}
			} else if !xFold && p.y > pos {
				newY := 2*pos - p.y
				newP = point{p.x, newY}
			}
			newPositions[newP] = true
			if newP.x > currentWidth {
				currentWidth = newP.x
			}
			if newP.y > currentHeight {
				currentHeight = newP.y
			}
		}
		positions = newPositions
		width = currentWidth
		height = currentHeight
	}
	fmt.Printf("Solution for part B:\n")
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			if positions[point{x, y}] == true {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
