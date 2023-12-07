package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Observations:
// - there's at least one number per line
// - there's never a 0
// - if there's only one number, it is used twice
func d1p1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	re := regexp.MustCompile("[1-9]")
	for scanner.Scan() {
		line := scanner.Text()
		xs := re.FindAllString(line, -1)
		s := fmt.Sprintf("%s%s", xs[0], xs[len(xs)-1])
		n, _ := strconv.Atoi(s)
		sum += n
	}
	return sum
}

// Only difference to d1p1 is line 49-60 and 63-68
func d1p2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	re := regexp.MustCompile("[1-9]")
	// Trick for overlap, e.g. eightwone -> e8twone -> e8t2one -> e8t2o1e
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
	for scanner.Scan() {
		line := scanner.Text()
		before := line
		after := replacer.Replace(before)
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
