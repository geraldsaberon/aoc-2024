package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

func readInput(inputFile string) []string {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	return lines
}

func isSafe(levels []string) bool {
	prevLevel := 0
	isIncreasing := 0
	isDecreasing := 0

	for _, v := range levels {
		currentLevel, _ := strconv.Atoi(v)

		if prevLevel == 0 {
			prevLevel = currentLevel
			continue
		}

		diff := currentLevel - prevLevel

		if diff == 0 || math.Abs(float64(diff)) > 3 {
			return false
		}

		if diff > 0 { isIncreasing += 1 }
		if diff < 0 { isDecreasing += 1 }

		if isIncreasing > 0 && isDecreasing > 0 {
			return false
		}

		prevLevel = currentLevel
	}

	return true
}

func remove(slice []string, i int) []string {
    return append(slice[:i], slice[i+1:]...)
}

func part1() int {
	lines := readInput("input.txt")
	isSafeSum := 0
	for _, line := range lines {
		levels := strings.Split(line, " ")
		if isSafe(levels) {
			isSafeSum += 1
		}
	}
	return isSafeSum
}

func part2() int {
	lines := readInput("input.txt")
	isSafeSum := 0
	for _, line := range lines {
		levels := strings.Split(line, " ")
		safe := isSafe(levels)
		for i, max := 0, len(levels); !safe && i < max; {
			levels = remove(strings.Split(line, " "), i)
			safe = isSafe(levels)
			if !safe {
				i += 1
			}
		}
		if safe {
			isSafeSum += 1
		}
	}
	return isSafeSum
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
