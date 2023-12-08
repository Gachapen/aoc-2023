package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example1(t *testing.T) {
	result := Part1("example1.txt")
	assert.Equal(t, 2, result)
}

func TestPart1Example2(t *testing.T) {
	result := Part1("example2.txt")
	assert.Equal(t, 6, result)
}

func TestPart2(t *testing.T) {
	result := Part2("example3.txt")
	assert.Equal(t, 6, result)
}
