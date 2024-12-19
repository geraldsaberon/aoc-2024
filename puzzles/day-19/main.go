package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(inputFile string) (towels []string, designs []string) {
	input, _ := os.ReadFile(inputFile)
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	towels = strings.Split(split[0], ", ")
	designs = strings.Split(split[1], "\n")
	return
}

func isPossible(design string, towels []string) bool {
	if design == "" {
		return true
	}
	for _, towel := range towels {
		if cut, ok := strings.CutPrefix(design, towel); ok {
			if p := isPossible(cut, towels); p {
				return p
			}
		}
	}
	return false
}

func part1() int {
	towels, designs := readInput("input.txt")
	possible := 0
	for _, design := range designs {
		if isPossible(design, towels) {
			possible += 1
		}
	}
	return possible
}

func main() {
	fmt.Println("Part 1:", part1())
}
