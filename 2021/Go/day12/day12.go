package main

import (
	"aoc2021/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	lines := util.GetLines("day12\\12.in")
	start := time.Now()
	day12a(lines)
	duration := time.Since(start)
	day12b(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func makeConnections(lines []string) map[string][]string {
	connections := map[string][]string{}
	for _, line := range lines {
		nodes := strings.Split(line, "-")
		nodeA, nodeB := nodes[0], nodes[1]
		if connections[nodeA] == nil {
			connections[nodeA] = []string{nodeB}
		} else {
			connections[nodeA] = append(connections[nodeA], nodeB)
		}
		if connections[nodeB] == nil {
			connections[nodeB] = []string{nodeA}
		} else {
			connections[nodeB] = append(connections[nodeB], nodeA)
		}
	}
	return connections
}

type node struct {
	name          string
	previousNodes string
	doubleVisit   bool
}

func day12a(lines []string) {
	connectionTable := makeConnections(lines)
	nodeQueue := []*node{}
	// add first node
	nodeQueue = append(nodeQueue, &node{"start", "start", false})
	counter := 0
	for len(nodeQueue) != 0 {
		newQueue := []*node{}
		for _, n := range nodeQueue {
			connections := connectionTable[n.name]
			for _, c := range connections {
				if c == "end" {
					counter++
					//fmt.Printf("%s - end\n", n.previousNodes)
				} else if strings.ToUpper(c) == c || !strings.Contains(n.previousNodes, c) {
					newPrevious := n.previousNodes + " - " + c
					newNode := node{c, newPrevious, false}
					newQueue = append(newQueue, &newNode)
				}
			}
		}
		nodeQueue = newQueue
	}

	fmt.Printf("Solution for part A: %d\n", counter)
}

func day12b(lines []string) {
	connectionTable := makeConnections(lines)
	nodeQueue := []*node{}
	// add first node
	nodeQueue = append(nodeQueue, &node{"start", "start", false})
	counter := 0
	for len(nodeQueue) != 0 {
		newQueue := []*node{}
		for _, n := range nodeQueue {
			connections := connectionTable[n.name]
			for _, c := range connections {
				doubleVisit := n.doubleVisit
				if c == "start" {
					continue
				} else if c == "end" {
					counter++
					//fmt.Printf("%s,end\n", n.previousNodes)
					continue
				} else if strings.ToUpper(c) != c && strings.Contains(n.previousNodes, c) {
					// skip small caves if already visit and double visit checked
					if doubleVisit {
						continue
					}
					doubleVisit = true
				}
				newPrevious := n.previousNodes + "," + c
				newNode := node{c, newPrevious, doubleVisit}
				newQueue = append(newQueue, &newNode)
			}
		}
		nodeQueue = newQueue
	}

	fmt.Printf("Solution for part B: %d\n", counter)
}
