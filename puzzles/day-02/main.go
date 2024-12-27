package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

func readInput(inputFile string) [][]int {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	reports := [][]int{}
	for _, line := range lines {
		report := []int{}
		for _, level := range strings.Split(line, " ") {
			intLevel, _ := strconv.Atoi(level)
			report = append(report, intLevel)
		}
		reports = append(reports, report)
	}
	return reports
}

func isSafe(levels []int) bool {
	increments, decrements := 0, 0
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if math.Abs(float64(diff)) > 3 || diff == 0 {
			return false
		}
		if diff > 0 {
			increments += 1
		} else {
			decrements += 1
		}
		if increments > 0 && decrements > 0 {
			return false
		}
	}
	return true
}

func remove(s []int, i int) []int {
	c := make([]int, len(s))
	copy(c, s)
    return append(c[:i], c[i+1:]...)
}

func part1() int {
	reports := readInput("input.txt")
	isSafeSum := 0
	for _, report := range reports {
		if isSafe(report) {
			isSafeSum += 1
		}
	}
	return isSafeSum
}

func part2() int {
	reports := readInput("input.txt")
	isSafeSum := 0
	for _, report := range reports {
		safe := isSafe(report)
		for i := 0; !safe && i < len(report); i++ {
			safe = isSafe(remove(report, i))
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
