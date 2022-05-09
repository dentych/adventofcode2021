package day2

import (
	"strconv"
	"strings"
)

type Command struct {
	Direction string
	Distance  int
}

func Part1(inputLines []string) int {
	var commands []Command
	for _, line := range inputLines {
		split := strings.Split(line, " ")
		cmd := Command{
			Direction: split[0],
			Distance:  convert(split[1]),
		}
		commands = append(commands, cmd)
	}

	x, y := 0, 0
	for _, cmd := range commands {
		switch cmd.Direction {
		case "forward":
			x += cmd.Distance
		case "down":
			y += cmd.Distance
		case "up":
			y -= cmd.Distance
		}
	}

	return x * y
}

func convert(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}