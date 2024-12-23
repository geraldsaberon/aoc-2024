package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func readInput(inputFile string) [][2]string {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	connections := [][2]string{}
	for _, line := range lines {
		connections = append(connections, [2]string(strings.Split(line, "-")))
	}
	return connections
}

func areAdjacent(g map[string][]string, a, b string) bool {
	if slices.Contains(g[a], b) || slices.Contains(g[b], a) {
		return true
	}
	return false
}

func part1() int {
	connections := readInput("input.txt")
	g := map[string][]string{}
	for _, con := range connections {
		a, b := con[0], con[1]
		g[a] = append(g[a], b)
	}
	threes := []string{}
	for    a := range g {
	for _, b := range g[a] {
	for _, c := range g[b] {
		if areAdjacent(g, a, b) && areAdjacent(g, b, c) && areAdjacent(g, a, c) {
			t := []string{a, b, c}
			slices.Sort(t)
			st := fmt.Sprint(t)
			if !slices.Contains(threes, st) && (a[0] == 't' || b[0] == 't' || c[0] == 't') {
				threes = append(threes, st)
			}
		}
	}}}
	return len(threes)
}

func main() {
	fmt.Println("Part 1:", part1())
}
