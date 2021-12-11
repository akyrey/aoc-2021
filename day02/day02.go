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
  f, err := os.Open("input.txt")
  check(err)
  orizontalMovements := [1]string{"forward"}
  verticalMovements := [2]string{"up","down"}

  scanner := bufio.NewScanner(f)

  horizontal := 0
  depth := 0

  for scanner.Scan() {
    line := scanner.Text()
    split := strings.Split(line, " ")
    movement := split[0]
    value, err := strconv.Atoi(split[1])
    check(err)

    if contains(orizontalMovements[:], movement) {
      horizontal += value
    } else if contains(verticalMovements[:], movement) {
      if movement == verticalMovements[0] {
        depth -= value
      } else {
        depth += value
      }
    }

    fmt.Printf("Horizontal position: %d, vertical position: %d\n", horizontal, depth)
  }

  fmt.Printf("Totals multiplied: %d\n", horizontal * depth)
}
