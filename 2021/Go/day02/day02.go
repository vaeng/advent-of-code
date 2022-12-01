package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day02\\02.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	depth := 0
	horiz := 0
	aim := 0
	depth2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := strings.Fields(scanner.Text())
		num, _ := strconv.Atoi(v[1])
		switch v[0] {
		case "forward":
			horiz += num
			depth2 += num * aim
		case "up":
			depth -= num
			aim -= num
		case "down":
			depth += num
			aim += num
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(depth * horiz)

	fmt.Println(depth2 * horiz)
}
