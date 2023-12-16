package main

import (
	"fmt"
)

func main() {
	part1 := part1("task.txt")
	fmt.Println("Part 1", part1)
	part2 := part2("task.txt")
	fmt.Println("Part 2", part2)
}

func part1(inputPath string) int {
	values := ParsePuzzleInputFile(inputPath)
	return FindSumOfHashes(values)
}

func part2(inputPath string) int {
	values := ParsePuzzleInputFile(inputPath)
	return FindFocusPower(values)
}
