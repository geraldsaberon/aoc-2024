package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getLists(inputFile string) ([]int, []int) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	left := []int{}
	right := []int{}
	for _, line := range lines  {
		split := strings.Split(line, "   ")
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func part1() int {
	left, right := getLists("./input.txt")
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := range left {
		diff := int(math.Abs(float64(left[i] - right[i])))
		sum += diff
	}
	return sum
}

func part2() int {
	left, right := getLists("./input.txt")
	rightMap := make(map[int]int)
	for _, v := range right {
		rightMap[v] += 1
	}
	sum := 0
	for _, v := range left {
		if count, ok := rightMap[v]; ok {
			sum += v * count
		}
	}
	return sum
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
