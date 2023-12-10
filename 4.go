package main

import (
	"regexp"
	"strings"
)

func d4p1(filename string) int {
	cards := read(filename)

	sum := 0
	whitespaces := regexp.MustCompile(`\s+`)
	for _, card := range cards {
		card = whitespaces.ReplaceAllString(card, " ")
		s := strings.Split(card, " | ")
		wins := strings.Split(s[0], " ")
		nums := strings.Split(s[1], " ")

		points := 0
		for _, n := range nums {
			for _, w := range wins {
				if n == w {
					if points == 0 {
						points = 1
					} else {
						points = points * 2
					}
				}
			}
		}
		sum += points
	}
	return sum
}
