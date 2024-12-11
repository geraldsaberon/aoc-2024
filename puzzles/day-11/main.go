package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readinput(inputfile string) map[int]int {
	input, _ := os.ReadFile(inputfile)
	stones := map[int]int{}
	for _, ns := range strings.Split(strings.TrimSpace(string(input)), " ") {
		n, _ := strconv.Atoi(ns)
		stones[n] += 1
	}
	return stones
}

func split(n int) (int, int) {
	ns := strconv.Itoa(n)
	ls, rs := ns[:len(ns)/2], ns[len(ns)/2:]
	l, _ := strconv.Atoi(ls)
	r, _ := strconv.Atoi(rs)
	return l, r
}

func intlen(n int) int {
	if n == 0 {
		return 1
	}
	l := 0
	for n > 0 {
		l += 1
		n /= 10
	}
	return l
}

func blink(stones map[int]int) {
	new := map[int]int{}
	for stone, count := range stones {
		if count <= 0 {
			continue
		}
		if stone == 0 {
			new[1] += 1 * count
		} else if intlen(stone)%2 == 0 {
			l, r := split(stone)
			new[l] += 1 * count
			new[r] += 1 * count
		} else {
			new[stone*2024] += 1 * count
		}
		stones[stone] -= 1 * count
	}
	for stone, count := range new {
		stones[stone] += count
	}
}

func part1() int {
	stones := readinput("input.txt")
	for i := 0; i < 25; i++ {
		blink(stones)
	}
	total := 0
	for _, count := range stones {
		total += count
	}
	return total
}

func part2() int {
	stones := readinput("input.txt")
	for i := 0; i < 75; i++ {
		blink(stones)
	}
	total := 0
	for _, count := range stones {
		total += count
	}
	return total
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
