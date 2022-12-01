package main

import (
	"aoc2021/util"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

func main() {
	lines := util.GetLines("day08\\08.in")
	start := time.Now()
	day08a(lines)
	duration := time.Since(start)
	day08b(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func countDigits(line string) int {
	numbers := 0
	groups := strings.Fields(line[60:])
	for _, group := range groups {
		l := len(group)
		switch l {
		case 2, 3, 4, 7:
			numbers++
		}
	}
	return numbers
}

func stringInside(a, b string) bool {
	for _, r := range a {
		if !strings.ContainsRune(b, r) {
			return false
		}
	}
	return true
}

func overlap(a, b string) int {
	overlap := 0
	for _, r := range a {
		if strings.ContainsRune(b, r) {
			overlap++
		}
	}
	return overlap
}

// Taken from: https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func decodeDigits(line string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	decoder := map[string]int{}
	encoder := map[int]string{}
	encoding := []string{}
	groups := []string{}

	//sort fields
	for i, v := range strings.Fields(line) {
		sorted := sortString(v)
		if i < 10 {
			encoding = append(encoding, sorted)
		} else if i > 10 {
			groups = append(groups, sorted)
		}
	}

	//set easy items
	for _, code := range encoding {
		switch len(code) {
		case 2:
			decoder[code] = 1
			encoder[1] = code
		case 3:
			decoder[code] = 7
			encoder[7] = code
		case 4:
			decoder[code] = 4
			encoder[4] = code
		case 7:
			decoder[code] = 8
			encoder[8] = code
		}
	}
	//set dependend
	for _, code := range encoding {
		switch len(code) {
		case 5:
			if stringInside(encoder[1], code) {
				decoder[code] = 3
			} else if overlap(encoder[4], code) == 3 {
				decoder[code] = 5
			} else {
				decoder[code] = 2
			}
		case 6:
			if stringInside(encoder[4], code) {
				decoder[code] = 9
			} else if stringInside(encoder[1], code) {
				decoder[code] = 0
			} else {
				decoder[code] = 6
			}
		}
	}
	ch <- decoder[groups[0]]*1000 + decoder[groups[1]]*100 +
		decoder[groups[2]]*10 + decoder[groups[3]]
}

func day08a(lines []string) {
	sum1478 := 0
	for _, line := range lines {
		sum1478 += countDigits(line)
	}
	fmt.Printf("Solution for part A: %d\n", sum1478)
}

func day08b(lines []string) {
	sum := 0
	sampleChan := make(chan int, len(lines))
	var wg sync.WaitGroup
	for _, line := range lines {
		wg.Add(1)
		go decodeDigits(line, sampleChan, &wg)
	}
	wg.Wait()
	close(sampleChan)
	for s := range sampleChan {
		sum += s
	}
	fmt.Printf("Solution for part B: %d\n", sum)
}
