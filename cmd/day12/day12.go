package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/akyrey/aoc-2021/internal"
)

func retrieveConnections(filename string) map[string][]string {
	f, err := os.Open(filename)
	internal.CheckError(err)
	defer f.Close()

	connections := make(map[string][]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "-")
		if len(values) != 2 {
			panic("Wrong input")
		}

		if _, ok := connections[values[0]]; !ok {
			connections[values[0]] = []string{values[1]}
		} else {
			connections[values[0]] = append(connections[values[0]], values[1])
		}
	}

	return connections
}

func main() {}
