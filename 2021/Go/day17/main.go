package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"time"
)

func main() {
	lines := util.GetLines(".\\day17\\in")
	start := time.Now()
	partA(lines)
	duration := time.Since(start)
	partB(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func partA(lines []string) {
	xMin, xMax, yMin, yMax := getParameters(lines[0])
	hits := 0
	fmt.Println(xMin, xMax, yMin, yMax)
	vyGlobalMax := -yMin - 100
	vxGlobalMax := 0
	// steps
steploop:
	for s := 1; s < 300; s++ {
		for vx := 0; ; vx++ {
			actualX := actualXPos(vx, s)
			// not far enough
			if actualX < xMin {
				continue
			}
			// too far
			if actualX > xMax {
				break
			}
			vy := 0
			for ; ; vy++ {
				actualY := s*vy - (s*(s-1))/2
				// break when lower than target
				if actualY > yMax {
					continue steploop
				}
				if actualY < yMin {
					continue
				}
				// found a hit
				if vy > vyGlobalMax {
					vyGlobalMax = vy
					vxGlobalMax = vx
					//fmt.Printf("Step: %d, vx: %d vy: %d - Hit: (%d, %d)\n", s, vx, vy, actualX, actualY)
				}
			}
		}
	}
	maxY := vyGlobalMax * (vyGlobalMax + 1) / 2
	fmt.Printf("Solution for part A: %v, %v: max y: %d (total hits: %d)\n", vxGlobalMax, vyGlobalMax, maxY, hits)
}

// calcualte actual x position from step and vx
func actualXPos(vx, step int) int {
	if step >= vx+1 {
		return vx * (vx + 1) / 2
	}
	return step*vx - step*(step-1)/2
}

func getParameters(s string) (int, int, int, int) {
	var xmin, xmax, ymin, ymax int
	i := 0
	intStr := ""
	for true {
		if s[15+i] != '.' {
			intStr += string(s[15+i])
			i++
		} else {
			xmin, _ = strconv.Atoi(intStr)
			i += 2
			intStr = ""
			break
		}
	}
	for true {
		if s[15+i] != ',' {
			intStr += string(s[15+i])
			i++
		} else {
			xmax, _ = strconv.Atoi(intStr)
			i += 4
			intStr = ""
			break
		}
	}
	for true {
		if s[15+i] != '.' {
			intStr += string(s[15+i])
			i++
		} else {
			ymin, _ = strconv.Atoi(intStr)
			i += 2
			intStr = ""
			break
		}
	}
	for i != len(s)-15 {
		intStr += string(s[15+i])
		i++
	}
	ymax, _ = strconv.Atoi(intStr)

	return xmin, xmax, ymin, ymax
}

func partB(lines []string) {
	xMin, xMax, yMin, yMax := getParameters(lines[0])
	fmt.Printf("Solution for part B: %v\n", stepper(xMin, yMin, xMax, yMax))
}

func stepper(xMin, yMin, xMax, yMax int) int {
	counter := 0
xVelocity:
	for vx := 0; vx <= xMax; vx++ {
	yVelocity:
		for vy := yMin; ; vy++ {
			if vy >= 300 {
				continue xVelocity
			}
			for s := 0; s < 300; s++ {
				x := actualXPos(vx, s)
				y := s*vy - (s*(s-1))/2
				if x <= xMax && x >= xMin && y <= yMax && y >= yMin {
					counter++
					continue yVelocity
				}
			}
		}
	}
	return counter
}
