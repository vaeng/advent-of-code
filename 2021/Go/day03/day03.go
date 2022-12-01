package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func tellMostCommon(codes []string, pos int) rune {
	l := len(codes)
	onecounter := 0
	for _, code := range codes {
		if code[pos] == '1' {
			onecounter++
		}
	}
	if onecounter*2 >= l {
		return '1'
	}
	return '0'
}

func tellLeastCommon(codes []string, pos int) rune {
	l := len(codes)
	onecounter := 0
	for _, code := range codes {
		if code[pos] == '1' {
			onecounter++
		}
	}
	if onecounter*2 >= l {
		return '0'
	}
	return '1'
}

func filterByPos(codes []string, pos int, r rune) []string {
	filtered := []string{}
	for _, code := range codes {
		if rune(code[pos]) == r {
			filtered = append(filtered, code)
		}
	}
	return filtered
}

func main() {
	file, err := os.Open("day03\\03.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ones []int
	var zeros []int
	var codes []string
	first := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := scanner.Text()
		codes = append(codes, v)
		for i, r := range v {
			if first {
				ones = append(ones, 0)
				zeros = append(zeros, 0)
			}
			if r == '1' {
				ones[i]++
			} else {
				zeros[i]++
			}
		}
		first = false
	}
	s := ""
	t := ""
	for i := range ones {
		if ones[i] > zeros[i] {
			s = s + "1"
			t = t + "0"
		} else {
			s = s + "0"
			t = t + "1"
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	a, _ := strconv.ParseInt(s, 2, 64)
	b, _ := strconv.ParseInt(t, 2, 64)
	fmt.Println(a * b)

	oxygen := codes
	co2 := codes
	for i := 0; i < len(codes[0]); i++ {
		if len(oxygen) > 1 {
			mc := tellMostCommon(oxygen, i)
			oxygen = filterByPos(oxygen, i, mc)
		}
		if len(co2) > 1 {
			lc := tellLeastCommon(co2, i)
			co2 = filterByPos(co2, i, lc)
		}
	}

	ab, _ := strconv.ParseInt(oxygen[0], 2, 64)
	bb, _ := strconv.ParseInt(co2[0], 2, 64)
	fmt.Printf("%d * %d = %d", ab, bb, ab*bb)
}
