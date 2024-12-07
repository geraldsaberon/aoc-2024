package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	row, col int
	dir rune // ^ V < >
}

var directions = map[rune][2]int{
	'^': {-1, 0},
	'<': { 0,-1},
	'V': {+1, 0},
	'>': { 0,+1},
}

var turns = map[rune]rune{
	'^': '>',
	'<': '^',
	'V': '<',
	'>': 'V',
}

func readInput(inputFile string) ([][]rune, Position) {
	input, _ := os.ReadFile(inputFile)
	inputlines := strings.Split(strings.TrimSpace(string(input)), "\n")
	lines := [][]rune{}
	guardpos := Position{}
	for row, line := range inputlines {
		lines = append(lines, []rune(line))
		for col, char := range line {
			if char == '^' {
				guardpos.row = row
				guardpos.col = col
				guardpos.dir = '^'
			}
		}
	}
	return lines, guardpos
}

func nextmove(g *Position, lines [][]rune) bool {
	rows, cols := len(lines), len(lines[0])

	if
	g.dir == '^' && g.row == 0 ||
	g.dir == '<' && g.col == 0 ||
	g.dir == 'V' && g.row == rows-1 ||
	g.dir == '>' && g.col == cols-1  {
		return false
	}

	for dir, move := range directions {
		x, y := move[0], move[1]
		if g.dir == dir {
			for lines[g.row+1*x][g.col+1*y] == '#' {
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

func runguard(g *Position, lines [][]rune) (map[[2]int]bool, bool) {
	visited := map[[2]int]bool{}
	visits := map[[3]int]int{}
	hasloop := false
	for ok := true; ok && !hasloop; ok = nextmove(g, lines) {
		key := [3]int{g.row, g.col, int(g.dir)}
		if visitcount, _ := visits[key]; visitcount > 0 {
			hasloop = true
		} else {
			visits[key] += 1
			visited[[2]int{g.row, g.col}] = true
		}
	}
	return visited, hasloop
}

func part1and2() (int, int) {
	lines, gp := readInput("input.txt")
	visited, _ := runguard(&Position{row: gp.row, col: gp.col, dir: gp.dir}, lines)
	obstructions := map[[2]int]bool{}
	for coord := range visited {
		x, y := coord[0], coord[1]
		if lines[x][y] != '^' {
			lines[x][y] = '#'
			_, hasloop := runguard(&Position{row: gp.row, col: gp.col, dir: gp.dir}, lines)
			if hasloop {
				_, seen := obstructions[coord]
				if !seen {
					obstructions[coord] = true
				}
			}
			lines[x][y] = '.'
		}
	}
	return len(visited), len(obstructions)
}

func main() {
	p1, p2 := part1and2()
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
