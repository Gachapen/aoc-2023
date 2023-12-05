package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1 := Part1("task.txt")
	part2 := Part2("task.txt")
	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func Part1(inputPath string) int {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := []byte(scanner.Text())
		var current byte
		var first byte
		for _, byte := range line {
			if byte >= '0' && byte <= '9' {
				current = byte
				if first == 0 {
					first = byte
				}
			}
		}
		last := current
		number, _ := strconv.Atoi(string([]byte{first, last}))
		sum += number
	}

	return sum
}

// var re = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)

var wordToDigitMap = make(map[string]byte)

func Part2(inputPath string) int {
	wordToDigitMap["one"] = '1'
	wordToDigitMap["two"] = '2'
	wordToDigitMap["three"] = '3'
	wordToDigitMap["four"] = '4'
	wordToDigitMap["five"] = '5'
	wordToDigitMap["six"] = '6'
	wordToDigitMap["seven"] = '7'
	wordToDigitMap["eight"] = '8'
	wordToDigitMap["nine"] = '9'

	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := []byte(scanner.Text())
		sum += Part2Line(line)
	}

	return sum
}

func Part2Line(line []byte) int {
	var currentChar byte
	var firstChar byte
	for index := range line {
		digitChar := getDigit(line[index:])
		if digitChar != 0 {
			currentChar = digitChar
			if firstChar == 0 {
				firstChar = digitChar
			}
		}
	}
	lastChar := currentChar
	number, _ := strconv.Atoi(string([]byte{firstChar, lastChar}))
	return number
}

func getDigit(bytes []byte) byte {
	byte := bytes[0]
	if byte >= '0' && byte <= '9' {
		return byte
	}

	for str, digit := range wordToDigitMap {
		if matchesString(bytes, str) {
			return digit
		}
	}

	return 0
}

func matchesString(first []byte, second string) bool {
	for i, secondByte := range []byte(second) {
		if i >= len(first) || secondByte != first[i] {
			return false
		}
	}

	return true
}

func toDigit(str []byte) byte {
	if len(str) == 1 {
		return str[0]
	}

	switch string(str) {
	case "one":
		return '1'
	case "two":
		return '2'
	case "three":
		return '3'
	case "four":
		return '4'
	case "five":
		return '5'
	case "six":
		return '6'
	case "seven":
		return '7'
	case "eight":
		return '8'
	case "nine":
		return '9'
	default:
		panic("not a digit")
	}
}
