package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func readInput(inputFile string) map[string][]string {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	graph := map[string][]string{}
	for _, line := range lines {
		split := strings.Split(line, "-")
		a, b := split[0], split[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	return graph
}

func areAdjacent(g map[string][]string, a, b string) bool {
	if slices.Contains(g[a], b) || slices.Contains(g[b], a) {
		return true
	}
	return false
}

func part1() int {
	graph := readInput("input.txt")
	threes := map[string]bool{}
	for    a := range graph {
	for _, b := range graph[a] {
	for _, c := range graph[a] {
		if areAdjacent(graph, b, c) && (a[0] == 't' || b[0] == 't' || c[0] == 't') {
			t := []string{a, b, c}
			slices.Sort(t)
			threes[fmt.Sprint(t)] = true
		}
	}}}
	return len(threes)
}

func main() {
	fmt.Println("Part 1:", part1())
}
