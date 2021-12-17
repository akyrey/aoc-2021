package main

import (
	"bufio"
	"fmt"
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
  var lanternFishes []int

  for scanner.Scan() {
    line := scanner.Text()
    fmt.Printf("Line read: %s\n", line)

    lanternFishesAsString := strings.Split(line, ",")
    lanternFishes = make([]int, len(lanternFishesAsString))

    for i, stringValue := range lanternFishesAsString {
      value, err := strconv.Atoi(stringValue)
      check(err)

      lanternFishes[i] = value
    }

    fmt.Printf("Initial lantern fishes array: %v\n", lanternFishes)
  }

  return lanternFishes
}

func calculateChildren(daysUntilChildren int, daysLeft int) int {
  daysUntilChildren += 1
  if daysLeft < daysUntilChildren {
    return 0
  }

  acc := 0
  for i := daysUntilChildren; i <= daysLeft; i += 7 {
    acc += 1 + calculateChildren(8, daysLeft - i)
  }

  return acc
}

func main() {
  lanternFishes := readFile("test06.txt")
  totalDays := 80

  totalFishes := len(lanternFishes)
  for _, lanternFishCountdown := range lanternFishes {
    totalFishes += calculateChildren(lanternFishCountdown, totalDays)
  }

  fmt.Printf("Total lantern fishes after %d days: %d\n", totalDays, totalFishes)
}
