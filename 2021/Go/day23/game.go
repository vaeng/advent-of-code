package main

import binaryheap "github.com/jhm/go-binaryheap/v2"

type game struct {
	// LL, L, AF, AB, AtB, BF, BB, BtC, CF, CB, CtD, DF, DB, R, RR amphipod
	layout [15]amphipod
}

func (g *game) nextGames() {

}

func (g *game) finished() bool {
	if g.layout[AF] == A &&
		g.layout[AB] == A &&
		g.layout[BF] == B &&
		g.layout[BB] == B &&
		g.layout[CF] == C &&
		g.layout[CB] == C &&
		g.layout[DF] == D &&
		g.layout[DB] == D {
		return true
	}
	return false
}

func (g *game) nextMoves() []game {
	gs := g.toGamestate()
	nextGamestates := gs.nextPossibleGamestates()
	nextGames := []game{}
	for _, ngs := range nextGamestates {
		nextGames = append(nextGames, ngs.toGame())
	}
	return nextGames
}

func energyCost(from, to game) uint {
	var fromPos, toPos position
	var content amphipod

	for i := 0; i < 15; i++ {
		if from.layout[i] != to.layout[i] && to.layout[i] == E {
			fromPos = position(i)
			content = from.layout[i]
		}
		if from.layout[i] != to.layout[i] && from.layout[i] == E {
			toPos = position(i)
		}
	}
	d := getDistance(fromPos, toPos)
	return d * energy[content]
}

func (g *game) toGamestate() gamestate {
	gs := gamestate{}
	a, b, c, d := false, false, false, false
	for i := 0; i < 15; i++ {
		if g.layout[i] == A {
			if !a {
				gs.A1 = position(i)
			} else {
				gs.A2 = position(i)
			}
		}
		if g.layout[i] == B {
			if !b {
				gs.B1 = position(i)
			} else {
				gs.B2 = position(i)
			}
		}
		if g.layout[i] == C {
			if !c {
				gs.C1 = position(i)
			} else {
				gs.C2 = position(i)
			}
		}
		if g.layout[i] == D {
			if !d {
				gs.D1 = position(i)
			} else {
				gs.D2 = position(i)
			}
		}
	}
	return gs
}

func (gs *gamestate) toGame() game {
	g := game{layout: [15]amphipod{E, E, E, E, E, E, E, E, E, E, E, E, E, E, E}}
	g.layout[gs.A1] = A
	g.layout[gs.A2] = A
	g.layout[gs.B1] = B
	g.layout[gs.B2] = B
	g.layout[gs.C1] = C
	g.layout[gs.C2] = C
	g.layout[gs.D1] = D
	g.layout[gs.D2] = D

	return g
}

var ggScores = map[game]uint{}
var gfScores = map[game]uint{}

func getgGScore(g game) uint {
	score, exists := ggScores[g]
	if exists {
		return score
	}
	return infinity
}

func setgGScore(g game, score uint) {
	ggScores[g] = score
}

func getgFScore(g game) uint {
	score, exists := gfScores[g]
	if exists {
		return score
	}
	return infinity
}

func setgFScore(g game, score uint) {
	gfScores[g] = score
}

func aStar(start game) uint {
	ggScores = map[game]uint{}
	gfScores = map[game]uint{}
	cameGFrom := map[game]game{}
	current := start
	openSet := binaryheap.New(func(a, b game) bool { return position(getgFScore(a)) < position(getgFScore(b)) })
	openSet.Push(start)

	setgGScore(start, 0)
	setgFScore(start, 0)
	for openSet.Len() > 0 {
		current, _ = openSet.Pop()

		if current.finished() {
			// totalPath = reconstructGame(cameFrom, current)
			return getgGScore(current)
			break
		}
		for _, nextState := range current.nextMoves() {

			tentativeGScore := getgGScore(current) + energyCost(current, nextState)
			if tentativeGScore < getgGScore(nextState) {
				cameGFrom[nextState] = current
				setgGScore(nextState, tentativeGScore)
				setgFScore(nextState, tentativeGScore+0) // h(nextState))
				openSet.Push(nextState)
			}
		}
	}
	panic("should not reach this point")
}
