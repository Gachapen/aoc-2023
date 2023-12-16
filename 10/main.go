package main

import (
	"bufio"
	"errors"
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
	tileMap := parseFile(inputPath)
	return findFarthestDistance(tileMap)
}

func Part2(inputPath string) int {
	tileMap := parseFile(inputPath)
	return findInsideArea(tileMap)
}

func findInsideArea(tiles TileMap) int {
	markPipeLoop(tiles)

	insideCount := 0
	for y := 0; y < tiles.height; y++ {
		inside := false
		for x := 0; x < tiles.width; x++ {
			print(string(getTile(tiles, Position{x: x, y: y})))
			tile := getTile(tiles, Position{x: x, y: y})
			if tile == 'X' {
				inside = !inside
			} else if tile != 'P' && inside {
				insideCount++
			}
		}
		println("")
	}

	return insideCount
}

func markPipeLoop(tiles TileMap) {
	start, err := findStart(tiles)
	if err != nil {
		panic(err)
	}

	startingConnections := findConnectionsFromStart(tiles, start)

	position := startingConnections[0]
	direction := Position{x: position.x - start.x, y: position.y - start.y}
	tile := getTile(tiles, position)
	steps := 1

	for tile != 'S' {
		previousTile := tile
		previousPosition := position

		steps += 1
		position, direction, tile = findNextTile(tiles, position, direction)

		markLoopTile(tiles, previousPosition, previousTile)
	}

	markLoopTile(tiles, position, tile)
}

func markLoopTile(tiles TileMap, position Position, tile byte) {
	var mark byte
	switch tile {
	case '|':
		mark = 'Y'
		break
	case '-':
		mark = 'X'
		break
	case '7':
		mark = '>'
		break
	case 'J':
		mark = '>'
		break
	case 'F':
		mark = '<'
		break
	case 'L':
		mark = '<'
		break
	case 'S':
		// TODO
		mark = '<'
		break
	}

	setTile(tiles, position, mark)
}

func findFarthestDistance(tiles TileMap) int {
	start, err := findStart(tiles)
	if err != nil {
		panic(err)
	}

	startingConnections := findConnectionsFromStart(tiles, start)

	position := startingConnections[0]
	direction := Position{x: position.x - start.x, y: position.y - start.y}
	tile := getTile(tiles, position)
	steps := 1

	for tile != 'S' {
		steps += 1
		position, direction, tile = findNextTile(tiles, position, direction)
	}

	return steps / 2
}

func findNextTile(tiles TileMap, position Position, direction Position) (nextPosition Position, nextDirection Position, nextTile byte) {
	tile := getTile(tiles, position)

	switch tile {
	case '|':
		if direction.y > 0 {
			nextDirection = down()
		} else {
			nextDirection = up()
		}
		break
	case '-':
		if direction.x > 0 {
			nextDirection = right()
		} else {
			nextDirection = left()
		}
		break
	case '7':
		if direction.y < 0 {
			nextDirection = left()
		} else {
			nextDirection = down()
		}
		break
	case 'F':
		if direction.y < 0 {
			nextDirection = right()
		} else {
			nextDirection = down()
		}
		break
	case 'L':
		if direction.y > 0 {
			nextDirection = right()
		} else {
			nextDirection = up()
		}
		break
	case 'J':
		if direction.y > 0 {
			nextDirection = left()
		} else {
			nextDirection = up()
		}
		break
	default:
		panic("Not a pipe")
	}

	nextPosition = Position{x: position.x + nextDirection.x, y: position.y + nextDirection.y}
	nextTile = getTile(tiles, nextPosition)

	return
}

func findStart(tiles TileMap) (position Position, err error) {
	for i, tile := range tiles.tiles {
		if tile == 'S' {
			return Position{
				x: i % tiles.width,
				y: i / tiles.width,
			}, nil
		}
	}

	return Position{}, errors.New("Could not find start")
}

func findConnectionsFromStart(tiles TileMap, start Position) []Position {
	connections := make([]Position, 0)

	up := Position{x: start.x, y: start.y + 1}
	down := Position{x: start.x, y: start.y - 1}
	left := Position{x: start.x - 1, y: start.y}
	right := Position{x: start.x + 1, y: start.y}

	if isTileOfTypes(tiles, up, '|', 'F', '7') {
		connections = append(connections, up)
	}

	if isTileOfTypes(tiles, down, '|', 'L', 'J') {
		connections = append(connections, down)
	}

	if isTileOfTypes(tiles, left, '-', 'F', 'L') {
		connections = append(connections, left)
	}

	if isTileOfTypes(tiles, right, '-', 'J', '7') {
		connections = append(connections, right)
	}

	return connections
}

func isTileOfTypes(tiles TileMap, position Position, tileTypes ...byte) bool {
	tile := getTile(tiles, position)
	if tile == 0 {
		return false
	}

	for _, tileType := range tileTypes {
		if tile == tileType {
			return true
		}
	}

	return false
}

func getTile(tiles TileMap, position Position) byte {
	if position.x >= tiles.width || position.x < 0 || position.y >= tiles.height || position.y < 0 {
		return 0
	}

	i := position.y*tiles.width + position.x
	return tiles.tiles[i]
}

func setTile(tiles TileMap, position Position, tile byte) {
	if position.x >= tiles.width || position.x < 0 || position.y >= tiles.height || position.y < 0 {
		panic("Outside tile map")
	}

	i := position.y*tiles.width + position.x
	tiles.tiles[i] = tile
}

func parseFile(inputPath string) TileMap {
	file, _ := os.Open(inputPath)
	defer file.Close()

	tiles := make([]byte, 0)

	scanner := bufio.NewScanner(file)
	var width int
	height := 0

	for scanner.Scan() {
		line := scanner.Text()
		tiles = append(tiles, []byte(line)...)
		width = len(line)
		height += 1
	}

	return TileMap{
		width:  width,
		height: height,
		tiles:  tiles,
	}
}

func up() Position {
	return Position{x: 0, y: -1}
}

func down() Position {
	return Position{x: 0, y: 1}
}

func right() Position {
	return Position{x: 1, y: 0}
}

func left() Position {
	return Position{x: -1, y: 0}
}

type Position struct {
	x int
	y int
}

type TileMap struct {
	width  int
	height int
	tiles  []byte
}
