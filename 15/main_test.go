package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	result := part1("example1.txt")
	assert.Equal(t, 1320, result)
}

func TestPart2(t *testing.T) {
	result := part2("example1.txt")
	assert.Equal(t, 145, result)
}

func TestCalculateHash(t *testing.T) {
	assert.Equal(t, 30, calculateHash([]byte("rn=1")))
	assert.Equal(t, 253, calculateHash([]byte("cm-")))
	assert.Equal(t, 9, calculateHash([]byte("ot=9")))
}

func TestParsePuzzleInputFile(t *testing.T) {
	parsed := ParsePuzzleInputFile("example1.txt")
	assert.Equal(t, 11, len(parsed))
	assert.Equal(t, []byte("rn=1"), parsed[0])
	assert.Equal(t, []byte("ot=7"), parsed[len(parsed)-1])
}
