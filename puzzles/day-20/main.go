package main

import (
	"fmt"
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

func (p position) adjacents() []position {
	adjs := []position{
		{p.x, p.y-1},
		{p.x, p.y+1},
		{p.x-1, p.y},
		{p.x+1, p.y},
	}
	return adjs
}

func (p position) isInBounds(map_ [][]string) bool {
	if p.x < 0 || p.y < 0 || p.x > len(map_[0])-1 || p.y > len(map_)-1 {
		return false
	}
	return true
}

func (p position) isEdge(map_ [][]string) bool {
	if p.x == 0 || p.y == 0 || p.x == len(map_[0])-1 || p.y == len(map_)-1 {
		return true
	}
	return false
}

func (p position) isPassable(map_ [][]string) bool {
	a := p.adjacents()
	u, d, l, r := a[0], a[1], a[2], a[3]
	if !l.isInBounds(map_) || !r.isInBounds(map_) || !u.isInBounds(map_) || !d.isInBounds(map_) {
		return false
	}
	if map_[l.y][l.x] == "." && map_[r.y][r.x] == "." {
		return true
	}
	if map_[u.y][u.x] == "." && map_[d.y][d.x] == "." {
		return true
	}
	return false
}

func readInput(inputFile string) (map_ [][]string, walls []position, start, end position) {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	map_ = [][]string{}
	walls = []position{}
	start = position{}
	end = position{}
	for y, row := range lines {
		map_ = append(map_, strings.Split(row, ""))
		for x, tile := range row {
			if tile == 'S' {
				start.x, start.y = x, y
				map_[y][x] = "."
			} else if tile == 'E' {
				end.x, end.y = x, y
				map_[y][x] = "."
			} else if tile == '#' {
				walls = append(walls, position{x, y})
			}
		}
	}
	return
}

func bfs(start, end position, map_ [][]string) int {
	queue := []*node{{start, 0}}
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
			if _, visited := seen[a]; !visited && a.isInBounds(map_) && map_[a.y][a.x] != "#" {
				queue = append(queue, &node{a, cur.steps+1})
			}
		}
	}
	return solution.steps
}

func part1() int {
	map_, walls, start, end := readInput("input.txt")
	baseTime := bfs(start, end, map_)
	total := 0
	for _, wall := range walls {
		if !wall.isEdge(map_) && wall.isPassable(map_) {
			map_[wall.y][wall.x] = "."
			t := bfs(start, end, map_)
			ts := baseTime - t
			if ts >= 100 {
				total += 1
			}
		}
		map_[wall.y][wall.x] = "#"
	}
	return total
}

func main() {
	fmt.Println("Part 1:", part1())
}
