package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/akyrey/aoc-2021/internal"
)

func readFile(test bool) []int {
	f, err := internal.GetFileToReadFrom(6, test)
	internal.CheckError(err)
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
			internal.CheckError(err)

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
		acc += 1 + calculateChildren(8, daysLeft-i)
	}

	return acc
}

func main() {
	lanternFishes := readFile(internal.Test)
	totalDays := 256

	// Calculate children per starting value
	fishingMap := make([]int, 9)
	for i := 1; i <= 8; i++ {
		fishingMap[i] = calculateChildren(i, totalDays)
	}

	totalFishes := len(lanternFishes)
	for _, lanternFishCountdown := range lanternFishes {
		// Sum children per starting value previously calculated, avoiding duplicated computations
		totalFishes += fishingMap[lanternFishCountdown]
	}

	fmt.Printf("Total lantern fishes after %d days: %d\n", totalDays, totalFishes)
}
