package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1 := Part1("task.txt")
	fmt.Println("Part 1", part1)
	part2 := Part2("task.txt")
	fmt.Println("Part 2", part2)
}

func Part1(inputPath string) int {
	histories := parseFile(inputPath)
	return findSumOfNextValues(histories)
}

func Part2(inputPath string) int {
	histories := parseFile(inputPath)
	return findSumOfPreviousValues(histories)
}

func findSumOfNextValues(histories [][]int) int {
	sum := 0
	for _, history := range histories {
		differenceSequences := findDifferenceSequences(history)
		sum += findNextValue(differenceSequences)
	}
	return sum
}

func findSumOfPreviousValues(histories [][]int) int {
	sum := 0
	for _, history := range histories {
		differenceSequences := findDifferenceSequences(history)
		sum += findPreviousValue(differenceSequences)
	}
	return sum
}

func findDifferenceSequences(history []int) [][]int {
	sequences := make([][]int, 1)
	sequences[0] = history
	currentSequence := history

	allZeroes := false
	for !allZeroes {
		allZeroes = true
		nextSequence := make([]int, len(currentSequence)-1)
		for i := 0; i < len(nextSequence); i++ {
			diff := currentSequence[i+1] - currentSequence[i]
			if diff != 0 {
				allZeroes = false
			}
			nextSequence[i] = diff
		}
		sequences = append(sequences, nextSequence)
		currentSequence = nextSequence
	}

	return sequences
}

func findNextValue(sequences [][]int) int {
	// Always starts with zero
	currentPrediction := 0
	for i := len(sequences) - 2; i >= 0; i-- {
		sequence := sequences[i]
		lastValue := sequence[len(sequence)-1]
		currentPrediction = lastValue + currentPrediction
	}

	return currentPrediction
}

func findPreviousValue(sequences [][]int) int {
	// Always starts with zero
	currentPrediction := 0
	for i := len(sequences) - 2; i >= 0; i-- {
		sequence := sequences[i]
		firstValue := sequence[0]
		currentPrediction = firstValue - currentPrediction
	}

	return currentPrediction
}

func parseFile(inputPath string) [][]int {
	file, _ := os.Open(inputPath)
	defer file.Close()

	histories := make([][]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		stringValues := strings.Split(line, " ")
		values := make([]int, len(stringValues))
		for i, stringValue := range stringValues {
			values[i], _ = strconv.Atoi(stringValue)
		}

		histories = append(histories, values)
	}

	return histories
}
