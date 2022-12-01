package main

import (
	"aoc2021/util"
	"fmt"
	"time"
)

func main() {
	lines := util.GetLines("day15\\15.example")
	start := time.Now()
	partA(lines)
	duration := time.Since(start)
	partB(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

type point struct {
	x, y int
}

var fScores = map[point]int{}
var gScores = map[point]int{}

const infinity = 2147483646

func getFScore(p point) int {
	value, ok := fScores[p]
	if !ok {
		fScores[p] = infinity
		return infinity
	}
	return value
}

func setFScore(p point, value int) {
	fScores[p] = value
}

func getGScore(p point) int {
	value, ok := gScores[p]
	if !ok {
		gScores[p] = infinity
		return infinity
	}
	return value
}

func setGScore(p point, value int) {
	gScores[p] = value
}

func h(p point, width, height int) int {
	// dijkstra
	return 0
	return width - p.x + height - p.y - 2
}

func lowestFScore(points []point) (point, int) {
	lowest := infinity
	lowestP := points[0]
	pos := -1
	for i, p := range points {
		fscore := getFScore(p)
		if fscore < lowest {
			lowest = fscore
			lowestP = p
			pos = i
		}
	}
	return lowestP, pos
}

func removeFromPointSlice(slice []point, pos int) []point {
	if len(slice) == 1 {
		return []point{}
	}
	if pos < 0 || pos >= len(slice) {
		fmt.Errorf("removed at invalid position from point slice")
	}
	newLength := len(slice) - 1
	slice[pos] = slice[newLength] // replace with last element
	return slice[:newLength]
}

func giveNeighbors(width, height int, p point) []point {
	neighbors := []point{}
	// no diag-cells included
	offsets := []point{
		{-1, 0}, {1, 0},
		{0, -1}, {0, 1},
	}
	for _, offset := range offsets {
		x := p.x + offset.x
		y := p.y + offset.y
		if x >= 0 && y >= 0 && x < width && y < height {
			neighbors = append(neighbors, point{x, y})
		}

	}
	return neighbors
}

func addToPointSet(newP point, slice []point) []point {

	// Naive implementaion
	for _, p := range slice {
		if p == newP {
			return slice
		}
	}
	return append(slice, newP)
}

func addToPrioPointQueue(newP point, queue []point) []point {
	//sorted by lowest fscore
	if len(queue) == 0 {
		fmt.Printf("Was empty, adding %v\n", newP)
		return []point{newP}
	}

	currentFScore := getFScore(newP)
	newQueue := []point{}
	notInserted := true

	for _, el := range queue {
		if el == newP {
			continue
		}
		if notInserted && getFScore(el) >= currentFScore {
			newQueue = append(newQueue, newP)
			notInserted = false
		}
		newQueue = append(newQueue, el)
	}

	if notInserted {
		newQueue = append(newQueue, newP)
	}

	fmt.Println(queue)
	for _, el := range newQueue {
		fmt.Printf("Point: %v, FScore: %d\n", el, getFScore(el))
	}
	return newQueue
}

func givePriorityElementAndNewQueue(queue []point) (point, []point) {
	fmt.Printf("giving back: %v and %v\n", queue[0], len(queue[1:]))
	return queue[0], queue[1:]
}

func reconstructPath(cameFrom map[point]point, n, start point) []point {
	path := []point{n}
	current := n
	for current != start {
		previous := cameFrom[current]
		path = append(path, previous)
		current = previous
	}
	return path
}

func pointIncludedInSlice(slice []point, candidate point) bool {
	for _, p := range slice {
		if p == candidate {
			return true
		}
	}
	return false
}

func partA(lines []string) {

	// setting up the map with its values
	width := len(lines[0])
	height := len(lines)
	pointScores := make([][]int, height)
	for i := 0; i < width; i++ {
		pointScores[i] = make([]int, width)
	}
	for i, line := range lines {
		for j, value := range line {
			pointScores[i][j] = int(value - '0')
		}
	}

	// setting up A*
	start := point{0, 0}
	goal := point{height - 1, width - 1}
	openSet := []point{start}
	cameFrom := map[point]point{}
	path := []point{}

	setGScore(start, 0)
	setFScore(start, h(start, width, height))

	for len(openSet) > 0 {
		// check next in line with lowest score
		current, pos := lowestFScore(openSet)
		if current == goal {
			path = reconstructPath(cameFrom, current, start)
			break
		}
		// remove from openSet
		openSet = removeFromPointSlice(openSet, pos)

		// check neighbors
		neighbors := giveNeighbors(width, height, current)
		currentGScore := getGScore(current)
		for _, n := range neighbors {
			tentativeGScore := currentGScore + pointScores[n.y][n.x]
			if tentativeGScore < getGScore(n) {
				setGScore(n, tentativeGScore)
				setFScore(n, tentativeGScore+h(n, width, height))
				cameFrom[n] = current
				openSet = addToPointSet(n, openSet)
			}
		}

	}
	//printPath(path, height, width, pointScores)
	totalScore := calculateScore(path, pointScores)
	fmt.Printf("Solution for part A: %v\n", totalScore)
}

func printPath(path []point, height, width int, pointScores [][]int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			onPath := pointIncludedInSlice(path, point{j, i})
			if onPath {
				fmt.Print("\033[33m", getPointScore(pointScores, j, i), "\033[0m")
			} else {
				fmt.Print(getPointScore(pointScores, j, i))
			}
		}
		fmt.Println()
	}
}

func calculateScore(path []point, pointScores [][]int) int {
	score := -getPointScore(pointScores, 0, 0)
	for _, p := range path {
		score += getPointScore(pointScores, p.x, p.y)
	}
	return score
}

func getPointScore(pointScores [][]int, x, y int) int {
	height := len(pointScores)
	width := len(pointScores[0])

	if x < width && y < height {
		return pointScores[y][x]
	}
	addX := x / width
	addY := y / width
	newValue := pointScores[y%height][x%width] + addX + addY
	for newValue > 9 {
		newValue -= 9
	}
	return newValue
}

func partB(lines []string) {
	// reset globals
	fScores = map[point]int{}
	gScores = map[point]int{}

	// setting up the map with its values
	width := len(lines[0])
	height := len(lines)
	pointScores := make([][]int, height)
	for i := 0; i < width; i++ {
		pointScores[i] = make([]int, width)
	}
	for i, line := range lines {
		for j, value := range line {
			pointScores[i][j] = int(value - '0')
		}
	}

	// scale map
	width, height = width*5, height*5

	// setting up A*
	start := point{0, 0}
	goal := point{height - 1, width - 1}
	openSet := []point{start}
	cameFrom := map[point]point{}
	path := []point{}

	setGScore(start, 0)
	setFScore(start, h(start, width, height))

	for len(openSet) > 0 {
		fmt.Println("New Iteration")
		fmt.Println(openSet)
		// check next in line with lowest score
		//current, pos := lowestFScore(openSet)
		current, openSet := givePriorityElementAndNewQueue(openSet)
		if current == goal {
			path = reconstructPath(cameFrom, current, start)
			break
		}
		fmt.Println(openSet)

		// remove from openSet
		//openSet = removeFromPointSlice(openSet, pos)

		// check neighbors
		neighbors := giveNeighbors(width, height, current)
		currentGScore := getGScore(current)
		for _, n := range neighbors {
			fmt.Printf("Checking element: %v\n", n)
			tentativeGScore := currentGScore + getPointScore(pointScores, n.x, n.y)
			fmt.Printf("Tentative: %v, GScore %v\n", tentativeGScore, getGScore(n))
			if tentativeGScore < getGScore(n) {
				setGScore(n, tentativeGScore)
				setFScore(n, tentativeGScore+h(n, width, height))
				cameFrom[n] = current
				openSet = addToPrioPointQueue(n, openSet)
				fmt.Printf("Added element: %v, newQueue: %v\n", n, openSet)
			}
		}

	}
	//printPath(path, height, width, pointScores)
	totalScore := calculateScore(path, pointScores)
	fmt.Printf("Solution for part B: %v\n", totalScore)
}
