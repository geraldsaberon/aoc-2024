package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	pos position
	prev *node
	steps int
}

type position struct {
	x, y int
}

func (p position) adjacents() []position {
	adjs := []position{
		{p.x, p.y-1},
		{p.x, p.y+1},
		{p.x-1, p.y},
		{p.x+1, p.y},
	}
	return adjs
}

func (p position) isValid(map_ [][]string) bool {
	if p.y < 0 || p.x < 0 || p.y >= len(map_) || p.x >= len(map_[0]) {
		return false
	}
	if map_[p.y][p.x] == "#" {
		return false
	}
	return true
}

func readInput(inputFile string) []position {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ps := []position{}
	for _, line := range lines {
		xy := strings.Split(line, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		ps = append(ps, position{x, y})
	}
	return ps
}

func makeMap(w, h int) [][]string {
	map_ := make([][]string, h)
	for y := range map_ {
		map_[y] = make([]string, w)
		for x := range map_[y] {
			map_[y][x] = "."
		}
	}
	return map_
}

func bfs(start, end position, map_ [][]string) int {
	queue := []*node{{start, &node{}, 0}}
	seen := map[position]bool{}
	solution := node{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.pos == end {
			solution = *cur
			break
		}
		if _, visited := seen[cur.pos]; visited {
			continue
		}
		seen[cur.pos] = true
		for _, a := range cur.pos.adjacents() {
			if _, visited := seen[a]; !visited && a.isValid(map_) {
				queue = append(queue, &node{a, cur, cur.steps+1})
			}
		}
	}
	return solution.steps
}

func part1() int {
	ps := readInput("input.txt")
	w, h, max := 71, 71, 1024
	map_ := makeMap(w, h)
	for i := 0; i < max; i++ {
		map_[ps[i].y][ps[i].x] = "#"
	}
	return bfs(position{0, 0}, position{70, 70}, map_)
}

func part2() position {
	ps := readInput("input.txt")
	w, h, max := 71, 71, len(ps)
	map_ := makeMap(w, h)
	for i := 0; i < max; i++ {
		map_[ps[i].y][ps[i].x] = "#"
		res := bfs(position{0, 0}, position{70, 70}, map_)
		if res <= 0 {
			return ps[i]
		}
	}
	return position{}
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
