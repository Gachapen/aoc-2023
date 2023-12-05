package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	result := Part1("example_part1.txt")
	assert.Equal(t, 142, result)
}

func TestPart2(t *testing.T) {
	result := Part2("example_part2.txt")
	assert.Equal(t, 281, result)
}
