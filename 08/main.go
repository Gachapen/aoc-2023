package main

import (
	"bufio"
	"cmp"
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
	instructions, nodeMap := parseFile(inputPath)
	steps := countStepsToEndNode(instructions, nodeMap)
	return steps
}

func Part2(inputPath string) int {
	instructions, nodeMap := parseFile(inputPath)
	steps := countSimultaniousStepsToEndNodes(instructions, nodeMap)
	return steps
}

func countStepsToEndNode(instructions []Instruction, nodeMap map[string]Node) int {
	instructionsCount := len(instructions)
	currentNodeName := "AAA"
	currentInstructionIndex := 0
	steps := 0

	for currentNodeName != "ZZZ" {
		steps += 1

		node := nodeMap[currentNodeName]
		instruction := instructions[currentInstructionIndex%instructionsCount]

		switch instruction {
		case Left:
			currentNodeName = node.left
		case Right:
			currentNodeName = node.right
		}

		currentInstructionIndex += 1
	}

	return steps
}

func countSimultaniousStepsToEndNodes(instructions []Instruction, nodeMap map[string]Node) int {
	startingNodeNames := make([]string, 0)
	for name := range nodeMap {
		if name[2] == 'A' {
			startingNodeNames = append(startingNodeNames, name)
		}
	}

	instructionsCount := len(instructions)
	stepsToEndPerNode := make([]int, len(startingNodeNames))

	for i, nodeName := range startingNodeNames {
		steps := 0
		currentInstructionIndex := 0

		for nodeName[2] != 'Z' {
			steps += 1

			instruction := instructions[currentInstructionIndex%instructionsCount]
			node := nodeMap[nodeName]
			var nextNode string

			switch instruction {
			case Left:
				nextNode = node.left
			case Right:
				nextNode = node.right
			}

			nodeName = nextNode
			currentInstructionIndex += 1
		}

		stepsToEndPerNode[i] = steps
	}

	slices.SortFunc(stepsToEndPerNode, func(a int, b int) int { return cmp.Compare(b, a) })
	for _, steps := range stepsToEndPerNode {
		println(steps)
	}

	// Assuming all paths loop from the end to the beginning, we ca find the multiple
	// of the longest path that is divisible by all the other path lengths.
	longestStepCount := stepsToEndPerNode[0]
	stepMultiplier := 0
	foundMatch := false
	for !foundMatch {
		stepMultiplier++
		totalLongestSteps := longestStepCount * stepMultiplier

		foundMatch = true
		for i := 1; i < len(stepsToEndPerNode); i++ {
			if totalLongestSteps%stepsToEndPerNode[i] != 0 {
				foundMatch = false
				break
			}
		}
	}

	return longestStepCount * stepMultiplier
}

func allNodesAreEndNodes(nodeNames []string) bool {
	for _, name := range nodeNames {
		if name[2] != 'Z' {
			return false
		}
	}

	return true
}

type Node struct {
	left  string
	right string
}

type Instruction uint8

const (
	Right Instruction = iota
	Left
)

func parseFile(inputPath string) ([]Instruction, map[string]Node) {
	file, _ := os.Open(inputPath)
	defer file.Close()

	instructions := make([]Instruction, 0)
	nodeMap := make(map[string]Node)

	scanner := bufio.NewScanner(file)

	// First line is instructions
	scanner.Scan()
	line := scanner.Text()
	for _, c := range line {
		var instruction Instruction
		switch c {
		case 'R':
			instruction = Right
		case 'L':
			instruction = Left
		}
		instructions = append(instructions, instruction)
	}

	// Skip blank line
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		key := line[0:3]
		left := line[7:10]
		right := line[12:15]
		node := Node{
			right: right,
			left:  left,
		}
		nodeMap[key] = node
	}

	return instructions, nodeMap
}
