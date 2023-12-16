package main

import (
	"fmt"
	"testing"
)

func BenchmarkFindLargestSumOfEnergizedTiles(b *testing.B) {
	grid := parseFile("task.txt")
	result := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result += findLargestSumOfEnergizedTiles(&grid)
	}
	b.StopTimer()
	fmt.Println(result)
}
