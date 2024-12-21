package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type node struct {
	pos position
	steps int
}

type position struct {
	x, y int
}

func (p position) distance(q position) int {
	dx := math.Abs(float64(p.x - q.x))
	dy := math.Abs(float64(p.y - q.y))
	return int(dx + dy)
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

func readInput(inputFile string) (position, position, [][]string) {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	map_ := [][]string{}
	start, end := position{}, position{}
	for y, row := range lines {
		map_ = append(map_, strings.Split(row, ""))
		for x, tile := range row {
			if tile == 'S' {
				start.x, start.y = x, y
				map_[y][x] = "."
			} else if tile == 'E' {
				end.x, end.y = x, y
				map_[y][x] = "."
			}
		}
	}
	return start, end, map_
}

func bfs(start, end position, map_ [][]string) []*node {
	queue := []*node{{start, 0}}
	seen := map[position]bool{}
	path := []*node{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		path = append(path, cur)
		if cur.pos == end {
			break
		}
		seen[cur.pos] = true
		for _, a := range cur.pos.adjacents() {
			if _, visited := seen[a]; !visited && map_[a.y][a.x] != "#" {
				queue = append(queue, &node{a, cur.steps+1})
			}
		}
	}
	return path
}

func cheatCount(inputFile string, cheatTime, minTimeSaved int) int {
	start, end, map_ := readInput(inputFile)
	path := bfs(start, end, map_)
	count := 0
	for i, p := range path {
		for _, q := range path[i+1:] {
			d := p.pos.distance(q.pos)
			ts := q.steps - p.steps - d
			if d <= cheatTime && ts >= minTimeSaved {
				count += 1
			}
		}
	}
	return count
}

func part1() int {
	return cheatCount("input.txt", 2, 100)
}

func part2() int {
	return cheatCount("input.txt", 20, 100)
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
