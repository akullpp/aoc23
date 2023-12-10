package main

import (
	"regexp"
	"strconv"
)

func d3p1_check(line string, i_end int, i_start int) bool {
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
	lines := read(filename)

	result := 0
	re := regexp.MustCompile(`(\d+)`)
	for i, line := range lines {
		for _, m := range re.FindAllStringIndex(line, -1) {
			v, _ := strconv.Atoi(line[m[0]:m[1]])
			i_begin := m[0]
			i_end := m[1]

			// Check current line
			if d3p1_check(line, i_end, i_begin) {
				result += v
				continue
			}
			// Check line before
			if i > 0 && d3p1_check(lines[i-1], i_end, i_begin) {
				result += v
				continue
			}
			// Check line after
			if i < len(lines)-1 && d3p1_check(lines[i+1], i_end, i_begin) {
				result += v
				continue
			}
		}
	}
	return result
}

// n_start-1 <= gear_idx && n_end >= gear_idx
//
// -- CASE: BEFORE --
//
// . . . * . . .
// . . 1 . . . .
// 3
// 2-1 3
// 1 <= 3 && 3 >= 3 => true
//
// -- CASE: AFTER --
//
// . . . * . . .
// . . . . 1 . .
// 3
// 4-1 5
// 3 <= 3 && 5 >= 3 => true
//
// -- CASE: OVERLAP
//
// . . . * . . .
// . . . 1 . . .
// 3
// 3-1 5
// 2 <= 3 && 5 >= 3 => true
func d3p2_check(line string, gear_idx int, ratios *[]int) {
	re := regexp.MustCompile(`(\d+)`)

	for _, m := range re.FindAllStringIndex(line, -1) {
		n_start := m[0]
		n_end := m[1]

		if n_start-1 <= gear_idx && n_end >= gear_idx {
			v, _ := strconv.Atoi(line[n_start:n_end])
			*ratios = append(*ratios, v)
		}
	}
}

func d3p2(filename string) int {
	lines := read(filename)

	result := 0
	re_gear := regexp.MustCompile(`(\*)`)
	for i, line := range lines {
		for _, m_gear := range re_gear.FindAllStringIndex(line, -1) {
			gear_idx := m_gear[0]
			ratios := []int{}

			// Check current line
			d3p2_check(line, gear_idx, &ratios)

			// Check previous line
			if i > 0 {
				d3p2_check(lines[i-1], gear_idx, &ratios)
			}

			// Check next line
			if i < len(lines)-1 {
				d3p2_check(lines[i+1], gear_idx, &ratios)
			}

			if len(ratios) == 2 {
				result += ratios[0] * ratios[1]
			}
		}
	}
	return result
}
