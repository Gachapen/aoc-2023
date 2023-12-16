package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example1(t *testing.T) {
	result := Part1("example1.txt")
	assert.Equal(t, 4, result)
}

func TestPart1Example2(t *testing.T) {
	result := Part1("example2.txt")
	assert.Equal(t, 8, result)
}

func TestPart2Example3(t *testing.T) {
	result := Part2("example3.txt")
	assert.Equal(t, 4, result)
}

func TestPart2Example4(t *testing.T) {
	result := Part2("example4.txt")
	assert.Equal(t, 4, result)
}

func TestPart2Example5(t *testing.T) {
	result := Part2("example5.txt")
	assert.Equal(t, 8, result)
}

func TestPart2Example6(t *testing.T) {
	result := Part2("example6.txt")
	assert.Equal(t, 10, result)
}
