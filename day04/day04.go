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

func readExtractions(line string) []int {
    split := strings.Split(line, ",")
    extractions := make([]int, len(split), len(split))

    for i, v := range split {
      value, err := strconv.Atoi(v)
      check(err)
      extractions[i] = value
    }

    return extractions
}

func readBoardLine(line string) []int {
  var lineNumbers []int
  split := strings.Split(line, " ")

  for _, v := range split {
    if (v != "") {
      value, err := strconv.Atoi(v)
      check(err)
      lineNumbers = append(lineNumbers, value)
    }
  }

  return lineNumbers
}

func readFile(filename string) (extractions []int, boards [][][]int) {
  f, err := os.Open(filename)
  check(err)
  defer f.Close()

  scanner := bufio.NewScanner(f)
  currentBoard := -1

  for lineNumber := 0; scanner.Scan(); lineNumber++ {
    line := scanner.Text()

    if lineNumber == 0 {
      extractions = readExtractions(line)
    } else if line == "" {
      boards = append(boards, make([][]int, 0))
      currentBoard++
      fmt.Printf("Created a new board\n")
    } else {
      boards[currentBoard] = append(boards[currentBoard], readBoardLine(line))
      fmt.Printf("Add row to board %d\n", currentBoard)
    }
  }

  return
}

func contains(s []int, e int) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}

func checkWinningRow(line []int, extracted []int) bool {
  fmt.Printf("Checking line %v with extraction %v", line, extracted)
  rowMatched := 0
  for _, v := range extracted {
    if contains(line, v) {
      rowMatched++
      fmt.Printf(" DOES CONTAIN VALUE, matched %d\n", rowMatched)
    }
  }

  return rowMatched == 5
}

func invertColsAndRows(board [][]int) (inverted [][]int) {
  inverted = make([][]int, 5, 5)
  for col := 0; col < 5; col++ {
    for row := 0; row < 5; row++ {
      inverted[row] = append(inverted[row], board[col][row])
    }
  }

  return
}

func checkWinningBoard(boards [][][]int, extracted []int) (winners []int) {
  for index, board := range boards {
    invertedBoard := invertColsAndRows(board)
    fmt.Printf("board %v, inverted: %v\n", board, invertedBoard)
    for col := 0; col < 5; col++ {
      winningRow := checkWinningRow(board[col], extracted) 
      winningColumn := checkWinningRow(invertedBoard[col], extracted)  
      if winningRow || winningColumn {
        winners = append(winners, index)
      }
    }
  }

  return
}

func calculateBoardScore(board [][]int, extracted []int) int {
  unmarkedNumbersTotal := 0
  for row := 0; row < 5; row++ {
    for col := 0; col < 5; col++ {
      if !contains(extracted, board[row][col]) {
        unmarkedNumbersTotal += board[row][col]
      }
    }
  }

  multiplier := extracted[len(extracted) - 1]

  return unmarkedNumbersTotal * multiplier
}

func main() {
  extractions, boards := readFile("input04.txt")
  var extracted []int
  var winners []int

  fmt.Printf("Extractions: %v\nBoards: %v\n", extractions, boards)

  for _, extraction := range extractions {
    extracted = append(extracted, extraction)
    fmt.Printf("Extracted %v\n", extracted)
    winners = checkWinningBoard(boards, extracted)
    if len(winners) > 0 {
      fmt.Printf("Finished, winning boards %v, extracted %v\n", winners, extracted)
      break
    }
  }

  boardScore := calculateBoardScore(boards[winners[0]], extracted)
  fmt.Printf("Board %d total score %d\n", winners[0], boardScore)
}
