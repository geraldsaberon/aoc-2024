package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(inputFile string) []string {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	return lines
}

func check1(lines []string, xrow , xcol int) int {
	count := 0

	rows := len(lines)
	cols := len(lines[0])

	minrow := xrow > 2
	mincol := xcol > 2
	maxrow := rows-xrow > 3
	maxcol := cols-xcol > 3

	directions := map[string][]int{
		"U" :{-1, 0},
		"D" :{+1, 0},
		"L" :{ 0,-1},
		"R" :{ 0,+1},
		"UR":{-1,+1},
		"DR":{+1,+1},
		"UL":{-1,-1},
		"DL":{+1,-1},
	}

	for dir, dirV := range directions {
		if
		dir == "U"  && minrow ||
		dir == "D"  && maxrow ||
		dir == "L"  && mincol ||
		dir == "R"  && maxcol ||
		dir == "UR" && minrow && maxcol ||
		dir == "DR" && maxrow && maxcol ||
		dir == "UL" && minrow && mincol ||
		dir == "DL" && maxrow && mincol {
			x, y := dirV[0], dirV[1]
			m := string(lines[xrow+1*x][xcol+1*y])
			a := string(lines[xrow+2*x][xcol+2*y])
			s := string(lines[xrow+3*x][xcol+3*y])
			if (m == "M" && a == "A" && s == "S") {
				count += 1
			}
		}
	}

	return count
}

func check2(lines []string, arow, acol int) int {
	rows := len(lines)
	cols := len(lines[0])
	count := 0

	minrow := arow > 0
	mincol := acol > 0
	maxrow := rows-arow > 1
	maxcol := cols-acol > 1

	if minrow && mincol && maxrow && maxcol {
		ul := string(lines[arow-1][acol-1])
		ur := string(lines[arow-1][acol+1])
		dl := string(lines[arow+1][acol-1])
		dr := string(lines[arow+1][acol+1])
		if
		(ul == "M" && dr == "S") && (ur == "M" && dl == "S") ||
		(ul == "S" && dr == "M") && (ur == "S" && dl == "M") ||
		(ul == "S" && dr == "M") && (ur == "M" && dl == "S") ||
		(ul == "M" && dr == "S") && (ur == "S" && dl == "M") {
			count += 1
		}
	}

	return count
}

func part1() int {
	lines := readInput("input.txt")
	sum := 0
	for row, line := range lines {
		for col, c := range line {
			char := string(c)
			if char == "X" {
				sum += check1(lines, row, col)
			}
		}
	}
	return sum
}

func part2() int {
	lines := readInput("input.txt")
	sum := 0
	for row, line := range lines {
		for col, c := range line {
			char := string(c)
			if char == "A" {
				sum += check2(lines, row, col)
			}
		}
	}
	return sum
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
