package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day01\\01.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	first := true
	last := 0
	counter := 0
	var slider []int
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		if !first {
			if value > last {
				counter++
			}
		} else {
			first = false
		}
		last = value
		slider = append(slider, last)
	}

	var windows []int
	for i := range slider {
		if i < len(slider)-2 {
			sum := slider[i] + slider[i+1] + slider[i+2]
			windows = append(windows, sum)
		}
	}

	counter2 := 0
	last = 0
	for i, v := range windows {
		if i != 0 {
			if v > last {
				counter2++
			}
		}
		last = v
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(counter)
	fmt.Println(counter2)

}
