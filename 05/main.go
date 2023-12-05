package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {
	part1 := Part1("task.txt")
	fmt.Println("Part 1", part1)
	part2 := Part2("task.txt")
	fmt.Println("Part 2", part2)
}

func Part1(inputPath string) int {
	seeds, mappings := parseFile(inputPath, false)
	return findLowestMappedValue(seeds, mappings)
}

func Part2(inputPath string) int {
	seeds, mappings := parseFile(inputPath, true)
	return findLowestMappedValue(seeds, mappings)
}

func findLowestMappedValue(seedRanges []IntRange, mappings [][]MappingRange) int {
	ch := make(chan int, len(seedRanges))
	wg := sync.WaitGroup{}

	for _, seedRange := range seedRanges {
		wg.Add(1)
		go findLowestMappedValueForRange(seedRange, mappings, ch, &wg)
	}

	wg.Wait()
	close(ch)

	lowest := math.MaxInt32
	for lowestInRange := range ch {
		if lowestInRange < lowest {
			lowest = lowestInRange
		}
	}

	return lowest
}

func findLowestMappedValueForRange(seedRange IntRange, mappings [][]MappingRange, ch chan int, wg *sync.WaitGroup) {
	lowest := math.MaxInt32

	end := seedRange.start + seedRange.length
	for seed := seedRange.start; seed < end; seed++ {
		currentValue := seed

		for _, mapping := range mappings {
			currentValue = mapSourceToDestination(currentValue, mapping)
		}

		if currentValue < lowest {
			lowest = currentValue
		}
	}

	println("Range", seedRange.start, "done")

	ch <- lowest
	wg.Done()
}

func mapSourceToDestination(source int, mapping []MappingRange) int {
	for _, mappingRange := range mapping {
		if source < mappingRange.sourceStart {
			return source
		}

		if source < mappingRange.sourceEnd {
			return mappingRange.destinationStart + (source - mappingRange.sourceStart)
		}
	}

	return source
}

func parseFile(inputPath string, parseSeedsAsRange bool) ([]IntRange, [][]MappingRange) {
	file, _ := os.Open(inputPath)
	defer file.Close()

	seeds := make([]IntRange, 0)
	mappings := make([][]MappingRange, 0, 7)
	var currentMapping []MappingRange

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds:") {
			seedsStr := line[len("seeds: "):]
			seedsStrList := strings.Split(seedsStr, " ")

			if parseSeedsAsRange {
				i := 0
				for i < len(seedsStrList) {
					seed_start, _ := strconv.Atoi(seedsStrList[i])
					seed_length, _ := strconv.Atoi(seedsStrList[i+1])
					seeds = append(seeds, IntRange{start: seed_start, length: seed_length})
					i += 2
				}
			} else {
				for _, seedStr := range seedsStrList {
					seed, _ := strconv.Atoi(seedStr)
					seeds = append(seeds, IntRange{start: seed, length: 1})
				}
			}
		} else if strings.Contains(line, "map:") {
			currentMapping = make([]MappingRange, 0)
		} else if len(line) == 0 {
			if currentMapping != nil {
				slices.SortFunc(currentMapping, func(a, b MappingRange) int {
					return cmp.Compare(a.sourceStart, b.sourceStart)
				})
				mappings = append(mappings, currentMapping)
			}
		} else {
			valueStrList := strings.Split(line, " ")
			destination, _ := strconv.Atoi(valueStrList[0])
			source, _ := strconv.Atoi(valueStrList[1])
			length, _ := strconv.Atoi(valueStrList[2])
			mappingRange := MappingRange{
				sourceStart:      source,
				sourceEnd:        source + length,
				destinationStart: destination,
			}
			currentMapping = append(currentMapping, mappingRange)
		}
	}

	mappings = append(mappings, currentMapping)
	return seeds, mappings
}

type MappingRange struct {
	sourceStart      int
	sourceEnd        int
	destinationStart int
}

type IntRange struct {
	start  int
	length int
}
