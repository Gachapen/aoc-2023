package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	result := Part1("example1.txt")
	assert.Equal(t, 374, result)
}

func TestPart2(t *testing.T) {
	result := Part2("example1.txt")
	assert.Equal(t, 8410, result)
}
