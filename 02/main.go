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
}

func Part1(inputPath string) int {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	possibleCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		game_split := strings.Split(line, ": ")
		game_str := game_split[0]
		game_id, _ := strconv.Atoi(game_str[5:])

		sets_str := game_split[1]
		sets := strings.Split(sets_str, "; ")

		for _, set_str := range sets {

		}

		possibleCount += game_id
	}

	return possibleCount
}

func readTo(str string, letter rune) (string, string) {
	for index, current := range str {
		if current == letter {
			return str[0:index], str[index+1:]
		}
	}

	return str, ""
}
