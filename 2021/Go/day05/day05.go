package main

import (
	"aoc2021/util"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main() {
	lines := util.GetLines("day05\\05.in")
	start := time.Now()
	day05a(lines)
	duration := time.Since(start)
	day05b(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

// Pair is a struct with x and y
type Pair struct {
	x, y int
}

func mustAtoi(line []byte) int {
	i, err := strconv.Atoi(string(line))
	if err != nil {
		panic(err)
	}
	return i
}

func parsePairs(line string) (Pair, Pair) {
	var start, end Pair
	var num []byte // accumulator
	var state int  // parser state: 0: x1 1: y1 2: x2 3: y2
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num = append(num, line[i])
		case ' ':
			start.y = mustAtoi(num)
			num = []byte{}
			i += 3 // jump past '-> '
			state = 2
		case ',':
			if state == 0 {
				start.x = mustAtoi(num)
				num = []byte{}
				state = 1
			}
			if state == 2 {
				end.x = mustAtoi(num)
				num = []byte{}
				state = 3
			}
		}
	}
	end.y = mustAtoi(num)
	return start, end
}

var (
	vents map[Pair]int
	re    = regexp.MustCompile("^(?P<x1>[0-9]+),(?P<y1>[0-9]+)+ -> (?P<x2>[0-9]+),(?P<y2>[0-9]+)$")
)

func day05a(lines []string) {
	vents = map[Pair]int{}
	counter := 0
	for _, line := range lines {
		var p1, p2 Pair
		p1, p2 = parsePairs(line)
		x1 := p1.x
		x2 := p2.x
		y1 := p1.y
		y2 := p2.y

		// no diags
		if x1 == x2 || y1 == y2 {
			var xs []int
			var ys []int
			if x2 < x1 {
				x1, x2 = x2, x1
			}
			if y2 < y1 {
				y1, y2 = y2, y1
			}
			for i := x1; i <= x2; i++ {
				xs = append(xs, i)
			}
			for i := y1; i <= y2; i++ {
				ys = append(ys, i)
			}
			for _, x := range xs {
				for _, y := range ys {
					p := Pair{x, y}
					_, ok := vents[p]
					if !ok {
						vents[p] = 1
					} else {
						vents[p]++
					}
				}
			}
		}
	}
	for _, value := range vents {
		if value > 1 {
			counter++
		}
	}
	fmt.Printf("Solution for part A: %d\n", counter)
}

func day05b(lines []string) {
	vents = map[Pair]int{}
	counter := 0
	for _, line := range lines {
		m := re.FindAllStringSubmatch(line, -1)
		x1, _ := strconv.Atoi(m[0][1])
		y1, _ := strconv.Atoi(m[0][2])
		x2, _ := strconv.Atoi(m[0][3])
		y2, _ := strconv.Atoi(m[0][4])

		// no diags
		if x1 == x2 || y1 == y2 {
			var xs []int
			var ys []int
			if x2 < x1 {
				x1, x2 = x2, x1
			}
			if y2 < y1 {
				y1, y2 = y2, y1
			}
			for i := x1; i <= x2; i++ {
				xs = append(xs, i)
			}
			for i := y1; i <= y2; i++ {
				ys = append(ys, i)
			}
			for _, x := range xs {
				for _, y := range ys {
					p := Pair{x, y}
					_, ok := vents[p]
					if !ok {
						vents[p] = 1
					} else {
						vents[p]++
					}
				}
			}
		}
		// special diags
		dx := x2 - x1
		xstep := 1
		dy := y2 - y1
		ystep := 1
		if dx < 0 {
			dx *= -1
			xstep *= -1
		}
		if dy < 0 {
			dy *= -1
			ystep *= -1
		}
		if dx == dy {
			pairs := []Pair{}
			x := x1
			y := y1
			for dx >= 0 {
				pairs = append(pairs, Pair{x, y})
				x += xstep
				y += ystep
				dx--
			}
			for _, p := range pairs {
				_, ok := vents[p]
				if !ok {
					vents[p] = 1
				} else {
					vents[p]++
				}
			}
		}
	}
	for _, value := range vents {
		if value > 1 {
			counter++
		}
	}
	fmt.Printf("Solution for part B: %d\n", counter)
}
