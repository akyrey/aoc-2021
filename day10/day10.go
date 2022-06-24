package day10

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/akyrey/aoc-2021/utils"
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

var openings = []string{OPEN_PARENTHESES, OPEN_SQUARE, OPEN_CURLY, OPEN_ANGLE}
var closings = []string{CLOSE_PARENTHESES, CLOSE_SQUARE, CLOSE_CURLY, CLOSE_ANGLE}

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

func readFile(test bool) []Corrupted {
	f, err := utils.GetFileToReadFrom(10, test)
	utils.CheckError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	corrupted := make([]Corrupted, 0)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "")
		stack := make([]string, 0)

		for _, value := range values {
			if utils.Contains(openings, value) {
				stack = append(stack, value)
			} else if isMatchingTag(stack, value) {
				stack = stack[:len(stack)-1]
			} else {
				corrupted = append(corrupted, Corrupted{char: value, line: line})
				break
			}
		}
	}

	return corrupted
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

func calcSyntaxError(corrupted []Corrupted) int {
	sum := 0

	for _, value := range corrupted {
		sum += syntaxErrorValue(value.char)
	}

	return sum
}

func Day10(test bool) {
	corrupted := readFile(test)

	fmt.Printf("Corrupted lines: %#v\n", corrupted)

	total := calcSyntaxError(corrupted)

	fmt.Printf("Total syntax error %d\n", total)
}
