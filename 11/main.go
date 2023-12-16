package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1 := Part1("task.txt")
	fmt.Println("Part 1", part1)
	part2 := Part2("task.txt")
	fmt.Println("Part 2", part2)
}

func Part1(inputPath string) int {
	galaxies := parseFile(inputPath, 2)
	return sumOfGalaxyDistances(galaxies)
}

func Part2(inputPath string) int {
	galaxies := parseFile(inputPath, 1000000)
	return sumOfGalaxyDistances(galaxies)
}

func sumOfGalaxyDistances(galaxies []Position) int {
	sum := 0

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			deltaX := galaxies[j].x - galaxies[i].x
			if deltaX < 0 {
				deltaX = -deltaX
			}

			deltaY := galaxies[j].y - galaxies[i].y
			if deltaY < 0 {
				deltaY = -deltaY
			}

			// println("galaxy1", galaxies[i].x, galaxies[i].y)
			// println("galaxy2", galaxies[j].x, galaxies[j].y)
			// println("delta", deltaX, deltaY)

			sum += deltaX + deltaY
		}
	}

	return sum
}

func parseFile(inputPath string, expansionMultiplier int) []Position {
	file, _ := os.Open(inputPath)
	defer file.Close()

	galaxies := make([]Position, 0)
	rowExpansions := make([]int, 0)
	var colGalaxies []int

	scanner := bufio.NewScanner(file)
	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		if colGalaxies == nil {
			colGalaxies = make([]int, len(line))
		}

		rowIsEmpty := true
		for x, c := range line {
			if c == '#' {
				galaxies = append(galaxies, Position{x, y})
				rowIsEmpty = false
				colGalaxies[x]++
			}
		}

		if rowIsEmpty {
			rowExpansions = append(rowExpansions, y)
		}

		y++
	}

	// println("Galaxies", len(galaxies))

	// for _, y := range rowExpansions {
	// 	// println("Empty row", y)
	// }

	colExpansions := make([]int, 0)

	for x, count := range colGalaxies {
		if count == 0 {
			// println("Empty col", x)
			colExpansions = append(colExpansions, x)
		}
	}

	for g, galaxy := range galaxies {
		expandX := 0
		for i := 0; i < len(colExpansions) && colExpansions[i] < galaxy.x; i++ {
			expandX += expansionMultiplier - 1
		}

		expandY := 0
		for i := 0; i < len(rowExpansions) && rowExpansions[i] < galaxy.y; i++ {
			expandY += expansionMultiplier - 1
		}

		galaxies[g] = Position{x: galaxy.x + expandX, y: galaxy.y + expandY}

		// println("Galaxy", galaxies[g].x, galaxies[g].y)
	}

	return galaxies
}

type Position struct {
	x int
	y int
}
