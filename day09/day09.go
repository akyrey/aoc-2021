package day09

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/akyrey/aoc-2021/utils"
)

func readLine(line string) []int {
	stringValues := strings.Split(line, "")
	values := make([]int, len(stringValues))

	for i, stringValue := range stringValues {
		value, err := strconv.Atoi(stringValue)
		utils.CheckError(err)

		values[i] = value
	}

	return values
}

func readBuffer(line string, matrix [][]int, index int) [][]int {
	return append(matrix, readLine(line))
}

func areAdjacentsHigherThanCurrentValue(value, i int, line []int) bool {
	if i == 0 {
		return value < line[i+1]
	}

	if i == len(line)-1 {
		return value < line[i-1]
	}

	return value < line[i-1] && value < line[i+1]
}

func isCurrentValueLowPoint(value, i int, index int, matrix [][]int) bool {
	currentLine := matrix[index]
	if !areAdjacentsHigherThanCurrentValue(value, i, currentLine) {
		return false
	}

	verticalLine := make([]int, 0)
	for x := range matrix {
		verticalLine = append(verticalLine, matrix[x][i])
	}
	if !areAdjacentsHigherThanCurrentValue(value, index, verticalLine) {
		return false
	}

	return true
}

func calcLineLowPoints(index int, matrix [][]int, lowPoints []int) []int {
	currentLine := matrix[index]
	for i, value := range currentLine {
		if isCurrentValueLowPoint(value, i, index, matrix) {
			lowPoints = append(lowPoints, value)
		}
		// fmt.Printf("Matrix %v, low points %v\n", matrix, lowPoints)
	}

	return lowPoints
}

func findBasin(matrix, basinMatrix [][]int, x, y int) int {
	if matrix[x][y] == 9 || basinMatrix[x][y] == 1 {
		basinMatrix[x][y] = 1
		return 0
	}
	basinMatrix[x][y] = 1

	adjacents := 0
	// Up
	if y != 0 {
		adjacents += findBasin(matrix, basinMatrix, x, y-1)
	}
	// Down
	if y != len(matrix[x])-1 {
		adjacents += findBasin(matrix, basinMatrix, x, y+1)
	}
	// Left
	if x != 0 {
		adjacents += findBasin(matrix, basinMatrix, x-1, y)
	}
	// Right
	if x != len(matrix)-1 {
		adjacents += findBasin(matrix, basinMatrix, x+1, y)
	}

	return 1 + adjacents
}

func findThreeLargestBasins(basins []int) []int {
    sort.Ints(basins)

    return basins[len(basins) - 3:]
}

func readFile(test bool) ([]int, []int) {
	f, err := utils.GetFileToReadFrom(9, test)
	utils.CheckError(err)
	defer f.Close()

	lowPoints := make([]int, 0)
	matrix := make([][]int, 0)
	index := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = readBuffer(line, matrix, index)
		if index > 0 && len(matrix) > 1 {
			lowPoints = calcLineLowPoints(len(matrix)-2, matrix, lowPoints)
		}

		index++
	}
	lowPoints = calcLineLowPoints(len(matrix)-1, matrix, lowPoints)
	basinMatrix := make([][]int, len(matrix))
	for x, line := range matrix {
		basinMatrix[x] = make([]int, len(line))
	}
	basins := make([]int, 0)
	for x, line := range basinMatrix {
		for y, value := range line {
			if value != 1 {
				basins = append(basins, findBasin(matrix, basinMatrix, x, y))
			}
		}
	}
	fmt.Printf("Basins found %v\n", basins)
    threeLargestBasins := findThreeLargestBasins(basins)
	fmt.Printf("Largest basins %v\n", threeLargestBasins)

	return lowPoints, threeLargestBasins
}

func Day09(test bool) {
	lowPoints, largestBasins := readFile(test)

	fmt.Printf("Low points found: %v\n", lowPoints)

	riskLevel := 0
	for _, value := range lowPoints {
		riskLevel += (1 + value)
	}

    basinsMultiplied := 1
    for _, value := range largestBasins {
        basinsMultiplied *= value
    }

	fmt.Printf("Risk level sum %d\n", riskLevel)
	fmt.Printf("Basins multiplied %d\n", basinsMultiplied)
}
