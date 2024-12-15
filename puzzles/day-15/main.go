package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct { x, y int }

var directions = map[rune]position {
	'^': { 0,-1},
	'v': { 0,+1},
	'<': {-1, 0},
	'>': {+1, 0},
}

func readInput(inputFile string) ([][]string, string, position) {
	input, _ := os.ReadFile(inputFile)
	split := strings.Split(string(input), "\n\n")
	map_ := [][]string{}
	for _, line := range strings.Split(split[0], "\n") {
		map_ = append(map_, strings.Split(line, ""))
	}
	robotPos := position{}
	for y := range map_ {
		for x, s := range map_[y] {
			if s == "@" {
				robotPos.x, robotPos.y = x, y
			}
		}
	}
	moves := split[1]
	return map_, moves, robotPos
}

func move(obj *position, dir position, map_ [][]string) bool {
	nx, ny := obj.x+1*dir.x, obj.y+1*dir.y
	nextPosition := map_[ny][nx]
	if nextPosition == "." {
		map_[ny][nx], map_[obj.y][obj.x] = map_[obj.y][obj.x], map_[ny][nx]
		obj.x, obj.y = nx, ny
		return true
	}
	if nextPosition == "O" {
		hasMoved := move(&position{nx, ny}, dir, map_)
		if hasMoved {
			map_[ny][nx], map_[obj.y][obj.x] = map_[obj.y][obj.x], map_[ny][nx]
			obj.x, obj.y = nx, ny
			return true
		}
	}
	return false
}

func getBoxGPSCoordsSum(map_ [][]string) int {
	sum := 0
	for y := range map_ {
		for x, s := range map_[y] {
			if s == "O" {
				sum += 100 * y + x
			}
		}
	}
	return sum
}

func part1() int {
	map_, moves, robotPos := readInput("input.txt")
	for _, m := range moves {
		if m == '\n' {
			continue
		}
		move(&robotPos, directions[m], map_)
	}
	return getBoxGPSCoordsSum(map_)
}

func main() {
	fmt.Println("Part 1:", part1())
}
