package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type position struct {
	x, y int
}

func readinput(inputfile string) (map[rune][]position, [2]int) {
	input, _ := os.ReadFile(inputfile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	antennas := map[rune][]position{}
	for x, line := range lines {
		for y, c := range line {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				antennas[c] = append(antennas[c], position{x, y})
			}
		}
	}
	return antennas, [2]int{len(lines), len(lines[0])}
}

func validnode(node position, mapsize [2]int) bool {
	w, h := mapsize[0], mapsize[1]
	return (node.x >= 0 && node.x < w) && (node.y >= 0 && node.y < h)
}

func antinodes(locations []position) []position {
	nodes := []position{}
	for _, a := range locations {
	for _, b := range locations {
		if a == b || b.x < a.x { continue }
		dx, dy := b.x - a.x, b.y - a.y
		node1 := position{a.x-dx, a.y-dy}
		node2 := position{b.x+dx, b.y+dy}
		nodes = append(nodes, node1, node2)
	}}
	return nodes
}

func antinodes2(locations []position, mapsize [2]int) []position {
	nodes := []position{}
	w, h := mapsize[0], mapsize[1]
	for _, a := range locations {
	for _, b := range locations {
		if a == b || b.x < a.x { continue }
		dx, dy := b.x - a.x, b.y - a.y
		for i := 0; dx*i < w && dy*i < h; i++ {
			node1 := position{a.x-dx*i, a.y-dy*i}
			node2 := position{b.x+dx*i, b.y+dy*i}
			nodes = append(nodes, node1, node2)
		}
	}}
	return nodes
}

func part1() int {
	antennas, mapsize := readinput("input.txt")
	unique := map[position]bool{}
	for _, locations := range antennas {
		nodes := antinodes(locations)
		for _, node := range nodes {
			if validnode(node, mapsize) {
				unique[node] = true
			}
		}
	}
	return len(unique)
}

func part2() int {
	antennas, mapsize := readinput("input.txt")
	unique := map[position]bool{}
	for _, locations := range antennas {
		nodes := antinodes2(locations, mapsize)
		for _, node := range nodes {
			if validnode(node, mapsize) {
				unique[node] = true
			}
		}
	}
	return len(unique)
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
