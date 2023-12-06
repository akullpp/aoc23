package d1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Observations:
// - there's at least one number per line
// - if there's only one, it is used twice
func d1p1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	re := regexp.MustCompile("[0-9]")
	for scanner.Scan() {
		line := scanner.Text()
		xs := re.FindAllString(line, -1)
		s := fmt.Sprintf("%s%s", xs[0], xs[len(xs)-1])
		n, _ := strconv.Atoi(s)
		sum += n
	}
	fmt.Println("Sum:", sum)
	return sum
}
