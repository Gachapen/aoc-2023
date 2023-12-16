package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	result := Part1("example1.txt")
	assert.Equal(t, 21, result)
}

func TestPart2(t *testing.T) {
	result := Part2("example1.txt")
	assert.Equal(t, 525152, result)
}

func TestFindGroupArrangements(t *testing.T) {
	assert.Equal(t, 1, len(FindGroupArrangements([]byte("?"), 0, 1, false)))
	assert.Equal(t, 1, len(FindGroupArrangements([]byte("#"), 0, 1, false)))
	assert.Equal(t, 9, len(FindGroupArrangements([]byte("?????????"), 0, 1, false)))
	assert.Equal(t, 1, len(FindGroupArrangements([]byte("#########"), 0, 9, false)))
	assert.Equal(t, 9, len(FindGroupArrangements([]byte("?????????.####"), 0, 1, false)))
	assert.Equal(t, 2, len(FindGroupArrangements([]byte("?.????????#???"), 7, 2, true)))
}

func TestFindGroupArrangements1(t *testing.T) {
	result := FindGroupArrangements([]byte("?###????????"), 5, 2, false)
	assert.Equal(t, 6, len(result))
}

func TestFindGroupArrangements2(t *testing.T) {
	result := FindGroupArrangements([]byte("?###????????"), 0, 3, false)
	assert.Equal(t, 1, len(result))
}

func TestFindGroupArrangements3(t *testing.T) {
	result := FindGroupArrangements([]byte("????.######..#####."), 0, 1, false)
	assert.Equal(t, 4, len(result))
}

func TestFindGroupArrangements4(t *testing.T) {
	result := FindGroupArrangements([]byte("????.######..#####."), 2, 6, false)
	assert.Equal(t, 1, len(result))
}

func TestCountArrangements(t *testing.T) {
	result := CountArrangements([]byte("????.######..#####."), []int{1, 6, 5})
	assert.Equal(t, 4, result)
}

func TestCountArrangementsTheDumbWay(t *testing.T) {
	// assert.Equal(t, 4, CountArrangementsTheDumbWay([]byte("????.######..#####."), []int{1, 6, 5}))
	// assert.Equal(t, 3, CountArrangementsTheDumbWay([]byte("??????#???"), []int{7, 1}))
	assert.Equal(t, 34, CountArrangementsTheDumbWay([]byte("?.????????#???"), []int{1, 2, 2}))
}

func TestCountArrangementsTask(t *testing.T) {
	testData := make([]TestData, 0)
	// testData = append(testData, makeTestData("??????#???", []int{7, 1}, 3))
	// testData = append(testData, makeTestData("#?#???##??#.?#?#?#?", []int{3, 3, 1, 7}, 2))
	// testData = append(testData, makeTestData(".??#?????.?.??", []int{5, 1, 1}, 15))
	testData = append(testData, makeTestData("?.????????#???", []int{1, 2, 2}, 34))

	for _, test := range testData {
		result := CountArrangements(test.arrangement, test.groups)
		assert.Equal(t, test.result, result)
	}
}

func makeTestData(arrangement string, groups []int, result int) TestData {
	return TestData{[]byte(arrangement), groups, result}
}

type TestData struct {
	arrangement []byte
	groups      []int
	result      int
}
