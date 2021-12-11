package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  f, err := os.Open("input.txt")
  check(err)

  scanner := bufio.NewScanner(f)
  var count int = -1
  var prevValue int
  for scanner.Scan() {
    var line string = scanner.Text()
    currentValue, err := strconv.Atoi(line)
    if err == nil && prevValue != -1 && prevValue < currentValue {
      count++
    }

    prevValue = currentValue
  }

  fmt.Printf("Total increases: %d\n", count)

  f.Close()
}
