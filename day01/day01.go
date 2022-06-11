package day01

import (
	"bufio"
	"fmt"
	"github.com/akyrey/aoc-2021/utils"
	"strconv"
)

func Day01(test bool) {
	f, err := utils.GetFileToReadFrom(1, test)
	utils.CheckError(err)

	scanner := bufio.NewScanner(f)
	var count int
	prevValue := -1
	currentWindow := []int{-1, -1, -1}

	for scanner.Scan() {
		line := scanner.Text()
		currentValue, err := strconv.Atoi(line)
		utils.CheckError(err)

		// Remove first array element
		currentWindow = currentWindow[1:]
		// Insert the newly read one as last element
		currentWindow = append(currentWindow, currentValue)
		if !utils.Contains(currentWindow, -1) {
			acc := 0
			for i := 0; i < len(currentWindow); i++ {
				acc += currentWindow[i]
			}

			fmt.Printf("Current window sum: %d", acc)

			if prevValue != -1 && acc > prevValue {
				fmt.Printf(" Increased\n")
				count++
			} else {
				fmt.Printf("\n")
			}

			prevValue = acc
		}
	}

	fmt.Printf("Total increases: %d\n", count)

	f.Close()
}
