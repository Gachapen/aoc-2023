package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	part1 := Part1("task.txt")
	fmt.Println("Part 1", part1)
	part2 := Part2("task.txt")
	fmt.Println("Part 2", part2)
}

func Part1(inputPath string) int {
	grid := parseFile(inputPath)
	return findSumOfEnergizedTiles(&grid, 0, 0, Right)
}

func Part2(inputPath string) int {
	grid := parseFile(inputPath)
	return findLargestSumOfEnergizedTiles(&grid)
}

func findLargestSumOfEnergizedTiles(grid *Grid) int {
	largest := 0
	yBottom := grid.height - 1
	xRight := grid.width - 1

	for x := 0; x < grid.width; x++ {
		top := findSumOfEnergizedTiles(grid, x, 0, Down)
		bottom := findSumOfEnergizedTiles(grid, x, yBottom, Up)

		if top > largest {
			largest = top
		}
		if bottom > largest {
			largest = bottom
		}
	}

	for y := 0; y < grid.height; y++ {
		left := findSumOfEnergizedTiles(grid, 0, y, Right)
		right := findSumOfEnergizedTiles(grid, xRight, y, Left)

		if left > largest {
			largest = left
		}
		if right > largest {
			largest = right
		}
	}

	return largest
}

func findSumOfEnergizedTiles(grid *Grid, xStart int, yStart int, directionStart Direction) int {
	resetTiles(grid)

	energizedCount := 0

	beams := make([]Beam, 1)
	beams[0] = Beam{direction: directionStart, x: xStart, y: yStart}

	for len(beams) != 0 {
		// printGrid(&grid, beams)

		currentIndex := len(beams) - 1
		beam := beams[currentIndex]
		beams = beams[:currentIndex]

		tile := getGridCell(grid, beam.x, beam.y)

		willRepeat := slices.Index(tile.beams, beam.direction) != -1
		hasBeenVisited := len(tile.beams) != 0

		if willRepeat {
			continue
		}

		if !hasBeenVisited {
			energizedCount++
		}

		tile.beams = append(tile.beams, beam.direction)

		var newBeams []Beam

		switch tile.content {
		case '.':
			newBeams = []Beam{moveBeamInCurrentDirection(beam)}
		case '/':
			newBeams = []Beam{reflectBeamRight(beam)}
		case '\\':
			newBeams = []Beam{reflectBeamLeft(beam)}
		case '-':
			newBeams = passBeamThroughHorizontalSplitter(beam)
		case '|':
			newBeams = passBeamThroughVerticalSplitter(beam)
		}

		for _, b := range newBeams {
			if !isOutside(b.x, b.y, grid.width, grid.height) {
				beams = append(beams, b)
			}
		}
	}

	return energizedCount
}

func resetTiles(grid *Grid) {
	for i := 0; i < len(grid.tiles); i++ {
		tile := &grid.tiles[i]
		tile.beams = tile.beams[0:0]
	}
}

func printGrid(grid *Grid, beams []Beam) {
	gridToPrint := make([]byte, len(grid.tiles))
	for i := 0; i < len(grid.tiles); i++ {
		tile := grid.tiles[i]
		gridToPrint[i] = tile.content
	}

	for i := 0; i < len(beams); i++ {
		beam := &beams[i]

		var symbol byte
		switch beam.direction {
		case Up:
			symbol = '^'
		case Down:
			symbol = 'v'
		case Right:
			symbol = '>'
		case Left:
			symbol = '<'
		}

		gridToPrint[beam.y*grid.width+beam.x] = symbol
	}

	for y := 0; y < grid.height; y++ {
		index := y * grid.width
		fmt.Println(string(gridToPrint[index : index+grid.width]))
	}

	fmt.Println()
}

func isOutside(x int, y int, width int, height int) bool {
	return x < 0 || y < 0 || x >= width || y >= height
}

func moveBeamInCurrentDirection(beam Beam) Beam {
	switch beam.direction {
	case Up:
		beam.y--
	case Right:
		beam.x++
	case Down:
		beam.y++
	case Left:
		beam.x--
	}

	return beam
}

func reflectBeamRight(beam Beam) Beam {
	switch beam.direction {
	case Up:
		beam.direction = Right
	case Right:
		beam.direction = Up
	case Down:
		beam.direction = Left
	case Left:
		beam.direction = Down
	}

	return moveBeamInCurrentDirection(beam)
}

func reflectBeamLeft(beam Beam) Beam {
	switch beam.direction {
	case Up:
		beam.direction = Left
	case Right:
		beam.direction = Down
	case Down:
		beam.direction = Right
	case Left:
		beam.direction = Up
	}

	return moveBeamInCurrentDirection(beam)
}

func passBeamThroughHorizontalSplitter(beam Beam) []Beam {
	var beams []Beam

	if beam.direction == Right || beam.direction == Left {
		beams = []Beam{beam}
	} else {
		beams = make([]Beam, 2)
		beams[0] = Beam{x: beam.x, y: beam.y, direction: Left}
		beams[1] = Beam{x: beam.x, y: beam.y, direction: Right}
	}

	for i, beam := range beams {
		beams[i] = moveBeamInCurrentDirection(beam)
	}

	return beams
}

func passBeamThroughVerticalSplitter(beam Beam) []Beam {
	var beams []Beam

	if beam.direction == Up || beam.direction == Down {
		beams = []Beam{beam}
	} else {
		beams = make([]Beam, 2)
		beams[0] = Beam{x: beam.x, y: beam.y, direction: Up}
		beams[1] = Beam{x: beam.x, y: beam.y, direction: Down}
	}

	for i, beam := range beams {
		beams[i] = moveBeamInCurrentDirection(beam)
	}

	return beams
}

func parseFile(inputPath string) Grid {
	file, _ := os.Open(inputPath)
	defer file.Close()

	tiles := make([]Tile, 0)
	height := 0
	var width int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := []byte(scanner.Text())

		row := make([]Tile, len(line))
		for i, c := range line {
			row[i] = Tile{content: c, beams: make([]Direction, 0)}
		}

		width = len(row)
		height += 1

		tiles = append(tiles, row...)
	}

	return Grid{
		tiles:  tiles,
		width:  width,
		height: height,
	}
}

func getGridCell(grid *Grid, x int, y int) *Tile {
	return &grid.tiles[y*grid.width+x]
}

type Direction uint8

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Tile struct {
	content byte
	beams   []Direction
}

type Grid struct {
	tiles  []Tile
	width  int
	height int
}

type Beam struct {
	direction Direction
	x         int
	y         int
}
