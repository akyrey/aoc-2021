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

func contains(s []string, e string) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}

func main() {
  f, err := os.Open("input02.txt")
  check(err)
  orizontalMovements := [1]string{"forward"}
  verticalMovements := [2]string{"up","down"}

  scanner := bufio.NewScanner(f)

  horizontal := 0
  aim := 0
  depth := 0

  for scanner.Scan() {
    line := scanner.Text()
    split := strings.Split(line, " ")
    movement := split[0]
    value, err := strconv.Atoi(split[1])
    check(err)

    if contains(orizontalMovements[:], movement) {
      horizontal += value
      depth += aim * value
    } else if contains(verticalMovements[:], movement) {
      if movement == verticalMovements[0] {
        aim -= value
      } else {
        aim += value
      }
    }

    fmt.Printf("Horizontal position: %d, current aim: %d, vertical position: %d\n", horizontal, aim, depth)
  }

  fmt.Printf("Totals multiplied: %d\n", horizontal * depth)
}
