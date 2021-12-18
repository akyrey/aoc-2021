package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func readFile(filename string) []int {
  f, err := os.Open(filename)
  check(err)
  defer f.Close()

  scanner := bufio.NewScanner(f)
  var positions []int

  for scanner.Scan() {
    line := scanner.Text()
    fmt.Printf("Line read: %s\n", line)

    positionsAsStrings := strings.Split(line, ",")
    positions = make([]int, len(positionsAsStrings))

    for i, stringValue := range positionsAsStrings {
      value, err := strconv.Atoi(stringValue)
      check(err)

      positions[i] = value
    }

    fmt.Printf("Initial crab positions: %v\n", positions)
  }

  return positions
}

func calculateMeanValue(positions []int) int {
  sum := 0
  for _, position := range positions {
    sum += position
  }

  mean := sum / len(positions)
  fmt.Printf("Mean %d\n", mean)

  return mean
}

func calculateFuelSpent(positions []int, target int) int {
  fuel := 0

  for _, position := range positions {
    fuel += int(math.Abs(float64(target) - float64(position)))
  }

  return fuel
}

func findPath(positions []int, target int, max int) int {
  if target < 0 || target > max {
    return int(math.MaxUint >> 1)
  }

  currentFuel := calculateFuelSpent(positions, target)
  fmt.Printf("Fuel: %d, current target: %d\n", currentFuel, target)

  pathMin := calculateFuelSpent(positions, target - 1)
  pathMaj := calculateFuelSpent(positions, target + 1)
  
  if currentFuel >= pathMin {
    return findPath(positions, target - 1, max)
  }

  if currentFuel >= pathMaj {
    return findPath(positions, target + 1, max)
  }

  return currentFuel
}

func findMaxValue(positions []int) int {
  max := 0

  for _, position := range positions {
    if position > max {
      max = position
    }
  }

  return max
}

func main() {
  positions := readFile("input07.txt")

  mean := calculateMeanValue(positions)

  fuel := findPath(positions, mean, findMaxValue(positions))

  fmt.Printf("Fuel %d\n", fuel)
}
