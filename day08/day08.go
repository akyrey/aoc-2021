package day08

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/akyrey/aoc-2021/utils"
)

func readFile(test bool) [][]string {
	f, err := utils.GetFileToReadFrom(8, test)
	utils.CheckError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	result := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line read: %s\n", line)

		signalPatternAndOutputValue := strings.Split(line, " | ")
		if len(signalPatternAndOutputValue) != 2 {
			panic("Line isn't formatted properly")
		}
		result = append(result, signalPatternAndOutputValue)
	}

	return result
}

func readSignal(line string) []string {
	return strings.Split(line, " ")
}

func aggregateTotalOutputValues(readings [][]string) map[int][]string {
    // Output values are saved in a map with indexes from 0 to 9
    // Each index contains the list of strings we are sure represents the values (at the moment only 1, 4, 7 and 8 that 
    // have unique lengths)
    totalOutputValues := make(map[int][]string)

    for _, display := range readings {
		if len(display) != 2 {
			panic("Wrong display separation")
		}
        currentOutputValues := strings.Split(display[1], " ")
        if len(currentOutputValues) != 4 {
			panic("Output values should always have 4 values")
        }
        for _, value := range currentOutputValues {
            length := len(value)
            switch length {
            case 2:
                // This represents the number 1
                totalOutputValues[length] = append(totalOutputValues[length], value)
                break
            case 3:
                // This represents the number 7
                totalOutputValues[length] = append(totalOutputValues[length], value)
                break
            case 4:
                // This represents the number 4
                totalOutputValues[length] = append(totalOutputValues[length], value)
                break
            case 7:
                // This represents the number 8
                totalOutputValues[length] = append(totalOutputValues[length], value)
                break
            default:
                fmt.Printf("Value of length %d isn't unique: %s\n", len(value), value)
            }
        }
    }

    return totalOutputValues
}

func Day08(test bool) {
	readings := readFile(test)

    totalOutputValues := aggregateTotalOutputValues(readings)

    totalUniqueDigits := 0
    for _, signals := range totalOutputValues {
        totalUniqueDigits += len(signals)
    }
    fmt.Printf("Unique digits in output values %d\n", totalUniqueDigits)
}
