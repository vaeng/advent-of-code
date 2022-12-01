package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	lines := util.GetLines(".\\day19\\in")
	start := time.Now()
	partA(lines)
	duration := time.Since(start)
	partB(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

type point struct {
	x, y, z int
}

func partA(lines []string) {
	scanners := loadScanners(lines)
	scannerPositions := []point{{0, 0, 0}}
outer:
	for len(scanners) > 1 {
		for i := range scanners {
			for j := range scanners {
				if i == j {
					continue
				}
				overlap, newPoints, scannerPosition := checkOverlapAndReturnNewPoints(scanners[i], scanners[j])
				if overlap {
					scannerPositions = append(scannerPositions, scannerPosition)
					scanners[i] = append(scanners[i], newPoints...)
					newScanners := scanners[:j]
					newScanners = append(newScanners, scanners[j+1:]...)
					scanners = newScanners
					continue outer
				}
			}
		}
	}
	fmt.Printf("Solution for part A: %v\n", len(scanners[0]))
	maxDistance := 0
	for i, p1 := range scannerPositions {
		for j, p2 := range scannerPositions {
			if j == i {
				continue
			}
			currentDistance := manhattenDistance(p1, p2)
			if currentDistance > maxDistance {
				maxDistance = currentDistance
			}
		}
	}

	fmt.Printf("Solution for part B: %v\n", maxDistance)
}

func manhattenDistance(p1, p2 point) int {
	x := p1.x - p2.x
	y := p1.y - p2.y
	z := p1.z - p2.z
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if z < 0 {
		z = -z
	}
	return x + y + z
}

func partB(lines []string) {

	fmt.Printf("Solution for part B: %v\n", 0)
}

func loadScanners(lines []string) [][]point {
	scanners := [][]point{}
	scanline := []point{}
	for i, line := range lines {
		// skip first line
		if i == 0 || len(line) == 0 {
			continue
		}
		// start new scanner
		if line[1] == '-' {
			scanners = append(scanners, scanline)
			scanline = []point{}
			continue
		}
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		scanline = append(scanline, point{x, y, z})
	}
	scanners = append(scanners, scanline)
	return scanners
}

func shufflePoint(p point, config int) point {
	x, y, z := p.x, p.y, p.z
	switch config {
	case 1:
		return point{x, z, y}
	case 2:
		return point{y, x, z}
	case 3:
		return point{y, z, x}
	case 4:
		return point{z, x, y}
	case 5:
		return point{z, y, x}
	default: //identity
		return point{x, y, z}
	}
}

func turnPoint(p point, config int) point {
	x, y, z := p.x, p.y, p.z
	switch config {
	case 1:
		return point{-x, y, z}
	case 2:
		return point{x, -y, z}
	case 3:
		return point{x, y, -z}
	case 4:
		return point{-x, -y, z}
	case 5:
		return point{-x, y, -z}
	case 6:
		return point{x, -y, -z}
	case 7:
		return point{-x, -y, -z}
	default: //identity
		return point{x, y, z}
	}
}

func checkRotation(origin, candidate point) (bool, int, int) {
	for shuffleIndex := 0; shuffleIndex < 6; shuffleIndex++ {
		for turnIndex := 0; turnIndex < 8; turnIndex++ {
			if origin == turnPoint(shufflePoint(candidate, shuffleIndex), turnIndex) {
				return true, shuffleIndex, turnIndex
			}
		}
	}
	return false, -1, -1
}

func generateRotatedScanner(scanner []point, shuffleIndex, turnIndex int) []point {
	rotatedScanner := []point{}
	for _, p := range scanner {
		rotatedP := turnPoint(shufflePoint(p, shuffleIndex), turnIndex)
		rotatedScanner = append(rotatedScanner, rotatedP)
	}
	return rotatedScanner
}

func generateRotatedScanners(scanner []point) [][]point {
	rotatedScanners := [][]point{}
	for shuffleIndex := 0; shuffleIndex < 6; shuffleIndex++ {
		for turnIndex := 0; turnIndex < 8; turnIndex++ {
			rotatedScanners = append(rotatedScanners, generateRotatedScanner(scanner, shuffleIndex, turnIndex))
		}
	}
	return rotatedScanners
}

func correctScannerByOffset(scanner []point, newReference, oldReference point) ([]point, point) {
	offset := subtractPoints(newReference, oldReference)
	corectedScanner := []point{}
	for _, p := range scanner {
		corectedScanner = append(corectedScanner, addPoints(p, offset))
	}
	return corectedScanner, offset
}

// if a is (0, 0, 0) and b is (1, 2, 3) than the offset should be (-1,-2,3)
// if a is (5, -4, 6) and b is (1, 2, 3) than the offset should be (4,-6,3)
func subtractPoints(a, b point) point {
	x := a.x - b.x
	y := a.y - b.y
	z := a.z - b.z
	return point{x, y, z}
}

func addPoints(a, b point) point {
	x := a.x + b.x
	y := a.y + b.y
	z := a.z + b.z
	return point{x, y, z}
}

func countAndListCommonPoints(base, tester []point) (int, []point, []point) {
	common := []point{}
	notIncluded := []point{}
	for _, p := range tester {
		if includedInScanner(base, p) {
			common = append(common, p)
		} else {
			notIncluded = append(notIncluded, p)
		}
	}
	return len(common), common, notIncluded
}

func includedInScanner(base []point, candidate point) bool {
	for _, p := range base {
		if p == candidate {
			return true
		}
	}
	return false
}

func checkOverlapAndReturnNewPoints(base, candidates []point) (bool, []point, point) {
	for shuffleIndex := 0; shuffleIndex < 6; shuffleIndex++ {
		for turnIndex := 0; turnIndex < 8; turnIndex++ {
			rotatedCandidate := generateRotatedScanner(candidates, shuffleIndex, turnIndex)
			for _, basePoint := range base {
				for _, candidatePoint := range rotatedCandidate {
					correctedAndOffsetCandidate, offset := correctScannerByOffset(rotatedCandidate, basePoint, candidatePoint)
					commonNumber, _, unknown := countAndListCommonPoints(base, correctedAndOffsetCandidate)
					if commonNumber >= 12 {
						return true, unknown, offset
					}
				}
			}

		}
	}
	return false, nil, point{}
}
