package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func d1p1(filename string) int {
	lines := read(filename)

	sum := 0
	// There's never a 0
	re := regexp.MustCompile("[1-9]")
	for _, line := range lines {
		// There's at least one number per line, so no need to check
		xs := re.FindAllString(line, -1)
		// If there's only one number, it is used twice, e.g. 9sixonefour -> 99
		s := fmt.Sprintf("%s%s", xs[0], xs[len(xs)-1])
		n, _ := strconv.Atoi(s)
		sum += n
	}
	return sum
}

func d1p2(filename string) int {
	lines := read(filename)

	sum := 0
	re := regexp.MustCompile("[1-9]")
	// Trick for the overlap, e.g. eightwone -> e8twone -> e8t2one -> e8t2o1e
	replacer := strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9e",
	)
	for _, line := range lines {
		before := line
		after := replacer.Replace(before)
		// Replace until there's nothing more to replace
		for before != after {
			before = after
			after = replacer.Replace(before)
		}
		xs := re.FindAllString(after, -1)
		s := fmt.Sprintf("%s%s", xs[0], xs[len(xs)-1])
		n, _ := strconv.Atoi(s)
		sum += n
	}
	return sum
}
