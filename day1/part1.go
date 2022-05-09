package day1

import (
	"fmt"
	"io"
)

func RunPart1(input io.Reader) int {
	numbers := readInput(input)
	fmt.Println("Done reading lines")

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
