package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"slices"
)

func readInput(inputFile string) (map[int][]int, [][]int) {
	input, _ := os.ReadFile(inputFile)
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	rules := map[int][]int{}
	updates := [][]int{}
	for _, rule := range strings.Split(split[0], "\n") {
		s := strings.Split(rule, "|")
		l, _ := strconv.Atoi(s[0])
		r, _ := strconv.Atoi(s[1])
		rules[r] = append(rules[r], l)
	}
	for _, update := range strings.Split(split[1], "\n") {
		pages := strings.Split(update, ",")
		intpages := []int{}
		for _, s := range pages {
			i, _ := strconv.Atoi(s)
			intpages = append(intpages, i)
		}
		updates = append(updates, intpages)
	}
	return rules, updates
}

func checkValid(pageIndex int, update []int, rules map[int][]int) (bool, int) {
	for p, pageAhead := range update[pageIndex:] {
		if slices.Contains(rules[update[pageIndex]], pageAhead) {
			return false, p+pageIndex
		}
	}
	return true, -1
}

func part1and2() (int, int) {
	rules, updates := readInput("input.txt")
	oksum, notoksum := 0, 0
	for _, update := range updates {
		ok := true
		for pageIndex := range update {
			if ok, _ = checkValid(pageIndex, update, rules); !ok {
				break
			}
		}
		if ok {
			oksum += update[len(update)/2]
		} else {
			for pageIndex := 0; pageIndex < len(update); {
				if ok, i := checkValid(pageIndex, update, rules); ok {
					pageIndex += 1
				} else {
					update[pageIndex], update[i] = update[i], update[pageIndex]
				}
			}
			notoksum += update[len(update)/2]
		}
	}
	return oksum, notoksum
}

func main() {
	p1, p2 := part1and2()
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
