package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type snailNum struct {
	isRegular, complete bool
	regularNum          int
	parent, left, right *snailNum
}

func buildSnailTree(s string, left bool, current *snailNum) {
	switch s[0] {
	case '[': // start new node
		newNode := snailNum{}
		newNode.parent = current
		if left {
			current.left = &newNode
		} else {
			current.right = &newNode
		}
		buildSnailTree(s[1:], true, &newNode)
	case ']': // done with the current node
		current.complete = true
		// check if string is completely consumed
		if len(s) == 1 {
			return // last element checked
		}
		// find previous unfinished nodes
		for current.complete {
			current = current.parent
		}
		if current == nil {
			panic("malformed input string: too many closing parens")
		}
		buildSnailTree(s[1:], false, current)
	case ',':
		buildSnailTree(s[1:], false, current)
	default:
		numString := ""
		offset := 0
	stringloop:
		for i, r := range s {
			switch r {
			case ',':
				if !left {
					panic(", found ] expected")
				}
				offset = i
				break stringloop
			case ']':
				if left {
					//fmt.Println(current)
					//fmt.Println(s)
					panic("] found , expected")
				}
				offset = i
				break stringloop
			default:
				numString += string(r)
			}
		}
		num, _ := strconv.Atoi(numString)
		regular := snailNum{}
		regular.regularNum = num
		regular.parent = current
		regular.isRegular = true
		regular.complete = true
		if left {
			current.left = &regular
		} else {
			current.right = &regular
		}
		buildSnailTree(s[offset:], false, current)
	}
}

func main() {
	lines := util.GetLines(".\\day18\\in")
	start := time.Now()
	partA(lines)
	duration := time.Since(start)
	partB(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func treeString(root *snailNum) string {
	if root == nil {
		return ""
	}

	if root.isRegular {
		return fmt.Sprint(root.regularNum)
	}
	return fmt.Sprintf("[%s,%s]", treeString(root.left), treeString(root.right))
}

func explodeNode(n *snailNum) {
	if !n.left.isRegular && !n.right.isRegular {
		panic("can't explode a pair with nonregular nodes")
	}
	firstleft, _ := findFirstregular(n, true, 0)
	firstright, _ := findFirstregular(n, false, 0)
	if firstleft != nil {
		firstleft.regularNum += n.left.regularNum
	}
	if firstright != nil {
		firstright.regularNum += n.right.regularNum
	}
	*n = snailNum{isRegular: true, complete: true, regularNum: 0, parent: n.parent}
}

func findFirstregular(start *snailNum, searchLeft bool, startDepth int) (*snailNum, int) {
	candidate := start
	depth := startDepth
	// find first left
	if searchLeft {
		for candidate.parent.left == candidate {
			// still on the left path, go up
			candidate = candidate.parent
			depth--
			// reached root:
			if candidate.parent == nil {
				return nil, -1
			}
		}
		candidate = candidate.parent
		depth--
		// take left path
		candidate = candidate.left
		depth++
		// found branching path, going down on the right
		for !candidate.isRegular {
			candidate = candidate.right
			depth++
		}
	} else { // find first right
		for candidate.parent.right == candidate {
			// still on the right path, go up
			candidate = candidate.parent
			depth--
			// reached root:
			if candidate.parent == nil {
				return nil, -1
			}
		}
		candidate = candidate.parent
		depth--
		// take right path
		candidate = candidate.right
		depth++
		// found branching path, going down on the left
		for !candidate.isRegular {
			candidate = candidate.left
			depth++
		}
	}
	return candidate, depth
}

func explode(root *snailNum) bool {
	depth := 0
	limit := 5
	candidate := root
	// go down left
	for !candidate.isRegular {
		candidate = candidate.left
		depth++
	}
	if depth >= limit {
		explodeNode(candidate.parent)
		return true
	}
	for {
		candidate, depth = findFirstregular(candidate, false, depth)
		if candidate == nil {
			break
		}
		if depth >= limit {
			explodeNode(candidate.parent)
			return true
		}
	}
	return false
}

func split(root *snailNum) bool {
	n := findSplit(root)
	if n == nil {
		// nothing to split
		return false
	}
	if !n.isRegular {
		panic("can't split unregular nodes")
	}
	// round down via int div
	newLeft := n.regularNum / 2
	newRight := newLeft
	if newLeft*2 < n.regularNum {
		// need to add 1 to round up
		newRight = newLeft + 1
	}
	*n = snailNum{parent: n.parent, complete: true}
	n.left = &snailNum{parent: n, complete: true, isRegular: true, regularNum: newLeft}
	n.right = &snailNum{parent: n, complete: true, isRegular: true, regularNum: newRight}
	return true
}

func findSplit(root *snailNum) *snailNum {
	candidate := root
	for !candidate.isRegular {
		candidate = candidate.left
	}
	for candidate.regularNum < 10 {
		candidate, _ = findFirstregular(candidate, false, 0)
		if candidate == nil {
			return nil
		}
	}
	if candidate.regularNum > 9 {
		return candidate
	}
	return nil
}

func partA(lines []string) {
	result := &snailNum{}
	for i, line := range lines {
		root := snailNum{}
		if line[0] != '[' {
			panic("malformed string, should start with [")
		}
		buildSnailTree(line[1:], true, &root)
		if i == 0 {
			result = &root
		} else {
			result = addSnailNums(result, &root)
			//println(treeString(result))
			result = reduceSnailNum(result)
			//println(treeString(result))
		}
	}
	result = reduceSnailNum(result)
	lastNode := magnitudeCalculation(result)
	fmt.Printf("Solution for part A: %v\n", lastNode.regularNum)
}

func magnitudeCalculation(n *snailNum) *snailNum {
	if n.isRegular {
		return n
	}
	candidate := n
	for !candidate.isRegular {
		candidate = candidate.left
	}
	for !candidate.parent.right.isRegular {
		candidate, _ = findFirstregular(candidate, false, 0)
		if candidate == nil {
			panic("failed to find magnipair")
		}
	}
	candidate = candidate.parent
	candidate.regularNum = 3*candidate.left.regularNum + 2*candidate.right.regularNum
	candidate.isRegular = true
	candidate.left = nil
	candidate.right = nil
	//fmt.Println(treeString(n))
	return magnitudeCalculation(n)
}

func reduceSnailNum(result *snailNum) *snailNum {
	exploded, splitted := true, true
	for exploded || splitted {
		exploded = explode(result)
		if exploded {
			//fmt.Println("Exploded:")
			//println(treeString(result))
			continue
		}
		splitted = split(result)
		if splitted {
			//fmt.Println("Splitted:")
			//println(treeString(result))
		}
	}
	return result
}

func addSnailNums(a *snailNum, b *snailNum) *snailNum {
	newRoot := snailNum{}
	a.parent = &newRoot
	b.parent = &newRoot
	newRoot.left = a
	newRoot.right = b
	newRoot.complete = true
	return &newRoot
}

func partB(lines []string) {
	maxMagni := 0
	problemSize := len(lines) * (len(lines) - 1)
	c := make(chan int, problemSize)
	var wg sync.WaitGroup
	for i := range lines {
		for j := range lines {
			if i == j {
				continue
			}
			wg.Add(1)
			go getMagni(i, j, lines, &wg, c)
		}
	}
	wg.Wait()
	close(c)
	for currentMagni := range c {
		if currentMagni > maxMagni {
			maxMagni = currentMagni
		}
	}

	fmt.Printf("Solution for part B: %v\n", maxMagni)
}

func getMagni(i, j int, lines []string, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	line := lines[i]
	otherline := lines[j]
	root := snailNum{}
	if line[0] != '[' {
		panic("malformed string, should start with [")
	}
	buildSnailTree(line[1:], true, &root)
	otherRoot := snailNum{}
	if otherline[0] != '[' {
		panic("malformed string, should start with [")
	}
	buildSnailTree(otherline[1:], true, &otherRoot)
	result := addSnailNums(&root, &otherRoot)
	result = reduceSnailNum(result)
	lastNode := magnitudeCalculation(result)
	c <- lastNode.regularNum
}
