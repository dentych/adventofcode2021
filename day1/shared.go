package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
)

func readInput(reader io.Reader) []int {
	bufReader := bufio.NewReader(reader)
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
	fmt.Println("Done reading input!")
	return numbers
}
