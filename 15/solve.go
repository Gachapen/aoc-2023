package main

import (
	"slices"
	"strconv"
)

func FindSumOfHashes(values [][]byte) int {
	sum := 0
	for _, value := range values {
		sum += calculateHash(value)
	}
	return sum
}

func FindFocusPower(values [][]byte) int {
	boxes := make([][]Lens, 256)

	for _, value := range values {
		instruction := parseInstruction(value)
		boxIndex := calculateHash([]byte(instruction.label))
		switch instruction.operation {
		case OperationRemove:
			boxes[boxIndex] = removeLens(instruction.label, boxes[boxIndex])
		case OperationInsert:
			boxes[boxIndex] = insertLens(instruction.label, instruction.focalLength, boxes[boxIndex])
		}
	}

	return findTotalFocalPower(boxes)
}

func findTotalFocalPower(boxes [][]Lens) int {
	totalFocalPower := 0
	for boxNumber, box := range boxes {
		totalFocalPower += findBoxFocalPower(box, boxNumber)
	}
	return totalFocalPower
}

func findBoxFocalPower(box []Lens, boxNumber int) int {
	boxFocalPower := 0
	for lensIndex, lens := range box {
		boxFocalPower += findLensFocalPower(boxNumber, lensIndex+1, lens.focalLength)
	}
	return boxFocalPower
}

func findLensFocalPower(boxNumber int, lensSlotNumber int, focalLength int) int {
	return (1 + boxNumber) * lensSlotNumber * focalLength
}

func removeLens(label string, box []Lens) []Lens {
	lensIndex := slices.IndexFunc(box, func(lens Lens) bool { return lens.label == label })
	if lensIndex == -1 {
		return box
	}

	for i := lensIndex; i < len(box)-1; i++ {
		box[i] = box[i+1]
	}

	return box[:len(box)-1]
}

func insertLens(label string, focalLength int, box []Lens) []Lens {
	lensIndex := slices.IndexFunc(box, func(lens Lens) bool { return lens.label == label })
	if lensIndex != -1 {
		box[lensIndex] = Lens{label: label, focalLength: focalLength}
		return box
	}

	return append(box, Lens{label: label, focalLength: focalLength})
}

func parseInstruction(value []byte) Instruction {
	operationIndex := slices.Index(value, '=')
	if operationIndex == -1 {
		operationIndex = slices.Index(value, '-')
	}

	label := string(value[:operationIndex])

	var operation Operation
	switch value[operationIndex] {
	case '=':
		operation = OperationInsert
	case '-':
		operation = OperationRemove
	}

	focalLength := 0
	if operationIndex < len(value)-1 {
		focalLength, _ = strconv.Atoi(string(value[operationIndex+1:]))
	}

	return Instruction{label: label, operation: operation, focalLength: focalLength}
}

func calculateHash(value []byte) int {
	hashValue := 0

	for _, c := range value {
		hashValue += int(c)
		hashValue *= 17
		hashValue %= 256
	}

	return hashValue
}

type Instruction struct {
	label       string
	operation   Operation
	focalLength int
}

type Lens struct {
	label       string
	focalLength int
}

type Operation uint8

const (
	OperationRemove Operation = iota
	OperationInsert Operation = iota
)
