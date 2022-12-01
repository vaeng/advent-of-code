package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"time"
)

func main() {
	lines := util.GetLines(".\\day16\\in")
	start := time.Now()
	partA(lines)
	duration := time.Since(start)
	partB(lines)
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}

func partA(lines []string) {
	totalVersionSum := 0
	for _, line := range lines {
		sum := versionSum(line)
		totalVersionSum += sum
		//fmt.Println(line, sum)
	}
	fmt.Printf("Solution for part A: %v\n", totalVersionSum)
}

func versionSum(line string) int {
	binString := hexToBinary(line)
	instructions := translateBinaryToInstructions(binString)
	//fmt.Println("Reading done")
	versionSum := 0
	for len(instructions) > 0 {
		versionSum += instructions[0].version
		if instructions[0].subPackets != nil {
			instructions = append(instructions, instructions[0].subPackets...)
		}
		instructions = instructions[1:]
	}
	return versionSum
}

type instruction struct {
	version    int
	typeID     int
	literal    uint
	subPackets []*instruction
}

func translateBinaryToInstruction(binString string) (*instruction, string) {
	version, _ := strconv.ParseInt(binString[0:3], 2, 32)
	//fmt.Printf("Version: %d (%s)\n", version, binString[0:3])
	typeID, _ := strconv.ParseInt(binString[3:6], 2, 32)
	//fmt.Printf("Type: %d (%s)\n", typeID, binString[3:6])
	switch typeID {
	case 4:
		literal, rest := processLiteral(binString[6:])
		newInstruction := instruction{int(version), int(typeID), literal, nil}
		return &newInstruction, rest
	default:
		lengthType := binString[6] - '0'
		if lengthType == 0 {
			lengthInBits, _ := strconv.ParseInt(binString[7:22], 2, 32)
			//fmt.Printf("Length in Bit Mode (%d)\n", lengthInBits)
			subBinString := binString[22 : 22+lengthInBits]
			subPackets := translateBinaryToInstructions(subBinString)
			newInstruction := instruction{int(version), int(typeID), 0, subPackets}
			return &newInstruction, binString[22+lengthInBits:]
		}
		lengthInPackets, _ := strconv.ParseInt(binString[7:18], 2, 32)
		//fmt.Printf("Length in Packets Mode (%d)\n", lengthInPackets)
		subPackets := []*instruction{}
		binString = binString[18:]
		for i := 0; i < int(lengthInPackets); i++ {
			newInstruction, rest := translateBinaryToInstruction(binString)
			subPackets = append(subPackets, newInstruction)
			binString = rest
		}
		newInstruction := instruction{int(version), int(typeID), 0, subPackets}
		return &newInstruction, binString
	}

}

func translateBinaryToInstructions(binString string) []*instruction {
	instructions := []*instruction{}
	for len(binString) > 8 {
		newInstr, newBinstring := translateBinaryToInstruction(binString)
		instructions = append(instructions, newInstr)
		binString = newBinstring
	}
	return instructions
}

func processLiteral(binString string) (uint, string) {
	binLiteral := ""
	moreToCome := true
	//fmt.Printf("Literals: ")
	for moreToCome {
		moreToCome = binString[0] == '1'
		binLiteral += binString[1:5]
		//fmt.Printf("%s ", binString[0:5])
		binString = binString[5:]
	}
	literal, _ := strconv.ParseInt(binLiteral, 2, 64)
	//fmt.Printf(" -> %d\n", literal)
	return uint(literal), binString
}

func hexToBinary(hexString string) string {
	bs := ""
	for _, r := range hexString {
		switch rune(r) {
		case '0':
			bs += "0000"
		case '1':
			bs += "0001"
		case '2':
			bs += "0010"
		case '3':
			bs += "0011"
		case '4':
			bs += "0100"
		case '5':
			bs += "0101"
		case '6':
			bs += "0110"
		case '7':
			bs += "0111"
		case '8':
			bs += "1000"
		case '9':
			bs += "1001"
		case 'A':
			bs += "1010"
		case 'B':
			bs += "1011"
		case 'C':
			bs += "1100"
		case 'D':
			bs += "1101"
		case 'E':
			bs += "1110"
		case 'F':
			bs += "1111"
		}
	}
	//fmt.Println(bs)
	return string(bs)
}

func partB(lines []string) {
	result := uint(0)
	for _, line := range lines {
		result += evaluateString(line)
		//fmt.Println(line, result)
	}
	fmt.Printf("Solution for part B: %v\n", result)
}

func evaluateString(line string) uint {
	binString := hexToBinary(line)
	instructions := translateBinaryToInstructions(binString)
	if len(instructions) != 1 {
		panic("error, not a single instruction but")
	}
	return evaluateInstruction(instructions[0])
}

func evaluateInstruction(ins *instruction) uint {
	switch ins.typeID {
	case 0:
		return evaluateSum(ins)
	case 1:
		return evaluateProduct(ins)
	case 2:
		return evaluateMinimum(ins)
	case 3:
		return evaluateMaximum(ins)
	case 4:
		return ins.literal
	case 5:
		return evaluateGreaterThan(ins)
	case 6:
		return evaluateLessThan(ins)
	case 7:
		return evaluateEqual(ins)
	default:
		panic("unknown type id")
	}
}

func evaluateEqual(ins *instruction) uint {
	if len(ins.subPackets) != 2 {
		panic("equal needs exactly two sub-packets")
	}
	pack1 := evaluateInstruction(ins.subPackets[0])
	pack2 := evaluateInstruction(ins.subPackets[1])
	//fmt.Printf("(%d == %d)\n", pack1, pack2)
	if pack1 == pack2 {
		return 1
	}
	return 0
}
func evaluateLessThan(ins *instruction) uint {
	if len(ins.subPackets) != 2 {
		panic("less than needs exactly two sub-packets")
	}
	pack1 := evaluateInstruction(ins.subPackets[0])
	pack2 := evaluateInstruction(ins.subPackets[1])
	//fmt.Printf("(%d < %d)\n", pack1, pack2)
	if pack1 < pack2 {
		return 1
	}
	return 0
}

func evaluateGreaterThan(ins *instruction) uint {
	if len(ins.subPackets) != 2 {
		panic("greater than needs exactly two sub-packets")
	}
	pack1 := evaluateInstruction(ins.subPackets[0])
	pack2 := evaluateInstruction(ins.subPackets[1])
	//fmt.Printf("(%d > %d)\n", pack1, pack2)
	if pack1 > pack2 {
		return 1
	}
	return 0
}

func evaluateMaximum(ins *instruction) uint {
	max := evaluateInstruction(ins.subPackets[0])
	for _, subins := range ins.subPackets {
		value := evaluateInstruction(subins)
		//fmt.Printf(" %d", value)
		if value > max {
			max = value
		}
	}
	return max
}

func evaluateMinimum(ins *instruction) uint {
	min := evaluateInstruction(ins.subPackets[0])
	for _, subins := range ins.subPackets {
		value := evaluateInstruction(subins)
		if value < min {
			min = value
		}
	}
	return min
}

func evaluateProduct(ins *instruction) uint {
	prod := uint(1)
	for _, subins := range ins.subPackets {
		prod *= evaluateInstruction(subins)
	}
	return prod
}

func evaluateSum(ins *instruction) uint {
	sum := uint(0)
	for _, subins := range ins.subPackets {
		sum += evaluateInstruction(subins)
	}
	return sum
}
