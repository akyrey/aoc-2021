package day09

import (
	"bufio"
	"fmt"
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
	matrix = append(matrix, readLine(line))
	if len(matrix) > 3 {
		return matrix[len(matrix)-3:]
	}

	return matrix
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

func readFile(test bool) []int {
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

	return lowPoints
}

func Day09(test bool) {
	lowPoints := readFile(test)

	fmt.Printf("Low points found: %v\n", lowPoints)

	riskLevel := 0
	for _, value := range lowPoints {
		riskLevel += (1 + value)
	}

	fmt.Printf("Risk level sum %d\n", riskLevel)
}
