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

func simulateDay(initialLanternFishes []int) (lanternFishes []int) {
  var fishesToAdd []int
  for _, countdown := range initialLanternFishes {
    // Reset countdown to 6 and create a new lantern fish with a countdown of 8
    if countdown == 0 {
      lanternFishes = append(lanternFishes, 6)
      fishesToAdd = append(fishesToAdd, 8)
    } else {
      // Simply decrement countdown
      lanternFishes = append(lanternFishes, countdown - 1)
    }
  }

  return append(lanternFishes, fishesToAdd...)
}

func main() {
  lanternFishes := readFile("input06.txt")
  totalDays := 80

  for day := 1; day <= totalDays; day++ {
    lanternFishes = simulateDay(lanternFishes)

    fmt.Printf("Lantern fishes after %d days: %v\n", day, lanternFishes)
  }

  fmt.Printf("Total lantern fishes after %d days: %d\n", totalDays, len(lanternFishes))
}
