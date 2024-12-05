package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func readInput(inputFile string) (map[int][]int, [][]int) {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	rules := map[int][]int{}
	updates := [][]int{}
	parseRules := true
	for _, line := range lines {
		if len(line) == 0 {
			parseRules = false
			continue
		}
		if parseRules {
			s := strings.Split(line, "|")
			l, _ := strconv.Atoi(s[0])
			r, _ := strconv.Atoi(s[1])
			rules[r] = append(rules[r], l)
		} else {
			pages := strings.Split(line, ",")
			intpages := []int{}
			for _, s := range pages {
				i, _ := strconv.Atoi(s)
				intpages = append(intpages, i)
			}
			updates = append(updates, intpages)
		}
	}
	return rules, updates
}

func part1() int {
	rules, updates := readInput("input.txt")
	sum := 0
	for u, update := range updates {
		ok := true
		for p, page := range update {
			if !ok { break }
			for _, pageAhead := range update[p:] {
				if !ok { break }
				for _, pageAheadRule := range rules[page] {
					if pageAhead == pageAheadRule {
						ok = false
						break
					}
				}
			}
		}
		if ok {
			middle := len(updates[u]) / 2
			sum += updates[u][middle]
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
