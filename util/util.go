package util

import (
	"bufio"
	"bytes"
	"io"
	"log"
)

func ReadInput(input io.Reader) []string {
	var output []string
	buf := bufio.NewReader(input)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			log.Fatalln("Failed to read line:", err)
		}
		if bytes.Equal(line, []byte("")) {
			break
		}

		output = append(output, string(line))
	}

	return output
}
