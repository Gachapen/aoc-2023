package main

import (
	"bufio"
	"os"
	"strings"
)

func ParsePuzzleInputFile(inputPath string) [][]byte {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	values := strings.Split(line, ",")
	byteValues := make([][]byte, len(values))
	for i, value := range values {
		byteValues[i] = []byte(value)
	}

	return byteValues
}
