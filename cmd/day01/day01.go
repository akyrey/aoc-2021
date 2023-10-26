package main

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/akyrey/aoc-2021/internal"
)

func main() {
	f, err := internal.GetFileToReadFrom(1, internal.Test)
	internal.CheckError(err)

	scanner := bufio.NewScanner(f)
	var count int
	prevValue := -1
	currentWindow := []int{-1, -1, -1}

	for scanner.Scan() {
		line := scanner.Text()
		currentValue, err := strconv.Atoi(line)
		internal.CheckError(err)

		// Remove first array element
		currentWindow = currentWindow[1:]
		// Insert the newly read one as last element
		currentWindow = append(currentWindow, currentValue)
		if !internal.Contains(currentWindow, -1) {
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
