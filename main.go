package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/akyrey/aoc-2021/day01"
	"github.com/akyrey/aoc-2021/day02"
	"github.com/akyrey/aoc-2021/day03"
	"github.com/akyrey/aoc-2021/day04"
	"github.com/akyrey/aoc-2021/day05"
	"github.com/akyrey/aoc-2021/day06"
	"github.com/akyrey/aoc-2021/day07"
	"github.com/akyrey/aoc-2021/day08"
	"github.com/akyrey/aoc-2021/day09"
	"github.com/akyrey/aoc-2021/day10"
	"github.com/akyrey/aoc-2021/day11"
)

func main() {
	const test = true
	fmt.Printf("Insert day:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		day, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Please insert a number (0 to exit)")
			continue
		}

		switch day {
		case 0:
			fmt.Println("Thanks for playing with us!")
			os.Exit(0)
		case 1:
			fmt.Println("Chosen day 1")
			day01.Day01(test)
		case 2:
			fmt.Println("Chosen day 2")
			day02.Day02(test)
		case 3:
			fmt.Println("Chosen day 3")
			day03.Day03(test)
		case 4:
			fmt.Println("Chosen day 4")
			day04.Day04(test)
		case 5:
			fmt.Println("Chosen day 5")
			day05.Day05(test)
		case 6:
			fmt.Println("Chosen day 6")
			day06.Day06(test)
		case 7:
			fmt.Println("Chosen day 7")
			day07.Day07(test)
		case 8:
			fmt.Println("Chosen day 8")
			day08.Day08(test)
		case 9:
			fmt.Println("Chosen day 9")
			day09.Day09(test)
		case 10:
			fmt.Println("Chosen day 10")
			day10.Day10(test)
		case 11:
			fmt.Println("Chosen day 11")
			day11.Day11(test)
		default:
			fmt.Println("Unknown day")
		}
	}
}
