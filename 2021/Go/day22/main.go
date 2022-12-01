package main

import (
	"aoc2021/util"
	"fmt"
	"time"
)

func main() {
	lines := util.GetLines(".\\day22\\in")
	start := time.Now()
	partA(lines)
	duration := time.Since(start)
	partB(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func parseLine(s string) rebootStep {
	// on x=-1..45,y=-26..27,z=-15..32
	// off x=-49..-34,y=-36..-20,z=19..28
	on := false
	var xMin, xMax, yMin, yMax, zMin, zMax int
	pointer := 6
	if s[1] == 'n' {
		on = true
		pointer = 5
	}
	n, err := fmt.Sscanf(s[pointer:], "%d..%d,y=%d..%d,z=%d..%d", &xMin, &xMax, &yMin, &yMax, &zMin, &zMax)
	if n != 6 || err != nil {
		fmt.Println(s[pointer:], on, xMin, xMax, yMin, yMax, zMin, zMax)
		panic(err)
	}
	return rebootStep{on, cuboidRange{xMin, xMax}, cuboidRange{yMin, yMax}, cuboidRange{zMin, zMax}}
}

type cuboidRange struct {
	min, max int
}

type rebootStep struct {
	on      bool
	x, y, z cuboidRange
}

func getRangeOverlap(r1, r2 cuboidRange) (bool, cuboidRange) {
	r := r1
	if r2.min > r1.min {
		r.min = r2.min
	}
	if r1.max > r2.max {
		r.max = r2.max
	}
	if r.min > r.max {
		return false, r
	}
	return true, r
}

func convertRangeForZeroIndex(r cuboidRange, negLimit, posLimit int) (bool, cuboidRange) {
	if r.max < negLimit {
		return false, r
	}
	if r.min > posLimit {
		return false, r
	}
	if r.min < negLimit {
		r.min = negLimit
	}
	if r.max > posLimit {
		r.max = posLimit
	}
	// with limits -10, +100
	// -50, 50 should be 0, 50 - (-10)
	// -2, 2 should be mapped to 0, 4
	if r.min < negLimit {
		r.min = 0
	} else {
		r.min = r.min - negLimit
	}
	if r.max > posLimit {
		r.max = posLimit - negLimit
	} else {
		r.max = r.max - negLimit
	}
	return true, r
}

func switchCuboid(reactor [][][]bool, rs rebootStep) [][][]bool {
	xInsideLimits, zeroIndexedX := convertRangeForZeroIndex(rs.x, -50, 50)
	yInsideLimits, zeroIndexedY := convertRangeForZeroIndex(rs.y, -50, 50)
	zInsideLimits, zeroIndexedZ := convertRangeForZeroIndex(rs.z, -50, 50)
	if !xInsideLimits || !yInsideLimits || !zInsideLimits {
		//fmt.Printf("Instructions %v is out of range x: %v, y%v, z%v\n", rs, zeroIndexedX, zeroIndexedY, zeroIndexedZ)
		return reactor
	}
	//fmt.Printf("zero indexed instruction from %v is x: %v, y%v, z%v\n", rs, zeroIndexedX, zeroIndexedY, zeroIndexedZ)
	for x := zeroIndexedX.min; x <= zeroIndexedX.max; x++ {
		for y := zeroIndexedY.min; y <= zeroIndexedY.max; y++ {
			for z := zeroIndexedZ.min; z <= zeroIndexedZ.max; z++ {
				//fmt.Printf("Accessing x: %d y: %d z: %d\n", x, y, z)
				reactor[x][y][z] = rs.on
			}
		}
	}
	return reactor
}

func cubesOn(reactor [][][]bool) uint64 {
	var count uint64
	reading := 0
	for x, xRange := range reactor {
		for y, yRange := range xRange {
			for z := range yRange {
				if reactor[x][y][z] {
					count++
				}
				reading++
			}
		}
	}
	//fmt.Printf("Read %d cubes, %d were on\n", reading, count)
	return count
}

func initializeReactor(posLimit, negLimit int) [][][]bool {
	// initilaize reactor cubes
	reactor := [][][]bool{}
	for x := negLimit; x <= posLimit; x++ {
		xRange := [][]bool{}
		for y := negLimit; y <= posLimit; y++ {
			yRange := []bool{}
			for z := negLimit; z <= posLimit; z++ {
				yRange = append(yRange, false)
			}
			xRange = append(xRange, yRange)
		}
		reactor = append(reactor, xRange)
	}
	//fmt.Printf("Initialized reactor with %d, %d (x 0-%d, y 0-%d, z 0-%d)\n", negLimit, posLimit, len(reactor), len(reactor[0]), len(reactor[0][0]))
	return reactor
}

func partA(lines []string) {
	// set limits
	posLimit := 50
	negLimit := -50

	// read Steps
	rebootSteps := []rebootStep{}
	for _, line := range lines {
		rebootSteps = append(rebootSteps, parseLine(line))
	}

	// set up reactor
	reactor := initializeReactor(posLimit, negLimit)
	// go through reboot routine
	var counter uint64
	for _, rebootStep := range rebootSteps {
		reactor = switchCuboid(reactor, rebootStep)
		counter = +cubesOn(reactor)
	}

	fmt.Printf("Solution for part A: %v\n", counter)
}

func subtractCuboid(base, substract rebootStep) []rebootStep {
	if false { // !base.on || substract.on {
		fmt.Printf("Base is %v, Subtract is %v\n", base.on, substract.on)
		panic("can only substract off from on")
	}
	returnSteps := []rebootStep{}
	// check if there is an overlap
	xOverlap, yOverlap, zOverlap := true, true, true
	xOverlap, substract.x = getRangeOverlap(base.x, substract.x)
	yOverlap, substract.y = getRangeOverlap(base.y, substract.y)
	zOverlap, substract.z = getRangeOverlap(base.z, substract.z)
	if !xOverlap || !yOverlap || !zOverlap {
		return append(returnSteps, base)
	}

	// check top block, where base.x.min < substract.x.min
	if base.x.min < substract.x.min {
		returnSteps = append(returnSteps, rebootStep{true, cuboidRange{base.x.min, substract.x.min - 1}, base.y, base.z})
	}
	// check bottom block, where base.x.max > substract.x.max
	if base.x.max > substract.x.max {
		returnSteps = append(returnSteps, rebootStep{true, cuboidRange{substract.x.max + 1, base.x.max}, base.y, base.z})
	}
	// check back block, where base.y.min < substract.y.min
	if base.y.min < substract.y.min {
		newXrange := substract.x
		newYrange := cuboidRange{base.y.min, substract.y.min - 1}
		newZrange := base.z
		returnSteps = append(returnSteps, rebootStep{true, newXrange, newYrange, newZrange})
	}
	// check front block, where base.y.max > substract.y.max
	if base.y.max > substract.y.max {
		newXrange := substract.x
		newYrange := cuboidRange{substract.y.max + 1, base.y.max}
		newZrange := base.z
		returnSteps = append(returnSteps, rebootStep{true, newXrange, newYrange, newZrange})
	}
	// check left side block, where base.z.min < substract.z.min
	if base.z.min < substract.z.min {
		newXrange := substract.x
		newYrange := substract.y
		newZrange := cuboidRange{base.z.min, substract.z.min - 1}
		returnSteps = append(returnSteps, rebootStep{true, newXrange, newYrange, newZrange})
	}

	// check right side block, where base.z.max > substract.z.max
	if base.z.max > substract.z.max {
		newXrange := substract.x
		newYrange := substract.y
		newZrange := cuboidRange{substract.z.max + 1, base.z.max}
		returnSteps = append(returnSteps, rebootStep{true, newXrange, newYrange, newZrange})
	}

	return returnSteps
}

func addCuboid(base, addition rebootStep) []rebootStep {
	if false { //|| !base.on || !addition.on {
		fmt.Printf("Base is %v, Addition is %v\n", base.on, addition.on)
		panic("can only substract off from on")
	}
	returnSteps := []rebootStep{}

	// check if there is an overlap
	xOverlap, yOverlap, zOverlap := true, true, true
	xOverlap, _ = getRangeOverlap(base.x, addition.x)
	yOverlap, _ = getRangeOverlap(base.y, addition.y)
	zOverlap, _ = getRangeOverlap(base.z, addition.z)
	if !xOverlap && !yOverlap && !zOverlap {
		fmt.Println("No overlap, both separate")
		returnSteps = append(returnSteps, base)
		return append(returnSteps, addition)
	}
	// check if base is already included, to avoid fragmentation
	if base.x.min > addition.x.min &&
		base.x.max < addition.x.min &&
		base.y.min > addition.y.min &&
		base.y.max < addition.y.min &&
		base.z.min > addition.z.min &&
		base.z.max < addition.z.min {
		fmt.Println("base is part of addition")
		return append(returnSteps, addition)
	}

	// include the base first
	returnSteps = append(returnSteps, base)

	// check if addition is already included in base, to avoid fragmentation
	if addition.x.min > base.x.min &&
		addition.x.max < base.x.min &&
		addition.y.min > base.y.min &&
		addition.y.max < base.y.min &&
		addition.z.min > base.z.min &&
		addition.z.max < base.z.min {
		fmt.Println("addition is part of base")
		return returnSteps // base is already included, nothing else
	}

	// union is the same as substraction + base (already included above)
	base.on = false // to be able to substract
	returnSteps = append(returnSteps, subtractCuboid(addition, base)...)

	return returnSteps
}

func countCubesFromInstructions(rs []rebootStep) uint64 {
	var counter uint64
	for _, r := range rs {
		if !r.on {
			panic("can't count switched off instructions")
		}
		counter += uint64((r.x.max - r.x.min + 1) * (r.y.max - r.y.min + 1) * (r.z.max - r.z.min + 1))
	}
	return counter
}

func partB(lines []string) {
	// read Steps
	rebootSteps := []rebootStep{}
	for _, line := range lines {
		rebootStep := parseLine(line)
		rebootSteps = append(rebootSteps, rebootStep)
	}

	// start the cuboids
	start := 0
	cuboidList := []rebootStep{}
	for i, rs := range rebootSteps {
		if rs.on {
			start = i + 1
			cuboidList = append(cuboidList, rs)
			break
		}
	}
	for i := start; i < len(rebootSteps); i++ {
		newCuboidList := []rebootStep{}
		if rebootSteps[i].on {
			// addition is adding the element,  but substracting it from the rest
			newCuboidList = append(newCuboidList, rebootSteps[i])
		}
		// substraction
		for _, current := range cuboidList {
			substractedCuboid := subtractCuboid(current, rebootSteps[i])
			newCuboidList = append(newCuboidList, substractedCuboid...)
		}
		cuboidList = newCuboidList
	}
	counter := countCubesFromInstructions(cuboidList)
	fmt.Printf("Solution for part B: %v\n", counter)
}
