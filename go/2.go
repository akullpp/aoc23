package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Observations:
// Games are listed sequentially, no need to parse the id
// The cubes are returned after each draw...
func d2p1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result := 0
	n := 1
	re := regexp.MustCompile(`(\d+)\s(\w+)`)
	for scanner.Scan() {
		line := scanner.Text()
		skip := false
		for _, m := range re.FindAllStringSubmatch(line, -1) {
			color := m[2]
			value, _ := strconv.Atoi(m[1])

			if (color == "red" && value > 12) ||
				(color == "green" && value > 13) ||
				(color == "blue" && value > 14) {
				skip = true
				break
			}
		}
		if !skip {
			result += n
		}
		n++
	}
	return result
}
