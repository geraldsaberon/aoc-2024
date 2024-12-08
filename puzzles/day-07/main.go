package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
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

func check1(res int, operands []int) bool {
	if len(operands) == 0 {
		return false
	}

	i := len(operands)-1
	last := operands[i]

	if i == 0 && res == last {
		return true
	}

	if res % last == 0 {
		return check1(res/last, operands[:i]) || check1(res-last, operands[:i])
	}

	return check1(res-last, operands[:i])
}

func intlen(n int) int {
	if n == 0 {
		return 1
	}
	len_ := 0
	for n > 0 {
		n /= 10
		len_ += 1
	}
	return len_
}

func checkconcat(res, n int) (bool, int) {
	l := intlen(n)
	div := int(math.Pow(float64(10), float64(l)))
	return res % div == n, div
}

func check2(res int, operands []int) bool {
	if len(operands) == 0 {
		return false
	}

	i := len(operands)-1
	last := operands[i]

	if i == 0 && res == last {
		return true
	}

	if res % last == 0 {
		if ok, div := checkconcat(res, last); ok {
			return check2(res/last, operands[:i]) || check2(res-last, operands[:i]) || check2(res/div, operands[:i])
		}
		return check2(res/last, operands[:i]) || check2(res-last, operands[:i])
	}

	if ok, div := checkconcat(res, last); ok {
		return check2(res/div, operands[:i]) || check2(res-last, operands[:i])
	}

	return check2(res-last, operands[:i])
}

func part1() int {
	results, operands := readinput("input.txt")
	sum := 0
	for i := 0; i < len(results); i++ {
		if ok := check1(results[i], operands[i]); ok {
			sum += results[i]
		}
	}
	return sum
}

func part2() int {
	results, operands := readinput("input.txt")
	sum := 0
	for i := 0; i < len(results); i++ {
		if ok := check2(results[i], operands[i]); ok {
			sum += results[i]
		}
	}
	return sum
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
