package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput(inputFile string) []string {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	return lines
}

func part1() int {
	sum := 0
	mulRegex := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	lines := readInput("input.txt")
	for _, line := range lines {
		muls := mulRegex.FindAllSubmatch([]byte(line), -1)
		for _, mul := range muls {
			a, _ := strconv.Atoi(string(mul[1]))
			b, _ := strconv.Atoi(string(mul[2]))
			sum += a * b
		}
	}
	return sum
}

func part2() int {
	sum := 0
	skip := false

	mulRegex := regexp.MustCompile("mul\\(\\d+,\\d+\\)|don't\\(\\)|do\\(\\)")
	operandsRegex := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

	lines := readInput("input.txt")

	for _, line := range lines {
		matches := mulRegex.FindAllSubmatch([]byte(line), -1)
		for _, match := range matches {
			token := string(match[0])
			if token == "don't()" {
				skip = true
			} else if token == "do()" {
				skip = false
			} else if !skip {
				operands := operandsRegex.FindSubmatch(match[0])
				a, _ := strconv.Atoi(string(operands[1]))
				b, _ := strconv.Atoi(string(operands[2]))
				sum += a * b
			}
		}
	}
	return sum
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
