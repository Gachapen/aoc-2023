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
		// printGrid(grid, beams)

		currentIndex := len(beams) - 1
		beam := &beams[currentIndex]

		tile := getGridCell(grid, beam.x, beam.y)

		hasBeenVisited := tile.visitedDown || tile.visitedLeft || tile.visitedRight || tile.visitedUp

		willRepeat := false
		switch beam.direction {
		case Up:
			willRepeat = tile.visitedUp
			tile.visitedUp = true
		case Down:
			willRepeat = tile.visitedDown
			tile.visitedDown = true
		case Left:
			willRepeat = tile.visitedLeft
			tile.visitedLeft = true
		case Right:
			willRepeat = tile.visitedRight
			tile.visitedRight = true
		}

		if willRepeat {
			beams = beams[:currentIndex]
			continue
		}

		if !hasBeenVisited {
			energizedCount++
		}

		wasSplit := false
		var split1 Beam
		var split2 Beam

		switch tile.content {
		case '.':
			moveBeamInCurrentDirection(beam)
		case '/':
			reflectBeamRight(beam)
		case '\\':
			reflectBeamLeft(beam)
		case '-':
			wasSplit, split1, split2 = passBeamThroughHorizontalSplitter(beam)
		case '|':
			wasSplit, split1, split2 = passBeamThroughVerticalSplitter(beam)
		}

		if wasSplit {
			beams = beams[:currentIndex]
			if !isOutside(split1.x, split1.y, grid.width, grid.height) {
				beams = append(beams, split1)
			}
			if !isOutside(split2.x, split2.y, grid.width, grid.height) {
				beams = append(beams, split2)
			}
		} else if isOutside(beam.x, beam.y, grid.width, grid.height) {
			beams = beams[:currentIndex]
		}
	}

	return energizedCount
}

func isHorizontal(direction Direction) bool {
	return direction == Right || direction == Left
}

func isVertical(direction Direction) bool {
	return direction == Up || direction == Down
}

func resetTiles(grid *Grid) {
	for i := 0; i < len(grid.tiles); i++ {
		tile := &grid.tiles[i]
		tile.visitedUp = false
		tile.visitedDown = false
		tile.visitedRight = false
		tile.visitedLeft = false
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

func moveBeamInCurrentDirection(beam *Beam) {
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
}

func reflectBeamRight(beam *Beam) {
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

	moveBeamInCurrentDirection(beam)
}

func reflectBeamLeft(beam *Beam) {
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

	moveBeamInCurrentDirection(beam)
}

func passBeamThroughHorizontalSplitter(beam *Beam) (bool, Beam, Beam) {
	if beam.direction == Right || beam.direction == Left {
		moveBeamInCurrentDirection(beam)
		return false, Beam{}, Beam{}
	} else {
		beam1 := Beam{x: beam.x, y: beam.y, direction: Left}
		beam2 := Beam{x: beam.x, y: beam.y, direction: Right}
		moveBeamInCurrentDirection(&beam1)
		moveBeamInCurrentDirection(&beam2)

		return true, beam1, beam2
	}
}

func passBeamThroughVerticalSplitter(beam *Beam) (bool, Beam, Beam) {
	if beam.direction == Up || beam.direction == Down {
		moveBeamInCurrentDirection(beam)
		return false, Beam{}, Beam{}
	} else {
		beam1 := Beam{x: beam.x, y: beam.y, direction: Up}
		beam2 := Beam{x: beam.x, y: beam.y, direction: Down}
		moveBeamInCurrentDirection(&beam1)
		moveBeamInCurrentDirection(&beam2)

		return true, beam1, beam2
	}
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
			row[i] = Tile{content: c}
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
	content      byte
	visitedRight bool
	visitedLeft  bool
	visitedUp    bool
	visitedDown  bool
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
