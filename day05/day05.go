package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
  start Point
  end Point
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func readPoint(text string) (point Point) {
  split := strings.Split(text, ",")

  if len(split) != 2 {
    panic("Wrong point format")
  }

  x, err := strconv.Atoi(split[0])
  check(err)
  y, err := strconv.Atoi(split[1])
  check(err)

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

func readFile(filename string) (lines []Line) {
  f, err := os.Open(filename)
  check(err)
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

  return diagram
}

func main() {
  lines := readFile("input05.txt")
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
