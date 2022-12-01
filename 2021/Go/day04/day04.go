package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := util.GetLines("day04\\04.in")

	day04a(lines)
	day04b(lines)
}

type board struct {
	row1, row2, row3, row4, row5 []int
	col1, col2, col3, col4, col5 []int
	winner                       bool
}

func newBoard(row1, row2, row3, row4, row5 []int) board {
	col1, col2, col3, col4, col5 := make([]int, 5), make([]int, 5), make([]int, 5), make([]int, 5), make([]int, 5)
	col1[0] = row1[0]
	col1[1] = row2[0]
	col1[2] = row3[0]
	col1[3] = row4[0]
	col1[4] = row5[0]

	col2[0] = row1[1]
	col2[1] = row2[1]
	col2[2] = row3[1]
	col2[3] = row4[1]
	col2[4] = row5[1]

	col3[0] = row1[2]
	col3[1] = row2[2]
	col3[2] = row3[2]
	col3[3] = row4[2]
	col3[4] = row5[2]

	col4[0] = row1[3]
	col4[1] = row2[3]
	col4[2] = row3[3]
	col4[3] = row4[3]
	col4[4] = row5[3]

	col5[0] = row1[4]
	col5[1] = row2[4]
	col5[2] = row3[4]
	col5[3] = row4[4]
	col5[4] = row5[4]

	return board{row1, row2, row3, row4, row5, col1, col2, col3, col4, col5, false}
}

func (b *board) addNum(num int) {
	for i := range b.row1 {
		switch num {
		case b.row1[i]:
			b.row1[i] -= 100
		case b.row2[i]:
			b.row2[i] -= 100
		case b.row3[i]:
			b.row3[i] -= 100
		case b.row4[i]:
			b.row4[i] -= 100
		case b.row5[i]:
			b.row5[i] -= 100
		}
		switch num {
		case b.col1[i]:
			b.col1[i] -= 100
		case b.col2[i]:
			b.col2[i] -= 100
		case b.col3[i]:
			b.col3[i] -= 100
		case b.col4[i]:
			b.col4[i] -= 100
		case b.col5[i]:
			b.col5[i] -= 100
		}
	}
}

func (b *board) isWinner() bool {
	if b.winner ||
		allNeg(b.row1) ||
		allNeg(b.row2) ||
		allNeg(b.row3) ||
		allNeg(b.row4) ||
		allNeg(b.row5) ||
		allNeg(b.col1) ||
		allNeg(b.col2) ||
		allNeg(b.col3) ||
		allNeg(b.col4) ||
		allNeg(b.col5) {
		b.winner = true
		return true
	}
	return false
}

func (b *board) unmarkedSum() int {
	sum := 0
	sum += unusedPerRow(b.row1)
	sum += unusedPerRow(b.row2)
	sum += unusedPerRow(b.row3)
	sum += unusedPerRow(b.row4)
	sum += unusedPerRow(b.row5)
	return sum
}

func allNeg(slice []int) bool {
	for _, v := range slice {
		if v >= 0 {
			return false
		}
	}
	return true
}

func unusedPerRow(slice []int) int {
	sum := 0
	for _, v := range slice {
		if v > 0 {
			sum += v
		}
	}
	return sum
}

func convStringSliceToIntSlice(s string) []int {
	slice := strings.Fields(s)
	intslice := make([]int, len(slice))
	for i, s := range slice {
		intslice[i], _ = strconv.Atoi(s)
	}
	return intslice
}

func day04a(lines []string) {
	var drawn []int
	boardInput := make([][]int, 5)
	var boards []board
	unmarkedSum := 0
	lastNum := 0
	lineCounter := 0
	for i, line := range lines {
		if i == 0 {
			drawn = convStringSliceToIntSlice(strings.Replace(line, ",", " ", -1))
		} else if len(line) != 0 {
			newIntSlice := convStringSliceToIntSlice(line)
			boardInput[lineCounter] = newIntSlice
			lineCounter++
			if lineCounter == 5 {
				anotherBoard := newBoard(boardInput[0], boardInput[1], boardInput[2], boardInput[3], boardInput[4])
				boardInput = make([][]int, 5)
				boards = append(boards, anotherBoard)
				lineCounter = 0
			}
		}
	}
out:
	for _, num := range drawn {
		for _, b := range boards {
			b.addNum(num)
			if b.isWinner() {
				unmarkedSum = b.unmarkedSum()
				lastNum = num
				break out
			}
		}
	}
	fmt.Printf("Solution for part A: %d * %d = %d\n", unmarkedSum, lastNum, unmarkedSum*lastNum)
}

func day04b(lines []string) {
	var drawn []int
	boardInput := make([][]int, 5)
	var boards []board
	unmarkedSum := 0
	lastNum := 0
	lineCounter := 0
	for i, line := range lines {
		if i == 0 {
			drawn = convStringSliceToIntSlice(strings.Replace(line, ",", " ", -1))
		} else if len(line) != 0 {
			newIntSlice := convStringSliceToIntSlice(line)
			boardInput[lineCounter] = newIntSlice
			lineCounter++
			if lineCounter == 5 {
				anotherBoard := newBoard(boardInput[0], boardInput[1], boardInput[2], boardInput[3], boardInput[4])
				boardInput = make([][]int, 5)
				boards = append(boards, anotherBoard)
				lineCounter = 0
			}
		}
	}
	winners := 0
out:
	for _, num := range drawn {
		for _, b := range boards {
			if !b.isWinner() {
				b.addNum(num)
				if b.isWinner() {
					unmarkedSum = b.unmarkedSum()
					lastNum = num
					winners++
				}
			}
			if winners == len(boards) {
				break out
			}
		}
	}
	fmt.Printf("Solution for part B: %d * %d = %d\n", unmarkedSum, lastNum, unmarkedSum*lastNum)
}
