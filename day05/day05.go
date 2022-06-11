package day05

import (
	"bufio"
	"fmt"
	"github.com/akyrey/aoc-2021/utils"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func readPoint(text string) (point Point) {
	split := strings.Split(text, ",")

	if len(split) != 2 {
		panic("Wrong point format")
	}

	x, err := strconv.Atoi(split[0])
	utils.CheckError(err)
	y, err := strconv.Atoi(split[1])
	utils.CheckError(err)

	return Point{x, y}
}

func readLine(line string) (start Point, end Point) {
	split := strings.Split(line, " -> ")

	if len(split) != 2 {
		panic("Wrong line format")
	}

	start = readPoint(split[0])
	end = readPoint(split[1])

	return
}

func readFile(test bool) (lines []Line) {
	f, err := utils.GetFileToReadFrom(5, test)
	utils.CheckError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		start, end := readLine(line)
		lines = append(lines, Line{start, end})
	}

	fmt.Printf("Lines %v\n", lines)

	return
}

func highestPoint(lines []Line) (point Point) {
	for _, line := range lines {
		if point.x < line.start.x {
			point.x = line.start.x
		}
		if point.x < line.end.x {
			point.x = line.end.x
		}
		if point.y < line.start.y {
			point.y = line.start.y
		}
		if point.y < line.end.y {
			point.y = line.end.y
		}
	}

	fmt.Printf("Highest point %v\n", point)
	return
}

func trackLine(line Line, diagram [][]int) [][]int {
	// Horizontal lines
	if line.start.x == line.end.x {
		start := int(math.Min(float64(line.start.y), float64(line.end.y)))
		end := int(math.Max(float64(line.start.y), float64(line.end.y)))

		for i := start; i <= end; i++ {
			diagram[line.start.x][i]++
		}
	}

	// Vertical lines
	if line.start.y == line.end.y {
		start := int(math.Min(float64(line.start.x), float64(line.end.x)))
		end := int(math.Max(float64(line.start.x), float64(line.end.x)))

		for i := start; i <= end; i++ {
			diagram[i][line.start.y]++
		}
	}

	// Diagonal lines
	if line.start.x != line.end.x && line.start.y != line.end.y {
		lenghtX := int(math.Abs(float64(line.start.x) - float64(line.end.x)))
		lenghtY := int(math.Abs(float64(line.start.y) - float64(line.end.y)))

		if lenghtX == lenghtY {
			for row, col := line.start.x, line.start.y; isCycleEnded(row, line.start.x, line.end.x) && isCycleEnded(col, line.start.y, line.end.y); row, col = nextValue(row, line.start.x, line.end.x), nextValue(col, line.start.y, line.end.y) {
				// fmt.Printf("This is a diagonal line, row: %d col: %d\n", row, col)
				diagram[row][col]++
			}
		}
	}

	return diagram
}

func isCycleEnded(current int, start int, end int) bool {
	if start > end {
		return end <= current
	}

	return end >= current
}

func nextValue(current int, start int, end int) int {
	if start > end {
		return current - 1
	}

	return current + 1
}

func Day05(test bool) {
	lines := readFile(test)
	highestPoint := highestPoint(lines)
	diagramLength := Point{highestPoint.x + 1, highestPoint.y + 1}

	diagram := make([][]int, diagramLength.x)
	for i := 0; i < diagramLength.x; i++ {
		diagram[i] = make([]int, diagramLength.y)
	}

	for _, line := range lines {
		diagram = trackLine(line, diagram)
	}

	fmt.Printf("Final diagram: %v\n", diagram)

	count := 0
	for col := 0; col < diagramLength.y; col++ {
		for row := 0; row < diagramLength.x; row++ {
			if diagram[row][col] > 1 {
				count++
			}
		}
	}

	fmt.Printf("Overlapping line points: %d\n", count)
}
