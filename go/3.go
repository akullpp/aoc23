package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func check(line string, i_end int, i_start int) bool {
	start := i_start - 1
	if start < 0 {
		start = 0
	}
	end := i_end + 1
	if end > len(line) {
		end = len(line)
	}
	seq := line[start:end]
	return regexp.MustCompile(`[^A-Za-z0-9\.]+`).MatchString(seq)
}

func d3p1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for i := 0; scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}

	result := 0
	re := regexp.MustCompile(`(\d+)`)
	for i, line := range lines {
		for _, m := range re.FindAllStringIndex(line, -1) {
			v, _ := strconv.Atoi(line[m[0]:m[1]])
			i_begin := m[0]
			i_end := m[1]

			// Check current line
			if check(line, i_end, i_begin) {
				result += v
				continue
			}
			// Check line before
			if i > 0 && check(lines[i-1], i_end, i_begin) {
				result += v
				continue
			}
			// Check line after
			if i < len(lines)-1 && check(lines[i+1], i_end, i_begin) {
				result += v
				continue
			}
		}
	}
	return result
}
