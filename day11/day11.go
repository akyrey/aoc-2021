package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/akyrey/aoc-2021/utils"
)

type Position struct {
	x int
	y int
}

func initMatrix(length int) [][]int {
	matrix := make([][]int, length)
	for i := range matrix {
		matrix[i] = make([]int, length)
	}

	return matrix
}

func buildInitialMatrix(f *os.File) [][]int {
	matrix := initMatrix(10)
	scanner := bufio.NewScanner(f)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		values := strings.Split(line, "")

		for x, value := range values {
			octopus, err := strconv.Atoi(value)
			utils.CheckError(err)

			matrix[y][x] = octopus
		}
	}

	return matrix
}

func initialStepIncrease(matrix [][]int) {
	for y := range matrix {
		for x := range matrix[y] {
			matrix[y][x] += 1
		}
	}
}

/**
* [b, a] since we need to translate [j, i] based on [y, x]
* [-1, -1] [-1, 0] [-1, 1]
* [0, -1]  [0, 0]  [0, 1]
* [1, -1]  [1, 0]  [1, 1]
 */
func increaseAdjacents(matrix [][]int, y, x int) {
	for j := -1; j < 2; j++ {
		for i := -1; i < 2; i++ {
			b := y + j
			a := x + i
			if
			// Don't increase self
			(j != 0 || i != 0) &&
				// Check octs is inside the matrix
				a >= 0 && b >= 0 && b < len(matrix) && a < len(matrix[y]) &&
				// Avoid increasing already flashed octs
				matrix[b][a] != -1 {
				matrix[b][a] += 1
				// fmt.Printf("Increasing %d, %d\n", b, a)
			}
		}
	}
}

func findOctopusesThatWillFlash(matrix [][]int) []Position {
	octs := make([]Position, 0)

	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] > 9 {
				octs = append(octs, Position{x, y})
			}
		}
	}

	return octs
}

func resetFlashedOcts(matrix [][]int) [][]int {
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == -1 {
				matrix[y][x] = 0
			}
		}
	}

	return matrix
}

func performStep(flashes int, matrix [][]int) int {
	initialStepIncrease(matrix)

	for octs := findOctopusesThatWillFlash(matrix); len(octs) > 0; octs = findOctopusesThatWillFlash(matrix) {
		flashes += len(octs)
		for _, pos := range octs {
			// Flash the octopus
			matrix[pos.y][pos.x] = -1
			// Increase all adjacent ones
			increaseAdjacents(matrix, pos.y, pos.x)
		}
	}

	resetFlashedOcts(matrix)

	return flashes
}

func Day11(test bool) {
	f, err := utils.GetFileToReadFrom(11, test)
	utils.CheckError(err)
	defer f.Close()

	matrix := buildInitialMatrix(f)
	steps := 100
	flashes := 0

	fmt.Printf("Initial energy level: %v\n", matrix)

	for step := 0; step < steps; step++ {
		flashes = performStep(flashes, matrix)
		fmt.Printf("Energy level after %d step: %v, flashes performed %d\n", step, matrix, flashes)
	}
}
