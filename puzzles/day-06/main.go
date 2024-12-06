package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	row, col int
	dir string // ^ V < >
}

func readInput(inputFile string) ([][]string, Position) {
	input, _ := os.ReadFile(inputFile)
	lines := [][]string{}
	guardpos := Position{}
	for row, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		chars := strings.Split(line, "")
		lines = append(lines, chars)
		for col, char := range chars {
			if char == "^" {
				guardpos.row = row
				guardpos.col = col
				guardpos.dir = "^"
			}
		}
	}
	return lines, guardpos
}

func nextmove(g *Position, lines [][]string) bool {
	block := "#"
	rows, cols := len(lines), len(lines[0])

	if
	g.dir == "^" && g.row == 0 ||
	g.dir == "<" && g.col == 0 ||
	g.dir == "V" && g.row == rows-1 ||
	g.dir == ">" && g.col == cols-1  {
		return false
	}

	directions := map[string][2]int{
		"^": {-1, 0},
		"<": { 0,-1},
		"V": {+1, 0},
		">": { 0,+1},
	}

	turns := map[string]string{
		"^": ">",
		"<": "^",
		"V": "<",
		">": "V",
	}

	for dir, move := range directions {
		x, y := move[0], move[1]
		if g.dir == dir {
			if lines[g.row+1*x][g.col+1*y] == block {
				turn := turns[g.dir]
				move = directions[turn]
				x, y = move[0], move[1]
				g.dir = turn
			}
			g.row += 1 * x
			g.col += 1 * y
			return true
		}
	}
	return false
}

func part1() int {
	lines, guardpos := readInput("input.txt")
	visited := map[[2]int]bool{}
	for ok := nextmove(&guardpos, lines); ok; ok = nextmove(&guardpos, lines) {
		key := [2]int{guardpos.row, guardpos.col}
		if _, ok := visited[key]; !ok {
			visited[key] = true
		}
	}
	return len(visited)
}

func part2() int {
	return -1
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
