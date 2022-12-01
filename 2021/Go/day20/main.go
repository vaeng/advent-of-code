package main

import (
	"aoc2021/util"
	"fmt"
	"time"
)

func main() {
	lines := util.GetLines(".\\day20\\in")
	start := time.Now()
	partA(lines)
	duration := time.Since(start)
	partB(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

var algo [512]bool

func fillAlgo(s string) {
	for i, r := range s {
		if r == '#' {
			algo[i] = true
		} else if r == '.' {
			algo[i] = false
		} else {
			panic("invalid character in algo init")
		}

	}
}

func loadImage(lines []string) [][]bool {
	image := [][]bool{}
	for _, line := range lines {
		row := []bool{}
		for _, r := range line {
			if r == '#' {
				row = append(row, true)
			} else if r == '.' {
				row = append(row, false)
			} else {
				panic("invalid character in image loading")
			}
		}
		image = append(image, row)
	}
	return image
}

func printImage(image [][]bool) {
	strImage := "\n"
	for _, row := range image {
		for _, b := range row {
			if b {
				strImage += "#"
			} else {
				strImage += "."
			}
		}
		strImage += "\n"
	}
	fmt.Print(strImage)
}

func countLightPixels(image [][]bool) int {
	counter := 0
	for _, row := range image {
		for _, b := range row {
			if b {
				counter++
			}
		}
	}
	return counter
}

func getPixel(image [][]bool, i, j, iteration int) bool {
	height := len(image)
	width := len(image[0])
	infinityValue := false
	// check if dark changes to bright in infinity
	if algo[0] {
		// check if complete bright changes to dark again
		if algo[511] { // stays bright
			infinityValue = true
		} else { // flip from dark to bright
			infinityValue = iteration%2 == 1
		}
	}
	if iteration < 1 {
		infinityValue = false
	}

	if i < 0 || i >= height || j < 0 || j >= width {
		return infinityValue
	}
	return image[i][j]
}

func enhancePixel(image [][]bool, i, j, iteration int) bool {
	// start with least important bit in the lower right corner
	multi := 1
	index := 0
	for row := 1; row >= -1; row-- {
		for col := 1; col >= -1; col-- {
			pixel := getPixel(image, i+row, j+col, iteration)
			if pixel {
				index += multi
			}
			multi *= 2
		}
	}
	return algo[index]
}

func recalibrateFrame(image [][]bool, iteration int) [][]bool {
	changes := true
	for changes {
		image, changes = addBoarder(image, false, iteration)
	}
	return image
}

func addBoarder(image [][]bool, forced bool, iteration int) ([][]bool, bool) {
	height := len(image)
	width := len(image[0])
	infinityValue := getPixel(image, -1, -1, iteration)
	// check left boarder, three deep
	left := false
	for row := 0; row < height; row++ {
		if image[row][0] != infinityValue ||
			image[row][1] != infinityValue ||
			image[row][2] != infinityValue {
			left = true
			break
		}
	}
	//check right boarder
	right := false
	for row := 0; row < height; row++ {
		if image[row][width-1] != infinityValue ||
			image[row][width-2] != infinityValue ||
			image[row][width-3] != infinityValue {
			right = true
			break
		}
	}
	// check top boarder
	top := false
	for col := 0; col < height; col++ {
		if image[0][col] != infinityValue ||
			image[1][col] != infinityValue ||
			image[2][col] != infinityValue {
			top = true
			break
		}
	}
	// check bottom boarder
	bottom := false
	for col := 0; col < height; col++ {
		if image[height-1][col] != infinityValue ||
			image[height-2][col] != infinityValue ||
			image[height-3][col] != infinityValue {
			bottom = true
			break
		}
	}
	if forced {
		left, right, bottom, top = true, true, true, true
	}
	recalibratedImage := [][]bool{}
	if top {
		emptyRow := []bool{}
		if left {
			emptyRow = append(emptyRow, infinityValue)
			emptyRow = append(emptyRow, infinityValue)
			emptyRow = append(emptyRow, infinityValue)
		}
		for range image[0] {
			emptyRow = append(emptyRow, infinityValue)
		}
		if right {
			emptyRow = append(emptyRow, infinityValue)
			emptyRow = append(emptyRow, infinityValue)
			emptyRow = append(emptyRow, infinityValue)
		}
		// make three new top lines
		recalibratedImage = append(recalibratedImage, emptyRow)
		recalibratedImage = append(recalibratedImage, emptyRow)
		recalibratedImage = append(recalibratedImage, emptyRow)
	}
	for _, originalLine := range image {
		newLine := []bool{}
		if left {
			newLine = append(newLine, infinityValue)
			newLine = append(newLine, infinityValue)
			newLine = append(newLine, infinityValue)
		}
		newLine = append(newLine, originalLine...)
		if right {
			newLine = append(newLine, infinityValue)
			newLine = append(newLine, infinityValue)
			newLine = append(newLine, infinityValue)
		}
		recalibratedImage = append(recalibratedImage, newLine)
	}
	if bottom {
		emptyRow := []bool{}
		for range recalibratedImage[0] {
			emptyRow = append(emptyRow, infinityValue)
		}
		recalibratedImage = append(recalibratedImage, emptyRow)
		recalibratedImage = append(recalibratedImage, emptyRow)
		recalibratedImage = append(recalibratedImage, emptyRow)
	}
	boardersAdded := left || right || top || bottom
	return recalibratedImage, boardersAdded
}

func enhanceImage(image [][]bool, iteration int) [][]bool {
	height := len(image)
	width := len(image[0])
	enhancedImage := [][]bool{}
	for i := 0; i < height; i++ {
		row := []bool{}
		for j := 0; j < width; j++ {
			row = append(row, enhancePixel(image, i, j, iteration-1))
		}
		enhancedImage = append(enhancedImage, row)
	}
	enhancedImage = recalibrateFrame(enhancedImage, iteration)
	return enhancedImage
}

func partA(lines []string) {
	// setup
	fillAlgo(lines[0])
	image := loadImage(lines[2:])
	//printImage(image)
	image, _ = addBoarder(image, true, 0)
	//printImage(image)
	goal := 2
	for i := 1; i <= goal; i++ {
		image = enhanceImage(image, i)
	}
	result := countLightPixels(image)
	fmt.Printf("Solution for part A: %v\n", result)
}

func partB(lines []string) {
	// setup
	fillAlgo(lines[0])
	image := loadImage(lines[2:])
	image, _ = addBoarder(image, true, 0)
	goal := 50
	for i := 1; i <= goal; i++ {
		image = enhanceImage(image, i)
	}
	result := countLightPixels(image)
	fmt.Printf("Solution for part B: %v\n", result)
}
