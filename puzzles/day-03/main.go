package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readInput(inputFile string) []byte {
	input, _ := os.ReadFile(inputFile)
	return input
}

func part1() int {
	sum := 0
	mulRegex := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	input := readInput("input.txt")
	muls := mulRegex.FindAllSubmatch(input, -1)
	for _, mul := range muls {
		a, _ := strconv.Atoi(string(mul[1]))
		b, _ := strconv.Atoi(string(mul[2]))
		sum += a * b
	}
	return sum
}

func part2() int {
	sum := 0
	skip := false

	mulRegex := regexp.MustCompile("mul\\(\\d+,\\d+\\)|don't\\(\\)|do\\(\\)")
	operandsRegex := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

	input := readInput("input.txt")
	matches := mulRegex.FindAllSubmatch(input, -1)
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
	return sum
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
