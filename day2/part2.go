package day2

import (
	"strings"
)

func Part2(inputLines []string) int {
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
	aim := 0
	for _, cmd := range commands {
		switch cmd.Direction {
		case "forward":
			x += cmd.Distance
			y += aim * cmd.Distance
		case "down":
			aim += cmd.Distance
		case "up":
			aim -= cmd.Distance
		}
	}

	return x * y
}
