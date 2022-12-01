package main

import (
	"aoc2021/util"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
	"sort"
	"time"
)

func main() {
	lines := util.GetLines("day09\\09.in")
	start := time.Now()
	day09a(lines)
	duration := time.Since(start)
	day09b(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
	//makeGif(lines)
}

func day09a(lines []string) {
	sum := 0
	height := len(lines)
	width := len(lines[0])
	for i, line := range lines {
		for j, r := range line {
			if i > 0 && rune(lines[i-1][j]) <= r {
				continue
			}
			if i < height-1 && rune(lines[i+1][j]) <= r {
				continue
			}
			if j < width-1 && rune(lines[i][j+1]) <= r {
				continue
			}
			if j > 0 && rune(lines[i][j-1]) <= r {
				continue
			}
			sum += int(r-'0') + 1
		}
	}
	fmt.Printf("Solution for part A: %d\n", sum)
}

type point struct {
	x, y int
}

func day09b(lines []string) {
	height := len(lines)
	width := len(lines[0])
	// mapping keeps track of scanned points and their assigned basin
	mapping := make([][]int, height)
	for i := range mapping {
		mapping[i] = make([]int, width)
	}
	// running list of basins with their respective points
	basins := map[int][]point{}

	for i, line := range lines {
		for j, r := range line {
			if r == '9' {
				mapping[i][j] = -1
				continue
			}
			// check previous element
			p := point{i, j}
			// map to same basin as previous element in row
			if j > 0 && rune(lines[i][j-1]) != '9' {
				mapping[i][j] = mapping[i][j-1]
				basins[mapping[i][j-1]] = append(basins[mapping[i][j-1]], p)
				// map to new basin
			} else {
				mapping[i][j] = len(basins)
				basins[len(basins)] = []point{p}
			}
			// if above is already mapped, take mapping and transform all connected points as well
			if i > 0 && rune(lines[i-1][j]) != '9' {
				oldMapping := mapping[i][j]
				newMapping := mapping[i-1][j]
				if oldMapping == newMapping {
					continue
				}
				for _, v := range basins[oldMapping] {
					mapping[v.x][v.y] = newMapping
					basins[newMapping] = append(basins[newMapping], v)
				}
				basins[oldMapping] = nil
			}
		}
	}
	sizes := []int{}
	for _, basin := range basins {
		sizes = append(sizes, len(basin))
	}
	sort.Ints(sizes)
	l := len(sizes)
	result := sizes[l-1] * sizes[l-2] * sizes[l-3]
	fmt.Printf("Solution for part B: %d\n", result)
}

func createPalette(numColors int) []color.Color {
	palette := palette.Plan9
	return palette
}

func makeGif(lines []string) {
	height := len(lines)
	width := len(lines[0])
	scale := 4
	palette := createPalette(2000)
	rect := image.Rect(0, 0, width*scale, height*scale)
	timings := []int{}
	frames := []*image.Paletted{}

	// mapping keeps track of scanned points and their assigned basin
	mapping := make([][]int, height)
	for i := range mapping {
		mapping[i] = make([]int, width)
	}
	// running list of basins with their respective points
	basins := map[int][]point{}
	for i, line := range lines {
		for j, r := range line {
			newImg := image.NewPaletted(rect, palette)
			if r == '9' {
				mapping[i][j] = 0
			} else {
				// check previous element
				p := point{i, j}
				// map to same basin as previous element in row
				if j > 0 && rune(lines[i][j-1]) != '9' {
					mapping[i][j] = mapping[i][j-1]
					basins[mapping[i][j-1]] = append(basins[mapping[i][j-1]], p)
					// map to new basin
				} else {
					mapping[i][j] = len(basins)
					basins[len(basins)] = []point{p}
				}
				// if above is already mapped, take mapping and transform all connected points as well
				if i > 0 && rune(lines[i-1][j]) != '9' {
					oldMapping := mapping[i][j]
					newMapping := mapping[i-1][j]
					if oldMapping == newMapping {
						continue
					}
					for _, v := range basins[oldMapping] {
						mapping[v.x][v.y] = newMapping
						basins[newMapping] = append(basins[newMapping], v)
					}
					basins[oldMapping] = nil
				}
			}
			for m := range lines {
				for n := range lines[0] {
					tileColor := uint8((mapping[m][n] + 1) % 256)
					for t := 0; t < scale; t++ {
						for k := 0; k < scale; k++ {
							if m < len(lines)-1 && n < len(lines[0])-1 {
								newImg.SetColorIndex(n*scale+k, m*scale+t, tileColor)
							}
						}
					}
				}

			}
			frames = append(frames, newImg)
			timings = append(timings, 0)
		}
	}
	anim := gif.GIF{Delay: timings, Image: frames, LoopCount: 0}
	out, err := os.Create("9b_scanning.gif")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	fmt.Println(len(lines[0]) * len(lines))
	fmt.Println(len(frames))
	e := gif.EncodeAll(out, &anim)
	fmt.Println(e)
}
