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

var cache = map[string]int{}

func combinationCount(design string, towels []string) int {
	if cached, hit := cache[design]; hit {
		return cached
	}
	if design == "" {
		return 1
	}
	possibleCombinations := 0
	for _, towel := range towels {
		if cut, ok := strings.CutPrefix(design, towel); ok {
			c := combinationCount(cut, towels)
			cache[cut] = c
			possibleCombinations += c
		}
	}
	return possibleCombinations
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

func part2() int {
	towels, designs := readInput("input.txt")
	possibleCombinations := 0
	for _, design := range designs {
		possibleCombinations += combinationCount(design, towels)
	}
	return possibleCombinations
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
