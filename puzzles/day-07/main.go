package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readinput(inputfile string) ([]int, [][]int) {
	input, _ := os.ReadFile(inputfile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	results := []int{}
	operands := [][]int{}
	for _, line := range lines {
		split := strings.Split(line, ": ")
		res, _ := strconv.Atoi(split[0])
		results = append(results, res)
		ops := []int{}
		for _, v := range strings.Split(split[1], " ") {
			iv, _ := strconv.Atoi(v)
			ops = append(ops, iv)
		}
		operands = append(operands, ops)
	}
	return results, operands
}

func check(res int, operands []int) bool {
	if len(operands) == 0 {
		return false
	}

	i := len(operands)-1
	last := operands[i]

	if i == 0 && res == last {
		return true
	}

	if res != 0 && res % last == 0 {
		return check(res/last, operands[:i]) || check(res-last, operands[:i])
	}

	return check(res-last, operands[:i])
}

func part1() int {
	results, operands := readinput("input.txt")
	sum := 0
	for i := 0; i < len(results); i++ {
		if ok := check(results[i], operands[i]); ok {
			sum += results[i]
		}
	}
	return sum
}

func part2() int {
	return -1
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
