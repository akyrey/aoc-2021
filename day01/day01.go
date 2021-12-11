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

func contains(s []int, e int) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}

func main() {
  f, err := os.Open("input01.txt")
  check(err)

  scanner := bufio.NewScanner(f)
  var count int
  prevValue := -1
  currentWindow := []int{-1, -1, -1}

  for scanner.Scan() {
    line := scanner.Text()
    currentValue, err := strconv.Atoi(line)
    check(err)

    // Remove first array element
    currentWindow = currentWindow[1:]
    // Insert the newly read one as last element
    currentWindow = append(currentWindow, currentValue)
    if !contains(currentWindow, -1) {
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
