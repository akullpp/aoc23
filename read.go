package main

import (
	"bufio"
	"os"
)

// Reads the lines in a file into a slice.
func read(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for i := 0; scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}
	return lines
}
