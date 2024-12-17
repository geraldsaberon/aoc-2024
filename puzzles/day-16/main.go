package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type PriorityQueue []*node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

type position struct {
	x, y int
}

func (p position) adjacents() []position {
	adjs := []position{
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
	}
	return adjs
}

func (p position) isValid(maze [][]string) bool {
	return maze[p.y][p.x] != "#"
}

func (p position) isEndGoal(maze [][]string) bool {
	return maze[p.y][p.x] == "E"
}

func (p position) isStartPos(maze [][]string) bool {
	return maze[p.y][p.x] == "S"
}

func (p position) diff(pd position) position {
	return position{p.x - pd.x, p.y - pd.y}
}

type node struct {
	pos   position
	prev  *node
	score int
	dir   position
	index int
}

type seenKey struct {
	p, d position
}

func readInput(inputFile string) ([][]string, position) {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	maze := [][]string{}
	for _, line := range lines {
		maze = append(maze, strings.Split(line, ""))
	}
	start := position{}
	for y, row := range maze {
		for x, c := range row {
			if c == "S" {
				start.x, start.y = x, y
			}
		}
	}
	return maze, start
}

func getBestPathsTiles(ends []*node, maze [][]string, bestScore int) int {
	tiles := map[position]bool{}
	for _, node := range ends {
		if node.score > bestScore {
			continue
		}
		for {
			tiles[node.pos] = true
			if node.pos.isStartPos(maze) {
				break
			}
			node = node.prev
		}
	}
	return len(tiles)
}

func part1and2() (int, int) {
	maze, start := readInput("input.txt")
	ends := []*node{}
	seen := map[seenKey]bool{}
	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &node{pos: start, prev: &node{}, score: 0, dir: position{1, 0}})
	bestScore := 999999
	for pq.Len() > 0 {
		curNode := heap.Pop(&pq).(*node)
		seen[seenKey{curNode.pos, curNode.dir}] = true
		if curNode.pos.isEndGoal(maze) {
			if curNode.score < bestScore {
				bestScore = curNode.score
			}
			ends = append(ends, curNode)
			continue
		}
		for _, a := range curNode.pos.adjacents() {
			if !a.isValid(maze) {
				continue
			}
			dir := a.diff(curNode.pos)
			score := curNode.score + 1
			if dir != curNode.dir {
				score += 1000
			}
			adjNode := node{pos: a, prev: curNode, score: score, dir: dir}
			if _, ok := seen[seenKey{adjNode.pos, adjNode.dir}]; !ok {
				heap.Push(&pq, &adjNode)
			}
		}
	}
	tiles := getBestPathsTiles(ends, maze, bestScore)
	return bestScore, tiles
}

func main() {
	p1, p2 := part1and2()
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
