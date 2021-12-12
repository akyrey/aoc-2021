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

func arrayToString(a []int, delim string) string {
  return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func calculateRate(totalCount int, reportBitSum []int, mostCommon bool) int {
  rate := make([]int, len(reportBitSum), cap(reportBitSum))

  for i := 0; i < len(reportBitSum); i++ {
    var value int
    if mostCommon {
      value = getMostCommonBit(totalCount, reportBitSum[i])
    } else {
      value = getLeastCommonBit(totalCount, reportBitSum[i])
    }

    rate[i] = value
  }

  binaryValue := arrayToString(rate[:], "")
  decimalValue, err := strconv.ParseInt(binaryValue, 2, 64)
  check(err)

  return int(decimalValue)
}

func getMostCommonBit(totalCount int, value int) int {
  zeroValueLimit := totalCount / 2;

  if value < zeroValueLimit {
    return 0
  }

  return 1
}

func getLeastCommonBit(totalCount int, value int) int {
  zeroValueLimit := totalCount / 2;

  if value < zeroValueLimit {
    return 1
  }

  return 0
}

func main() {
  f, err := os.Open("input03.txt")
  check(err)

  scanner := bufio.NewScanner(f)
  totalCount := 0
  var reportBitSum []int

  for scanner.Scan() {
    line := scanner.Text()
    split := strings.Split(line, "")
    if reportBitSum == nil {
      reportBitSum = make([]int, len(split), len(split))
    }

    fmt.Printf("Split %v\n", split)
    for i := 0; i < len(split); i++ {
      value, err := strconv.Atoi(split[i])
      check(err)

      reportBitSum[i] += value
    }

    totalCount++
  }

  gammaRate := calculateRate(totalCount, reportBitSum, true)
  epsilonRate := calculateRate(totalCount, reportBitSum, false)

  fmt.Printf("Total lines: %d, reportBitSum: %v, gammaRate: %d, epsilonRate: %d, power consumption: %d\n", totalCount, reportBitSum, gammaRate, epsilonRate, gammaRate * epsilonRate)
}
