package day1

import "io"

func RunPart2(input io.Reader) int {
	numbers := readInput(input)

	previous := 0
	count := -1
	for i := 0; i+2 < len(numbers); i++ {
		slidingWindowSum := numbers[i] + numbers[i+1] + numbers[i+2]
		if slidingWindowSum > previous {
			count++
		}

		previous = slidingWindowSum
	}

	return count
}
