package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	data := parseFile(inputPath, false)
	return CountSumOfArrangements(data)
}

func Part2(inputPath string) int {
	data := parseFile(inputPath, true)
	return CountSumOfArrangements(data)
}

func CountArrangementsTheDumbWay(states []byte, groups []int) int {
	stack := make([][]byte, 1)
	stack[0] = states

	permutations := make([][]byte, 0, 1)

	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		i := slices.Index(current, '?')
		if i == -1 {
			permutations = append(permutations, current)
		} else {
			permutation1 := make([]byte, len(current))
			copy(permutation1, current)
			permutation1[i] = '#'

			permutation2 := make([]byte, len(current))
			copy(permutation2, current)
			permutation2[i] = '.'

			stack = append(stack, permutation1, permutation2)
		}
	}

	validPermutationCount := 0

	for _, permutation := range permutations {
		charIndex := 0
		isValidPermutation := true

		for _, groupSize := range groups {
			if charIndex >= len(permutation) {
				isValidPermutation = false
				break
			}

			groupStart := slices.Index(permutation[charIndex:], '#') + charIndex
			groupEnd := groupStart + groupSize

			if groupStart == -1 || groupEnd > len(permutation) || (groupEnd < len(permutation) && permutation[groupEnd] == '#') || !hasOnly(permutation[groupStart:groupEnd], '#') {
				isValidPermutation = false
				break
			}

			charIndex = groupEnd + 1
		}

		if charIndex < len(permutation) && slices.Index(permutation[charIndex:], '#') != -1 {
			isValidPermutation = false
		}

		if isValidPermutation {
			fmt.Println(string(permutation))
			validPermutationCount++
		}
	}

	return validPermutationCount
}

func hasOnly(slice []byte, value byte) bool {
	for _, v := range slice {
		if v != value {
			return false
		}
	}

	return true
}

func CountSumOfArrangements(data []DataRow) int {
	sum := 0
	for _, row := range data {
		a := CountArrangements(row.arrangement, row.groups)
		// println(a)
		// b := CountArrangementsTheDumbWay(row.arrangement, row.groups)
		// println(b)
		sum += a
	}
	return sum
}

func CountArrangements(states []byte, groups []int) int {
	groupArrangements := make([]Range, 1)
	groupArrangements[0] = Range{start: 0, end: -1}

	for g, groupSize := range groups {
		nextGroupArrangements := make([]Range, 0)
		for _, r := range groupArrangements {
			result := FindGroupArrangements(states, r.end+1, groupSize, g == len(groups)-1)
			nextGroupArrangements = append(nextGroupArrangements, result...)

			// for _, a := range result {
			// 	example := make([]byte, len(states))
			// 	copy(example, states)
			// 	for i := 0; i < a.end; i++ {
			// 		if i >= a.start {
			// 			example[i] = '#'
			// 		} else if example[i] == '?' {
			// 			example[i] = '.'
			// 		}
			// 	}
			// 	// fmt.Printf("%v %02d %v\n", g, r.end-1, string(example))
			// }
		}
		groupArrangements = nextGroupArrangements
	}

	numArrangements := len(groupArrangements)
	return numArrangements
}

func FindGroupArrangements(states []byte, startingIndex int, groupSize int, isFinal bool) []Range {
	nextIndexes := make([]Range, 0)

	if startingIndex >= len(states) {
		return nextIndexes
	}

	for i := startingIndex; i < len(states); i++ {
		groupEndIndex := i + groupSize
		if groupEndIndex > len(states) {
			break
		}

		group := states[i:groupEndIndex]
		groupIsPossible := true
		for _, groupChar := range group {
			if groupChar == '.' {
				groupIsPossible = false
				break
			}
		}

		if groupEndIndex < len(states) && (states[groupEndIndex] == '#' || (isFinal && slices.Index(states[groupEndIndex:], '#') != -1)) {
			groupIsPossible = false
		}

		if groupIsPossible {
			nextIndexes = append(nextIndexes, Range{start: i, end: groupEndIndex})
			// fmt.Println(i, groupEndIndex, startingIndex)
		}

		if states[i] == '#' {
			break
		}
	}

	return nextIndexes
}

func parseFile(inputPath string, unfold bool) []DataRow {
	file, _ := os.Open(inputPath)
	defer file.Close()

	data := make([]DataRow, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		arrangement := []byte(split[0])

		groupStrings := strings.Split(split[1], ",")
		groups := make([]int, len(groupStrings))
		for i, str := range groupStrings {
			groups[i], _ = strconv.Atoi(str)
		}

		if unfold {
			unfoldedArrangement := make([]byte, len(arrangement)*5)
			for i := 0; i < 5; i++ {
				startIndex := i * len(arrangement)
				copy(unfoldedArrangement[startIndex:startIndex+len(arrangement)], arrangement)
			}

			unfoldedGroups := make([]int, len(groups)*5)
			for i := 0; i < 5; i++ {
				startIndex := i * len(groups)
				copy(unfoldedGroups[startIndex:startIndex+len(groups)], groups)
			}

			arrangement = unfoldedArrangement
			groups = unfoldedGroups
		}

		data = append(data, DataRow{arrangement, groups})
	}

	return data
}

type DataRow struct {
	arrangement []byte
	groups      []int
}

type Range struct {
	start int
	end   int
}
