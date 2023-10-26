package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/akyrey/aoc-2021/internal"
)

type Digit struct {
	Encoded      *string
	Decoded      *string
	Count        int
	EncodedBytes byte
	DecodedBytes byte
}

func (digit *Digit) String() string {
	return fmt.Sprintf("Digit{Encoded: '%s', EncodedBytes: '%b', Decoded: '%s', DecodedBytes: '%b', Count: %d}", internal.StringPtrToString(digit.Encoded), digit.EncodedBytes, internal.StringPtrToString(digit.Decoded), digit.DecodedBytes, digit.Count)
}

func displayDigitRepresentation() map[rune]byte {
	return map[rune]byte{
		'a': 0b1000000,
		'b': 0b0100000,
		'c': 0b0010000,
		'd': 0b0001000,
		'e': 0b0000100,
		'f': 0b0000010,
		'g': 0b0000001,
	}
}

func representationToDisplayDigit() map[byte]rune {
	return map[byte]rune{
		0b1000000: 'a',
		0b0100000: 'b',
		0b0010000: 'c',
		0b0001000: 'd',
		0b0000100: 'e',
		0b0000010: 'f',
		0b0000001: 'g',
	}
}

/**
*   0:      1:      2:      3:      4:
*  aaaa    ....    aaaa    aaaa    ....
* b    c  .    c  .    c  .    c  b    c
* b    c  .    c  .    c  .    c  b    c
*  ....    ....    dddd    dddd    dddd
* e    f  .    f  e    .  .    f  .    f
* e    f  .    f  e    .  .    f  .    f
*  gggg    ....    gggg    gggg    ....
*
*   5:      6:      7:      8:      9:
*  aaaa    aaaa    aaaa    aaaa    aaaa
* b    .  b    .  .    c  b    c  b    c
* b    .  b    .  .    c  b    c  b    c
*  dddd    dddd    ....    dddd    dddd
* .    f  e    f  .    f  e    f  .    f
* .    f  e    f  .    f  e    f  .    f
*  gggg    gggg    ....    gggg    gggg
*
*
* [a, b, c, d, e, f, g]
*
 */
func getDecodedBytes(index int) byte {
	switch index {
	case 0:
		// abcefg
		return 0b1110111
	case 1:
		// cf
		return 0b0010010
	case 2:
		// acdeg
		return 0b1011101
	case 3:
		// acdfg
		return 0b1011011
	case 4:
		// bcdf
		return 0b0111010
	case 5:
		// abdfg
		return 0b1101011
	case 6:
		// abdefg
		return 0b1101111
	case 7:
		// acf
		return 0b1010010
	case 8:
		// abcdefg
		return 0b1111111
	case 9:
		// abcdfg
		return 0b1111011
	}

	return 0b0000000
}

func getIntFromBytes(value byte) int {
	switch value {
	case 0b1110111:
		// abcefg
		return 0
	case 0b0010010:
		// cf
		return 1
	case 0b1011101:
		// acdeg
		return 2
	case 0b1011011:
		// acdfg
		return 3
	case 0b0111010:
		// bcdf
		return 4
	case 0b1101011:
		// abdfg
		return 5
	case 0b1101111:
		// abdefg
		return 6
	case 0b1010010:
		// acf
		return 7
	case 0b1111111:
		// abcdefg
		return 8
	case 0b1111011:
		// abcdfg
		return 9
	}

	return 0
}

func getEncodedBytes(value string) byte {
	encode := byte(0b0000000)
	representation := displayDigitRepresentation()

	for _, key := range value {
		encode |= representation[key]
	}

	return encode
}

func readFile(test bool) [][]string {
	f, err := internal.GetFileToReadFrom(8, test)
	internal.CheckError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	result := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("Line read: %s\n", line)

		signalPatternAndOutputValue := strings.Split(line, " | ")
		if len(signalPatternAndOutputValue) != 2 {
			panic("Line isn't formatted properly")
		}
		result = append(result, signalPatternAndOutputValue)
	}

	return result
}

func updateOutputValues(totalOutputValues map[int]Digit, index int, value string) {
	// fmt.Printf("Value '%s' represents %d\n", value, index)
	if entry, present := totalOutputValues[index]; present {
		entry.Count++
		totalOutputValues[index] = entry
		return
	}

	sortedVal := internal.SortString(value)
	totalOutputValues[index] = Digit{
		Decoded:      nil,
		DecodedBytes: getDecodedBytes(index),
		Encoded:      &sortedVal,
		EncodedBytes: getEncodedBytes(value),
		Count:        1,
	}
}

func decodeAllDigits(digits map[int]Digit, fiveLetters, sixLetters []string, decodedDigits map[rune]byte) map[rune]byte {
	actualValue := representationToDisplayDigit()
	fiveLettersEncoded := make([]byte, 0)
	sixLettersEncoded := make([]byte, 0)
	for _, value := range fiveLetters {
		fiveLettersEncoded = append(fiveLettersEncoded, getEncodedBytes(value))
		fmt.Printf("%s: %b\n", value, getEncodedBytes(value))
	}
	for _, value := range sixLetters {
		sixLettersEncoded = append(sixLettersEncoded, getEncodedBytes(value))
		fmt.Printf("%s: %b\n", value, getEncodedBytes(value))
	}
	// 'a' can be decoded using 7 and 1
	one, oneOk := digits[1]
	seven, sevenOk := digits[7]
	if oneOk && sevenOk {
		decodedDigits['a'] = seven.EncodedBytes - one.EncodedBytes
		fmt.Printf("Given that 1 is %b, 7 is %b, then 'a' must be %b (%c)\n", one.EncodedBytes, seven.EncodedBytes, decodedDigits['a'], actualValue[decodedDigits['a']])
	}
	// 'g' is common to all five and six letter values, minus 'a' that is common too
	decodedDigits['g'] = 0b1111111
	for _, value := range append(fiveLettersEncoded, sixLettersEncoded...) {
		decodedDigits['g'] = decodedDigits['g'] & value
	}
	decodedDigits['g'] -= decodedDigits['a']
	fmt.Printf("Since 'g' is the only common letter among all five and six letters with 'a', it must be %b (%c)\n", decodedDigits['g'], actualValue[decodedDigits['g']])
	// 'd' can be retrieved from six letters knowing 'a' and 'g'
	decodedDigits['d'] = 0b1111111
	for _, value := range fiveLettersEncoded {
		decodedDigits['d'] = decodedDigits['d'] & value
	}
	decodedDigits['d'] = decodedDigits['d'] - decodedDigits['a'] - decodedDigits['g']
	fmt.Printf("Given 'a' and 'g', the only common letter on five letters left is 'd', so it must be %b (%c)\n", decodedDigits['d'], actualValue[decodedDigits['d']])
	// 'b' can be retrieved from 9, knowing it has 6 letters, contains representation of 1 and has 'a', 'd' and 'g'
	// 'e' can be retrieved from 0, knowing it has 6 letters, contains representation of 1 and has 'a', 'b' and 'g'
	// 'f' can be retrieved from 6, given we now know 'a', 'b', 'd', 'e' and 'g'
	nineIndex := -1
	zeroIndex := -1
	sixIndex := -1
	test := representationToDisplayDigit()
	d := string(test[decodedDigits['d']])
	for index, value := range sixLetters {
		if strings.Contains(value, (*one.Encoded)[0:1]) && strings.Contains(value, (*one.Encoded)[1:2]) && strings.Contains(value, d) {
			nineIndex = index
		} else if strings.Contains(value, (*one.Encoded)[0:1]) && strings.Contains(value, (*one.Encoded)[1:2]) && !strings.Contains(value, d) {
			zeroIndex = index
		} else {
			sixIndex = index
		}
	}
	if oneOk && nineIndex != -1 {
		decodedDigits['b'] = sixLettersEncoded[nineIndex] - one.EncodedBytes - decodedDigits['a'] - decodedDigits['d'] - decodedDigits['g']
	}
	fmt.Printf("Given 'a', 'd', 'g' and knowing encoded representation of 1, we can retrieve 'b' from 9 as %b (%c)\n", decodedDigits['b'], actualValue[decodedDigits['b']])
	if oneOk && zeroIndex != -1 {
		decodedDigits['e'] = sixLettersEncoded[zeroIndex] - one.EncodedBytes - decodedDigits['a'] - decodedDigits['b'] - decodedDigits['g']
	}
	fmt.Printf("Given 'a', 'b', 'g' and knowing encoded representation of 1, we can retrieve 'e' from 0 as %b (%c)\n", decodedDigits['e'], actualValue[decodedDigits['e']])
	if sixIndex != -1 {
		decodedDigits['f'] = sixLettersEncoded[sixIndex] - decodedDigits['a'] - decodedDigits['b'] - decodedDigits['d'] - decodedDigits['e'] - decodedDigits['g']
	}
	fmt.Printf("Given 'a', 'b', 'd', 'e', 'g' and knowing encoded representation of 6, we can retrieve 'f' %b (%c)\n", decodedDigits['f'], actualValue[decodedDigits['f']])
	// 'c' can be retrieved as last available value
	decodedDigits['c'] = 0b1111111 - decodedDigits['a'] - decodedDigits['b'] - decodedDigits['d'] - decodedDigits['e'] - decodedDigits['f'] - decodedDigits['g']
	fmt.Printf("Given all other values, 'c' can be retrieved as %b (%c)\n", decodedDigits['c'], actualValue[decodedDigits['c']])

	return decodedDigits
}

func decodeSingleLine(pattern, display string) int {
	// Output values are saved in a map with indexes from 0 to 9
	// Each index contains the list of strings we are sure represents the values (at the moment only 1, 4, 7 and 8 that
	// have unique lengths)
	outputDigits := make(map[int]Digit)
	fiveLetters := make([]string, 0)
	sixLetters := make([]string, 0)
	encodedDigits := map[rune]byte{}

	patternValues := strings.Split(pattern, " ")
	displayValues := strings.Split(display, " ")
	if len(patternValues) != 10 {
		panic("Pattern values should always be 10")
	}
	if len(displayValues) != 4 {
		panic("Display values should always be 4")
	}
	for _, value := range patternValues {
		length := len(value)
		// fmt.Printf("Read '%s' with length %d\n", value, length)
		switch length {
		case 2:
			// This represents the number 1
			updateOutputValues(outputDigits, 1, value)
		case 3:
			// This represents the number 7
			updateOutputValues(outputDigits, 7, value)
		case 4:
			// This represents the number 4
			updateOutputValues(outputDigits, 4, value)
		case 5:
			// fmt.Printf("Value with length %d %s could be a 2, 3 or 5\n", length, value)
			sortedValue := internal.SortString(value)
			if !internal.Contains(fiveLetters, sortedValue) {
				fiveLetters = append(fiveLetters, sortedValue)
			}
		case 6:
			// fmt.Printf("Value with length %d %s could be a 0, 6 or 9\n", length, value)
			sortedValue := internal.SortString(value)
			if !internal.Contains(sixLetters, sortedValue) {
				sixLetters = append(sixLetters, sortedValue)
			}
		case 7:
			// This represents the number 8
			updateOutputValues(outputDigits, 8, value)
		}
	}

	fmt.Printf("Found 5 letters: %v\nand 6 letters: %v\n", fiveLetters, sixLetters)
	encodedDigits = decodeAllDigits(outputDigits, fiveLetters, sixLetters, encodedDigits)
	displayedValueToBytes := displayDigitRepresentation()
	translatedValues := translateDecodedValues(encodedDigits)

	// decodedDigits := representationToDisplayDigit()
	// for key, value := range decodedDigits {
	//     fmt.Printf("%c => %b %d\n", key, value, getIntFromBytes(value))
	// }
	totalValue := 0
	for index, value := range displayValues {
		currentValue := ""
		for _, char := range value {
			currentValue += string(translatedValues[displayedValueToBytes[char]])
		}
		orderedValue := internal.SortString(currentValue)
		encodedValue := getEncodedBytes(orderedValue)
		intValue := getIntFromBytes(encodedValue)
		fmt.Printf("Value: %s %s => %d\n", value, orderedValue, intValue)
		if index == 0 {
			totalValue += 1000 * intValue
		} else if index == 1 {
			totalValue += 100 * intValue
		} else if index == 2 {
			totalValue += 10 * intValue
		} else if index == 3 {
			totalValue += intValue
		}
	}
	fmt.Printf("Total line value %d\n", totalValue)

	return totalValue
}

func translateDecodedValues(encodedDigits map[rune]byte) map[byte]rune {
	result := make(map[byte]rune, 0)

	for key, value := range encodedDigits {
		result[value] = key
	}

	return result
}

func aggregateDigitsPerLine(readings [][]string) []int {
	result := make([]int, 0)

	for _, display := range readings {
		if len(display) != 2 {
			panic("Wrong display separation")
		}

		result = append(result, decodeSingleLine(display[0], display[1]))
	}

	return result
}

func main() {
	readings := readFile(internal.Test)

	outputLines := aggregateDigitsPerLine(readings)
	totalAdd := 0

	for _, line := range outputLines {
		totalAdd += line
	}

	fmt.Printf("Adding all displayed lines gives: %d\n", totalAdd)
}
