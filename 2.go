package main

import (
	"regexp"
	"strconv"
)

func d2p1(filename string) int {
	lines := read(filename)

	result := 0
	// The games are listed sequentially, so we don't need to parse the ID
	re := regexp.MustCompile(`(\d+)\s(\w+)`)
	for i, line := range lines {
		skip := false
		for _, m := range re.FindAllStringSubmatch(line, -1) {
			color := m[2]
			value, _ := strconv.Atoi(m[1])

			// The cubes are returned after each draw, this cost me so much time...
			if (color == "red" && value > 12) ||
				(color == "green" && value > 13) ||
				(color == "blue" && value > 14) {
				skip = true
				break
			}
		}
		if !skip {
			result += i + 1
		}
	}
	return result
}

func d2p2(filename string) int {
	lines := read(filename)

	result := 0
	re := regexp.MustCompile(`(\d+)\s(\w+)`)
	for _, line := range lines {
		config := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, m := range re.FindAllStringSubmatch(line, -1) {
			color := m[2]
			value, _ := strconv.Atoi(m[1])

			if value > config[color] {
				config[color] = value
			}
		}

		power := 1
		for _, v := range config {
			power *= v
		}
		result += power
	}
	return result
}
