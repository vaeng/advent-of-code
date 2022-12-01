package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"time"
)

func main() {
	lines := util.GetLines(".\\day24\\in")
	start := time.Now()
	partA(lines)
	duration := time.Since(start)
	partB(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func parseInstruction(w, x, y, z *int64, instruction string) {
	op := instruction[:3]
	register := instruction[4]
	var target *int64
	switch string(register) {
	case "x":
		target = x
	case "y":
		target = y
	case "z":
		target = z
	case "w":
		target = w
	default:
		panic("invalid target register")
	}

	if op == "inp" {
		panic("should not be reached")
	}

	source := instruction[6:]
	intLiteral, isNum := strconv.Atoi(source)
	literal := int64(intLiteral)
	// is literal
	if isNum == nil {
		switch op {
		case "add":
			*target = *target + literal
		case "div":
			*target = *target / literal
		case "mod":
			*target = *target % literal
		case "mul":
			*target = *target * literal
		case "eql":
			if literal == *target {
				*target = 1
			} else {
				*target = 0
			}
		default:
			panic("invald operation found")
		}
	} else {
		// not a literal but another register
		var sourceRegister *int64
		switch source {
		case "x":
			sourceRegister = x
		case "y":
			sourceRegister = y
		case "z":
			sourceRegister = z
		case "w":
			sourceRegister = w
		}
		switch op {
		case "add":
			*target = *target + *sourceRegister
		case "div":
			*target = *target / *sourceRegister
		case "mod":
			*target = *target % *sourceRegister
		case "mul":
			*target = *target * *sourceRegister
		case "eql":
			if *target == *sourceRegister {
				*target = 1
			} else {
				*target = 0
			}
		default:
			panic("invald operation found")
		}
	}

}

func parseInstructions(ws []int64, instructions []string) bool {
	var w, x, y, z int64
	index := 0
	for _, instruction := range instructions {
		if string(instruction[0]) == "i" {
			w = ws[index]
			//fmt.Printf("z%%26 == %d\n", z%26)
			// fmt.Printf("Packet %d w: %d, x: %d, y: %d, z: %d\n", index, w, x, y, z)

			index++
		} else {
			if instruction == "div z 1" {
				// println("div 1")
			} else if instruction == "div z 26" {
				// println("div 26")
			} else if instruction != "add x z" && instruction[:6] == "add x " {
				// fmt.Printf("x add %s\n (x %d, y%d, z%d, w%d)\n", instruction[6:], x, y, z, w)
			}
			parseInstruction(&w, &x, &y, &z, instruction)
		}
	}
	if z == 0 {
		return true
	}
	return false
}

func partA(lines []string) {
	ws := []int64{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
outer:
	for w0 := 9; w0 > 0; w0-- {
		t := time.Now()
		fmt.Printf("w0: %d reached at %s\n", w0, t)
		for w1 := 9; w1 > 0; w1-- {
			t := time.Now()
			fmt.Printf("w1: %d reached at %s\n", w1, t)
			for w2 := 9; w2 > 0; w2-- {
				for w3 := 9; w3 > 0; w3-- {
					for w4 := 9; w4 > 0; w4-- {
						for w5 := 9; w5 > 0; w5-- {
							for w6 := 9; w6 > 0; w6-- {
								for w7 := 9; w7 > 0; w7-- {
									for w8 := 9; w8 > 0; w8-- {
										for w9 := 9; w9 > 0; w9-- {
											for w10 := 9; w10 > 0; w10-- {
												for w11 := 9; w11 > 0; w11-- {
													for w12 := 9; w12 > 0; w12-- {
														for w13 := 9; w13 > 0; w13-- {
															ws = []int64{int64(w0), int64(w1), int64(w2), int64(w3), int64(w4), int64(w5), int64(w6), int64(w7), int64(w8), int64(w9), int64(w10), int64(w11), int64(w12), int64(w13)}
															success := parseInstructions(ws, lines)
															if success {
																break outer
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("Solution for part A: %v\n", ws)
}

func partB(lines []string) {

	fmt.Printf("Solution for part B: %v\n", 0)
}
