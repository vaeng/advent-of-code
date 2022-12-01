package util

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

// GetLines reads a file in path and returns them in a
// slice. An element for each new line.
func GetLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := scanner.Text()
		lines = append(lines, v)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

// StringInside checks if string a is in string b.
func StringInside(a, b string) bool {
	for _, r := range a {
		if !strings.ContainsRune(b, r) {
			return false
		}
	}
	return true
}

// Taken from: https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// minMaxIntSlice returns the max and min of an int slice.
func MinMaxIntSlice(slice []int) (int, int) {
	min, max := slice[0], slice[0]
	for _, v := range slice {
		switch {
		case v < min:
			min = v
		case v > max:
			max = v
		}
	}
	return min, max
}
