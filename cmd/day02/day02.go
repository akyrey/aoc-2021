package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/akyrey/aoc-2021/internal"
)

func main() {
	f, err := internal.GetFileToReadFrom(2, internal.Test)
	internal.CheckError(err)
	defer f.Close()

	orizontalMovements := [1]string{"forward"}
	verticalMovements := [2]string{"up", "down"}

	scanner := bufio.NewScanner(f)

	horizontal := 0
	aim := 0
	depth := 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		movement := split[0]
		value, err := strconv.Atoi(split[1])
		internal.CheckError(err)

		if internal.Contains(orizontalMovements[:], movement) {
			horizontal += value
			depth += aim * value
		} else if internal.Contains(verticalMovements[:], movement) {
			if movement == verticalMovements[0] {
				aim -= value
			} else {
				aim += value
			}
		}

		fmt.Printf("Horizontal position: %d, current aim: %d, vertical position: %d\n", horizontal, aim, depth)
	}

	fmt.Printf("Totals multiplied: %d\n", horizontal*depth)
}
