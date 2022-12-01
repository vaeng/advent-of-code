package main

import (
	"aoc2021/util"
	"fmt"
	"time"

	binaryheap "github.com/jhm/go-binaryheap/v2"
)

func main() {
	lines := util.GetLines(".\\day23\\example")
	start := time.Now()
	partA(lines)
	duration := time.Since(start)
	partB(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

type gamestate struct {
	A1, A2, B1, B2, C1, C2, D1, D2 position
}

const infinity uint = 4294967294

type position uint8

const (
	LL = iota
	L
	AB
	AF
	AtB
	BB
	BF
	BtC
	CB
	CF
	CtD
	DB
	DF
	R
	RR
)

type distance struct {
	from, to position
}

type coordinate struct {
	x, y uint
}

var coordinates = map[position]coordinate{
	LL:  {0, 0},
	L:   {1, 0},
	AB:  {2, 2},
	AF:  {2, 1},
	AtB: {3, 0},
	BB:  {4, 2},
	BF:  {4, 1},
	BtC: {5, 0},
	CB:  {6, 2},
	CF:  {6, 1},
	CtD: {7, 0},
	DB:  {8, 2},
	DF:  {8, 1},
	R:   {9, 0},
	RR:  {10, 0},
}

var coordinateToPosition = map[coordinate]position{
	{0, 0}:  LL,
	{1, 0}:  L,
	{2, 2}:  AB,
	{2, 1}:  AF,
	{3, 0}:  AtB,
	{4, 2}:  BB,
	{4, 1}:  BF,
	{5, 0}:  BtC,
	{6, 2}:  CB,
	{6, 1}:  CF,
	{7, 0}:  CtD,
	{8, 2}:  DB,
	{8, 1}:  DF,
	{9, 0}:  R,
	{10, 0}: RR,
}

type amphipod uint8

const (
	A = iota
	B
	C
	D
	E // empty
)

var energy = map[amphipod]uint{
	A: 1,
	B: 10,
	C: 100,
	D: 1000,
}

var distances = map[distance]uint{}

func getDistance(from, to position) uint {
	d := distance{from, to}
	distance, cached := distances[d]
	if cached {
		return distance
	}
	// calculate it once, manhatten distance
	fromCoord := coordinates[from]
	toCoord := coordinates[to]
	var xDist, yDist uint
	if fromCoord.x > toCoord.x {
		xDist = fromCoord.x - toCoord.x
	} else {
		xDist = toCoord.x - fromCoord.x
	}
	if fromCoord.y > toCoord.y {
		yDist = fromCoord.y - toCoord.y
	} else {
		yDist = toCoord.y - fromCoord.y
	}
	manhattenDistance := xDist + yDist
	distances[d] = manhattenDistance
	// invert
	d.to = from
	d.from = to
	distances[d] = manhattenDistance
	return manhattenDistance
}

func h(gs gamestate) uint {
	//return 0
	var minimumCost uint
	// A1, A2, B1, B2, C1, C2, D1, D2
	if !(gs.A1 == AF || gs.A1 == AB) {
		d := getDistance(gs.A1, AF)
		minimumCost += d * energy[A]
	}
	if !(gs.A2 == AF || gs.A2 == AB) {
		d := getDistance(gs.A2, AF)
		minimumCost += d * energy[A]
	}
	// B
	if !(gs.B1 == BF || gs.B1 == BB) {
		d := getDistance(gs.B1, BF)
		minimumCost += d * energy[B]
	}
	if !(gs.B2 == BF || gs.B2 == BB) {
		d := getDistance(gs.B2, BF)
		minimumCost += d * energy[B]
	}
	// C
	if !(gs.C1 == CF || gs.C1 == CB) {
		d := getDistance(gs.C1, CF)
		minimumCost += d * energy[C]
	}
	if !(gs.C2 == CF || gs.C2 == CB) {
		d := getDistance(gs.C2, CF)
		minimumCost += d * energy[C]
	}
	// D
	if !(gs.D1 == DF || gs.D1 == DB) {
		d := getDistance(gs.D1, DF)
		minimumCost += d * energy[D]
	}
	if !(gs.D2 == DF || gs.D2 == DB) {
		d := getDistance(gs.D2, DF)
		minimumCost += d * energy[D]
	}
	return minimumCost
}

func (gs *gamestate) isOccupied(pos position) bool {
	if gs.A1 == pos || gs.A2 == pos ||
		gs.B1 == pos || gs.B2 == pos ||
		gs.C1 == pos || gs.C2 == pos ||
		gs.D1 == pos || gs.D2 == pos {
		return true
	}
	return false
}

func possibleDestinations(pos position, amph amphipod) []position {
	if amph == A && pos == AB {
		return []position{}
	}
	if amph == B && pos == BB {
		return []position{}
	}
	if amph == C && pos == CB {
		return []position{}
	}
	if amph == D && pos == DB {
		return []position{}
	}

	homePos := map[amphipod][]position{
		A: {AF, AB},
		B: {BF, BB},
		C: {CF, CB},
		D: {DF, DB},
	}
	switch pos {
	case LL, L, AtB, BtC, CtD, R, RR:
		return homePos[amph]
	default:
		return []position{LL, L, AtB, BtC, CtD, R, RR}
	}
}

var cameFrom = map[gamestate]gamestate{}

func getWayInPosVisited(from, to position) []position {
	coordinateList := []position{}
	startX := coordinates[from].x
	startY := coordinates[from].y
	endX := coordinates[to].x
	endY := coordinates[to].y
	switch from {
	case LL, L, AtB, BtC, CtD, R, RR:
		for x := endX; x != startX; {
			coord := coordinate{x, startY}
			pos, exists := coordinateToPosition[coord]
			if exists {
				coordinateList = append(coordinateList, pos)
			}
			if startX < endX {
				x--
			} else {
				x++
			}
		}
		switch to {
		case AB, BB, CB, DB:
			coordinateList = append(coordinateList, coordinateToPosition[coordinate{endX, endY - 1}])
		}
		coordinateList = append(coordinateList, to)
	default:
		coordinateList = getWayInPosVisited(to, from)
		// remove last element, which would be start
		coordinateList = coordinateList[:len(coordinateList)-1]
		// add destination
		coordinateList = append(coordinateList, to)
	}
	return coordinateList
}

func (gs *gamestate) isWayPossible(from, to position) bool {
	for _, pos := range getWayInPosVisited(from, to) {
		if gs.isOccupied(pos) {
			return false
		}
	}
	return true
}

func (gs *gamestate) amphipodsDone(pod amphipod) bool {
	switch pod {
	case A:
		if gs.A1 == AF && gs.A2 == AB {
			return true
		} else if gs.A1 == AB && gs.A2 == AF {
			return true
		}
	case B:
		if gs.B1 == BF && gs.B2 == BB {
			return true
		} else if gs.B1 == BB && gs.B2 == BF {
			return true
		}
	case C:
		if gs.C1 == CF && gs.C2 == CB {
			return true
		} else if gs.C1 == CB && gs.C2 == CF {
			return true
		}
	case D:
		if gs.D1 == DF && gs.D2 == DB {
			return true
		} else if gs.D1 == DB && gs.D2 == DF {
			return true
		}
	default:
		panic("invalids amphipod")
	}
	return false
}

func (gs *gamestate) nextPossibleGamestates() []gamestate {
	nextGamestates := []gamestate{}
	// for A
	if !gs.amphipodsDone(A) {
		for _, nextPos := range possibleDestinations(gs.A1, A) {
			if nextPos == AF && gs.A2 != AB {
				continue
			}
			if gs.isWayPossible(gs.A1, nextPos) {
				newGamestate := *gs
				newGamestate.A1 = nextPos
				nextGamestates = append(nextGamestates, newGamestate)
			}
		}
		for _, nextPos := range possibleDestinations(gs.A2, A) {
			if nextPos == AF && gs.A1 != AB {
				continue
			}
			if gs.isWayPossible(gs.A2, nextPos) {
				newGamestate := *gs
				newGamestate.A2 = nextPos
				nextGamestates = append(nextGamestates, newGamestate)
			}
		}
	}

	// For B
	if !(gs.amphipodsDone(B)) {
		for _, nextPos := range possibleDestinations(gs.B1, B) {
			if nextPos == BF && gs.B2 != BB {
				continue
			}
			if gs.isWayPossible(gs.B1, nextPos) {
				newGamestate := *gs
				newGamestate.B1 = nextPos
				nextGamestates = append(nextGamestates, newGamestate)
			}
		}
		for _, nextPos := range possibleDestinations(gs.B2, B) {
			if nextPos == BF && gs.B1 != BB {
				continue
			}
			if gs.isWayPossible(gs.B2, nextPos) {
				newGamestate := *gs
				newGamestate.B2 = nextPos
				nextGamestates = append(nextGamestates, newGamestate)
			}
		}
	}
	// For C
	if !(gs.amphipodsDone(C)) {
		for _, nextPos := range possibleDestinations(gs.C1, C) {
			if nextPos == CF && gs.C2 != CB {
				continue
			}
			if gs.isWayPossible(gs.C1, nextPos) {
				newGamestate := *gs
				newGamestate.C1 = nextPos
				nextGamestates = append(nextGamestates, newGamestate)
			}
		}
		for _, nextPos := range possibleDestinations(gs.C2, C) {
			if nextPos == CF && gs.C1 != CB {
				continue
			}
			if gs.isWayPossible(gs.C2, nextPos) {
				newGamestate := *gs
				newGamestate.C2 = nextPos
				nextGamestates = append(nextGamestates, newGamestate)
			}
		}
	}
	// For D
	if !(gs.amphipodsDone(D)) {
		for _, nextPos := range possibleDestinations(gs.D1, D) {
			if nextPos == DF && gs.D2 != DB {
				continue
			}
			if gs.isWayPossible(gs.D1, nextPos) {
				newGamestate := *gs
				newGamestate.D1 = nextPos
				nextGamestates = append(nextGamestates, newGamestate)
			}
		}
		for _, nextPos := range possibleDestinations(gs.D2, D) {
			if nextPos == DF && gs.D1 != DB {
				continue
			}
			if gs.isWayPossible(gs.D2, nextPos) {
				newGamestate := *gs
				newGamestate.D2 = nextPos
				nextGamestates = append(nextGamestates, newGamestate)
			}
		}
	}
	return nextGamestates
}

func (gs *gamestate) stringPosition(pos position) string {
	if gs.A1 == pos || gs.A2 == pos {
		return "A"
	}
	if gs.B1 == pos || gs.B2 == pos {
		return "B"
	}
	if gs.C1 == pos || gs.C2 == pos {
		return "C"
	}
	if gs.D1 == pos || gs.D2 == pos {
		return "D"
	}
	return "."
}

func (gs *gamestate) print() {
	sll := gs.stringPosition(LL)
	sl := gs.stringPosition(L)
	satb := gs.stringPosition(AtB)
	sbtc := gs.stringPosition(BtC)
	sctd := gs.stringPosition(CtD)
	sr := gs.stringPosition(R)
	srr := gs.stringPosition(RR)
	saf := gs.stringPosition(AF)
	sab := gs.stringPosition(AB)
	sbf := gs.stringPosition(BF)
	sbb := gs.stringPosition(BB)
	scf := gs.stringPosition(CF)
	scb := gs.stringPosition(CB)
	sdb := gs.stringPosition(DB)
	sdf := gs.stringPosition(DF)

	s := "#############\n"
	s += fmt.Sprintf("#%s%s.%s.%s.%s.%s%s#\n", sll, sl, satb, sbtc, sctd, sr, srr)
	s += fmt.Sprintf("###%s#%s#%s#%s###\n", saf, sbf, scf, sdf)
	s += fmt.Sprintf("  #%s#%s#%s#%s#\n", sab, sbb, scb, sdb)
	s += "  #########\n\n"
	print(s)
}

func normalize(gs gamestate) gamestate {
	if gs.A1 > gs.A2 {
		tmp := gs.A1
		gs.A1 = gs.A2
		gs.A2 = tmp
	}
	if gs.B1 > gs.B2 {
		tmp := gs.B1
		gs.B1 = gs.B2
		gs.B2 = tmp
	}
	if gs.C1 > gs.C2 {
		tmp := gs.C1
		gs.C1 = gs.C2
		gs.C2 = tmp
	}
	if gs.D1 > gs.D2 {
		tmp := gs.D1
		gs.D1 = gs.D2
		gs.D2 = tmp
	}
	return gs
}

func (gs *gamestate) isComplete() bool {
	if !(gs.A1 == AF || gs.A1 == AB) {
		return false
	}
	if !(gs.A2 == AF || gs.A2 == AB) {
		return false
	}
	if !(gs.B1 == BF || gs.B1 == BB) {
		return false
	}
	if !(gs.B2 == BF || gs.B2 == BB) {
		return false
	}
	if !(gs.C1 == CF || gs.C1 == CB) {
		return false
	}
	if !(gs.C2 == CF || gs.C2 == CB) {
		return false
	}
	if !(gs.D1 == DF || gs.D1 == DB) {
		return false
	}
	if !(gs.D2 == DF || gs.D2 == DB) {
		return false
	}
	return true
}

func aStar2(start gamestate) uint {
	//start.normalize()
	current := start
	openSet := binaryheap.New(func(a, b gamestate) bool { return position(a.getFScore()) < position(b.getFScore()) })
	openSet.Push(start)

	start.setGScore(0)
	start.setFScore(h(start))
	for openSet.Len() > 0 {
		current, _ = openSet.Pop()
		//current.normalize()
		if current.isComplete() {
			// totalPath = reconstructGame(cameFrom, current)
			break
		}
		for _, nextState := range current.nextPossibleGamestates() {
			//nextState.normalize()
			tentativeGScore := current.getGScore() + moveCost(current, nextState)
			if tentativeGScore < nextState.getGScore() {
				cameFrom[nextState] = current
				nextState.setGScore(tentativeGScore)
				nextState.setFScore(tentativeGScore + h(nextState))
				openSet.Push(nextState)
			}
		}
	}
	goal := gamestate{AF, AB, BF, BB, CF, CB, DF, DB}
	cost := goal.getFScore()
	return cost
}

var gScores = map[gamestate]uint{}

func (gs *gamestate) getGScore() uint {
	n := normalize(*gs)
	score, exists := gScores[n]
	if exists {
		return score
	}
	return infinity
}

func (gs *gamestate) setGScore(score uint) {
	n := normalize(*gs)
	gScores[n] = score
}

var fScores = map[gamestate]uint{}

func (gs *gamestate) getFScore() uint {
	n := normalize(*gs)
	score, exists := fScores[n]
	if exists {
		return score
	}
	return infinity
}

func (gs *gamestate) setFScore(score uint) {
	n := normalize(*gs)
	fScores[n] = score
}

func reconstructGame(cameFrom map[gamestate]gamestate, current gamestate) []gamestate {
	totalPath := []gamestate{current}
	previousExists := true
	previous := gamestate{}
	for previousExists {
		previous, previousExists = cameFrom[current]
		if previousExists {
			totalPath = append(totalPath, previous)
			current = previous
		}
	}
	return totalPath
}

func (gs *gamestate) getAllPermutations() []gamestate {
	allPermutations := []gamestate{*gs}
	// Flip A
	flippedA := *gs
	flippedA.A1, flippedA.A2 = gs.A2, gs.A1
	allPermutations = append(allPermutations, flippedA)
	// flip B
	withoutB := allPermutations[:]
	for _, base := range withoutB {
		flipped := base
		flipped.B1, flipped.B2 = base.B2, base.B1
		allPermutations = append(allPermutations, flipped)
	}
	// flip C
	withoutC := allPermutations[:]
	for _, base := range withoutC {
		flipped := base
		flipped.C1, flipped.C2 = base.C2, base.C1
		allPermutations = append(allPermutations, flipped)
	}
	// flip D
	withoutD := allPermutations[:]
	for _, base := range withoutD {
		flipped := base
		flipped.D1, flipped.D2 = base.D2, base.D1
		allPermutations = append(allPermutations, flipped)
	}
	return allPermutations
}

func partA(lines []string) {
	// example
	start := gamestate{AB, DB, AF, CF, BF, CB, BB, DF}
	// LL, L, AF, AB, AtB, BF, BB, BtC, CF, CB, CtD, DF, DB, R, RR
	//start := game{layout: [15]amphipod{E, E, E, A, E, E, E, E, E, E, E, A, E, E, E}}
	//example := start
	// input
	//start = gamestate{AF, BB, AB, CF, BF, DB, CB, DF}

	//totalCost := aStar(start.toGame())
	totalCost := aStar2(start)
	// printPath(totalPath)
	fmt.Printf("Solution for part A: %v\n", totalCost)
}

func printPath(totalPath []gamestate) {
	for i := len(totalPath) - 1; i >= 0; i-- {
		current := totalPath[i]
		step := len(totalPath) - i
		fs := current.getFScore()
		gs := current.getGScore()
		moveCost := moveCost(current, cameFrom[current])
		fmt.Printf("Step: %d GScore: %d, FScore: %d, Movecost: %d\n", step, gs, fs, moveCost)
		current.print()
	}
}

func moveCost(current, nextState gamestate) uint {
	gc := current.toGame()
	gn := nextState.toGame()
	return energyCost(gc, gn)

	if current.A1 != nextState.A1 && current.A1 != nextState.A2 {
		d := getDistance(current.A1, nextState.A1)
		return d * energy[A]
	}
	if current.A2 != nextState.A2 && current.A2 != nextState.A1 {
		d := getDistance(current.A2, nextState.A2)
		return d * energy[A]
	}

	if current.B1 != nextState.B1 && current.B1 != nextState.B2 {
		d := getDistance(current.B1, nextState.B1)
		return d * energy[B]
	}
	if current.B2 != nextState.B2 && current.B2 != nextState.B1 {
		d := getDistance(current.B2, nextState.B2)
		return d * energy[B]
	}

	if current.C1 != nextState.C1 && current.C1 != nextState.C2 {
		d := getDistance(current.C1, nextState.C1)
		return d * energy[C]
	}
	if current.C2 != nextState.C2 && current.C2 != nextState.C1 {
		d := getDistance(current.C2, nextState.C2)
		return d * energy[C]
	}

	if current.D1 != nextState.D1 && current.D1 != nextState.D2 {
		d := getDistance(current.D1, nextState.D1)
		return d * energy[D]
	}
	if current.D2 != nextState.D2 && current.D2 != nextState.D1 {
		d := getDistance(current.D2, nextState.D2)
		return d * energy[D]
	}

	return 0
}

func partB(lines []string) {

	fmt.Printf("Solution for part B: %v\n", 0)
}
