package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"

	"github.com/akyrey/aoc-2021/internal"
)

type Corrupted struct {
	char string
	line string
}

func (corrupted *Corrupted) String() string {
	return fmt.Sprintf("Corrupted{char '%s', line '%s'}", corrupted.char, corrupted.line)
}

const (
	OPEN_PARENTHESES  string = "("
	OPEN_SQUARE       string = "["
	OPEN_CURLY        string = "{"
	OPEN_ANGLE        string = "<"
	CLOSE_PARENTHESES string = ")"
	CLOSE_SQUARE      string = "]"
	CLOSE_CURLY       string = "}"
	CLOSE_ANGLE       string = ">"
)

var (
	openings = []string{OPEN_PARENTHESES, OPEN_SQUARE, OPEN_CURLY, OPEN_ANGLE}
	// closings = []string{CLOSE_PARENTHESES, CLOSE_SQUARE, CLOSE_CURLY, CLOSE_ANGLE}
)

func isMatchingTag(stack []string, value string) bool {
	if len(stack) == 0 {
		return false
	}

	opening := stack[len(stack)-1]
	switch value {
	case CLOSE_PARENTHESES:
		return opening == OPEN_PARENTHESES
	case CLOSE_SQUARE:
		return opening == OPEN_SQUARE
	case CLOSE_CURLY:
		return opening == OPEN_CURLY
	case CLOSE_ANGLE:
		return opening == OPEN_ANGLE
	}

	return false
}

func completeLine(stack []string) []string {
	result := make([]string, 0)

	for i := len(stack) - 1; i >= 0; i-- {
		switch stack[i] {
		case OPEN_PARENTHESES:
			result = append(result, CLOSE_PARENTHESES)
		case OPEN_SQUARE:
			result = append(result, CLOSE_SQUARE)
		case OPEN_CURLY:
			result = append(result, CLOSE_CURLY)
		case OPEN_ANGLE:
			result = append(result, CLOSE_ANGLE)
		}
	}

	return result
}

func readFile(test bool) ([]Corrupted, []int) {
	f, err := internal.GetFileToReadFrom(10, test)
	internal.CheckError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	corrupted := make([]Corrupted, 0)
	completionScores := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "")
		stack := make([]string, 0)
		incomplete := true

		for _, value := range values {
			if internal.Contains(openings, value) {
				stack = append(stack, value)
			} else if isMatchingTag(stack, value) {
				stack = stack[:len(stack)-1]
			} else {
				corrupted = append(corrupted, Corrupted{char: value, line: line})
				incomplete = false
				break
			}
		}

		if incomplete {
			completedLine := completeLine(stack)
			fmt.Printf("Completed line %v\n", completedLine)
			score := 0
			for _, value := range completedLine {
				score = (score * 5) + completeLineScore(value)
			}
			completionScores = append(completionScores, score)
		}
	}

	return corrupted, completionScores
}

func syntaxErrorValue(value string) int {
	switch value {
	case CLOSE_PARENTHESES:
		return 3
	case CLOSE_SQUARE:
		return 57
	case CLOSE_CURLY:
		return 1197
	case CLOSE_ANGLE:
		return 25137
	}

	return 0
}

func completeLineScore(value string) int {
	switch value {
	case CLOSE_PARENTHESES:
		return 1
	case CLOSE_SQUARE:
		return 2
	case CLOSE_CURLY:
		return 3
	case CLOSE_ANGLE:
		return 4
	}

	return 0
}

func calcSyntaxError(corrupted []Corrupted) int {
	sum := 0

	for _, value := range corrupted {
		sum += syntaxErrorValue(value.char)
	}

	return sum
}

func main() {
	corrupted, completionScores := readFile(internal.Test)

	fmt.Printf("Corrupted lines: %#v\n", corrupted)

	total := calcSyntaxError(corrupted)

	fmt.Printf("Total syntax error %d\n", total)

	sort.Ints(completionScores)
	middleScore := completionScores[len(completionScores)/2]

	fmt.Printf("Middle score %d\n", middleScore)
}
