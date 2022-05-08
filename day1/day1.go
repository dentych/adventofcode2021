package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func Run(input io.Reader) int {
	bufReader := bufio.NewReader(input)
	var numbers []int
	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			log.Fatalf("Encountered unexpected error while reading from input reader: %s\n", err)
		}
		if len(line) == 0 {
			break
		}
		number, err := strconv.Atoi(string(line))
		if err != nil {
			log.Fatalf("Failed to convert number '%s': %s", string(line), err)
		}
		numbers = append(numbers, number)
	}
	fmt.Fprintln(os.Stderr, "Done reading lines")

	previous := 0
	increaseCount := -1
	for _, number := range numbers {
		if number > previous {
			increaseCount++
		}
		previous = number
	}

	return increaseCount
}
