package day02

import (
	"bufio"
	"fmt"
	"github.com/akyrey/aoc-2021/utils"
	"strconv"
	"strings"
)

func Day02(test bool) {
	f, err := utils.GetFileToReadFrom(2, test)
	utils.CheckError(err)
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
		utils.CheckError(err)

		if utils.Contains(orizontalMovements[:], movement) {
			horizontal += value
			depth += aim * value
		} else if utils.Contains(verticalMovements[:], movement) {
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
