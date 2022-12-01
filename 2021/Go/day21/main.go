package main

import (
	"fmt"
	_ "sync"
	"time"
)

func main() {
	myInput := [2]int{4, 1}
	example := [2]int{4, 8}
	fmt.Println(myInput)
	fmt.Println(example)
	lines := myInput

	start := time.Now()
	partA(lines)
	duration := time.Since(start)
	partB(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

var determinisitcDie int = 1

func getDeterNDiesRollResult(n int) int {
	sum := 0
	for n > 0 {
		if determinisitcDie == 101 {
			determinisitcDie = 1
		}
		sum += determinisitcDie
		determinisitcDie++
		n--
	}
	if sum > 101+100+99 {
		panic("something wrong with the dice")
	}
	return sum
}

func newPos(oldPos, diceResult int) int {
	movement := diceResult % 10
	newPos := oldPos + movement
	for newPos >= 11 {
		newPos -= 10
	}
	return newPos
}

func partA(lines [2]int) {
	p1score := 0
	p2score := 0
	p1pos := lines[0]
	p2pos := lines[1]
	diceRolls := 0
	result := 0
	for i := 0; ; i++ {
		// Player 1 rolls 1+2+3 and moves to space 10 for a total score of 10.
		diceResult := getDeterNDiesRollResult(3)
		p1pos = newPos(p1pos, diceResult)
		if p1pos < 0 || p1pos > 10 {
			panic("player1pos error")
		}
		p1score += p1pos
		diceRolls += 3
		if p1score >= 1000 {
			fmt.Printf("Player 1 rolls %d and moves to space %d for a total score of %d.\n", diceResult, p1pos, p1score)
			fmt.Printf("Game ends: Dice Rolls: %d, P1 Score: %d P2 Score: %d\n", diceRolls, p1score, p2score)
			result = diceRolls * p2score
			break
		}
		diceResult = getDeterNDiesRollResult(3)
		p2pos = newPos(p2pos, diceResult)
		if p2pos < 0 || p2pos > 10 {
			println(p2pos)

			panic("player2pos error")
		}
		p2score += p2pos
		diceRolls += 3
		if p2score >= 1000 {
			fmt.Printf("Player 2 rolls %d and moves to space %d for a total score of %d.\n", diceResult, p2pos, p2score)
			fmt.Printf("Game ends: Dice Rolls: %d, P1 Score: %d P2 Score: %d\n", diceRolls, p1score, p2score)
			result = diceRolls * p1score

			break
		}
	}
	fmt.Printf("Solution for part A: %v\n", result)
}

var universesPerDiceResult = map[int]uint64{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1}

type universe struct {
	p1score, p2score, p1pos, p2pos, totalscore int
}

func partB(lines [2]int) {
	limit := 1000
	start1, start2 := lines[0], lines[1]

	start := time.Now()
	p1WinsA, p2WinsA := simulateAllrounds(start1, start2, limit)
	duration := time.Since(start)
	//p1WinsB, p2WinsB := checkAllUniverses(start1, start2, limit)
	duration2 := time.Since(start)

	result := p1WinsA
	if p1WinsA < p2WinsA {
		result = p2WinsA
	}
	fmt.Printf("allRounds: %s, allUniverses: %s\n", duration, duration2-duration)
	fmt.Printf("Solution for part B: %v\n", result)
}

func bruteForce(start1, start2, limit int) (uint64, uint64) {
	var p1Wins uint64 = 0
	var p2Wins uint64 = 0
	universeCounter := map[universe]uint64{}
	startUniverse := universe{0, 0, start1, start2, 0}
	universeCounter[startUniverse] = 1
	for len(universeCounter) > 0 {
		for uni, counter := range universeCounter {
			// handle first player
			for diceResult, addedUnis1 := range universesPerDiceResult {
				pos1 := newPos(uni.p1pos, diceResult)
				p1Score := uni.p1score + pos1
				// check if plaxer 1 has reached the limit
				if p1Score >= limit {
					delete(universeCounter, uni)
					p1Wins += counter + addedUnis1
					continue
				}
				for diceResult, addedUnis2 := range universesPerDiceResult {
					pos2 := newPos(uni.p2pos, diceResult)
					p2score := uni.p2score + pos2
					// check if plaxer 2 has reached the limit
					if p2score >= limit {
						delete(universeCounter, uni)
						p2Wins += counter + addedUnis1*addedUnis2
						continue
					}
					// no win yet
					newUniverse := universe{p1score: p1Score, p2score: p2score, p1pos: pos1, p2pos: pos2}
					universeCounter[newUniverse] += uint64(addedUnis1 * addedUnis2)
				}
			}
		}
	}
	return p1Wins, p2Wins
}

func checkAllUniverses(start1, start2, limit int) (uint64, uint64) {
	var p1Wins uint64 = 0
	var p2Wins uint64 = 0

	// initialize multiverse
	universeCounter := map[universe]uint64{}
	startUniverse := universe{0, 0, start1, start2, 0}
	universeCounter[startUniverse] = 1
	// check universe with ascending totalscore.
	// that way "lower" universes are checked first and will not be visited again
	for totalscore := 0; totalscore < 2*limit-1; totalscore++ {
		// check all possible scores for p1
		for p1score := 0; p1score < limit; p1score++ {
			// scores for p2 can be calculated from previous loops
			p2score := totalscore - p1score
			if p1score > 20 {
				continue
			}
			// check all possible p1 positions
			for p1pos := 1; p1pos < 11; p1pos++ {
				// check all possible p2 positions
				for p2pos := 1; p2pos < 11; p2pos++ {
					// create matching universe
					currentUni := universe{
						totalscore: totalscore,
						p1score:    p1score,
						p2score:    p2score,
						p1pos:      p1pos,
						p2pos:      p2pos}
					counter := universeCounter[currentUni]
					// no further need for the universe:
					delete(universeCounter, currentUni)
					// universe can't be reached if it has no visits yet
					if counter == 0 {
						continue
					}
					// check if player 1 can win
					for diceResult, addedUnis1 := range universesPerDiceResult {
						newP1pos := newPos(p1pos, diceResult)
						newP1Score := p1score + newP1pos
						if newP1Score >= limit {
							// add the universes from the start multiplied by the ones with this future
							p1Wins += counter * addedUnis1
							// check next player 1 dice throw possibility
							continue
						}
						// check if player 2 can win
						for diceResult, addedUnis2 := range universesPerDiceResult {
							newP2pos := newPos(p2pos, diceResult)
							newP2Score := p2score + newP2pos
							if newP2Score >= limit {
								p2Wins += counter * addedUnis1 * addedUnis2
								// check another throw of player 2
								continue
							}
							// no player has won
							nextUni := universe{
								totalscore: newP1Score + newP2Score,
								p1score:    newP1Score,
								p2score:    newP2Score,
								p1pos:      newP1pos,
								p2pos:      newP2pos}
							// write result to the universeCounter for later calcualtions
							universeCounter[nextUni] += counter * addedUnis1 * addedUnis2
							//checked all player 2 throw possibilites
						}
						//checked all player 1 throw possibilites
					}
					// checked all p2 positions
				}
				//checked all p1 positions
			}
			//checked all p1 scores
		}
		//checked all totalscores
	}

	return p1Wins, p2Wins
}

func simulateAllrounds(start1, start2, limit int) (uint64, uint64) {
	var p1Wins uint64 = 0
	var p2Wins uint64 = 0

	// initialize multiverse
	universeCounter := map[universe]uint64{}
	startUniverse := universe{
		totalscore: 0,
		p1score:    0,
		p2score:    0,
		p1pos:      start1,
		p2pos:      start2}
	universeCounter[startUniverse] = 1
	activeUniverses := 1
	for activeUniverses > 0 {
		//fmt.Printf("Participating Universes: %d\n", len(universeCounter))
		activeUniverses = 0
		newUniverseCounter := map[universe]uint64{}
		for uni, counter := range universeCounter {
			if counter == 0 {
				continue
			}
			activeUniverses++
			for diceResult, multi1 := range universesPerDiceResult {
				newP1pos := newPos(uni.p1pos, diceResult)
				newP1Score := uni.p1score + newP1pos
				if newP1Score >= limit {
					p1Wins += counter * multi1
					universeCounter[uni] = 0
					continue
				}
				for diceResult, multi2 := range universesPerDiceResult {
					newP2pos := newPos(uni.p2pos, diceResult)
					newP2Score := uni.p2score + newP2pos
					if newP2Score >= limit {
						p2Wins += counter * multi1 * multi2
						universeCounter[uni] = 0
						continue
					}
					nextUniverse := universe{
						totalscore: newP1Score + newP2Score,
						p1score:    newP1Score,
						p2score:    newP2Score,
						p1pos:      newP1pos,
						p2pos:      newP2pos}
					universeCounter[uni] = 0
					newUniverseCounter[nextUniverse] += counter * multi1 * multi2
				}
			}
		}
		universeCounter = newUniverseCounter
	}
	return p1Wins, p2Wins
}
