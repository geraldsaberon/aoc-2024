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

func widenMap(map_ [][]string) ([][]string, position) {
	newMap := [][]string{}
	robotPos := position{}
	for y := range map_ {
		row := []string{}
		for x, s := range map_[y] {
			if s == "#" {
				row = append(row, "#", "#")
			} else if s == "O" {
				row = append(row, "[", "]")
			} else if s == "." {
				row = append(row, ".", ".")
			} else {
				row = append(row, "@", ".")
				robotPos.x, robotPos.y = x*2, y
			}
		}
		newMap = append(newMap, row)
	}
	return newMap, robotPos
}

func move(obj position, dir position, map_ [][]string) bool {
	np := position{obj.x+dir.x, obj.y+dir.y}
	npSpot := map_[np.y][np.x]
	if npSpot == "." {
		map_[np.y][np.x], map_[obj.y][obj.x] = map_[obj.y][obj.x], map_[np.y][np.x]
		return true
	}
	if npSpot == "O" || npSpot == "[" || npSpot == "]" {
		if hasMoved := move(position{np.x, np.y}, dir, map_); hasMoved {
			map_[np.y][np.x], map_[obj.y][obj.x] = map_[obj.y][obj.x], map_[np.y][np.x]
			return true
		}
	}
	return false
}

func copyMap(map_ [][]string) [][]string {
	sCopy := make([][]string, len(map_))
	for i := range map_ {
		sCopy[i] = make([]string, len(map_[0]))
		copy(sCopy[i], map_[i])
	}
	return sCopy
}

func swap(curr, next position, map_ [][]string) {
	map_[curr.y][curr.x]   = "."
	map_[curr.y][curr.x+1] = "."
	map_[next.y][next.x]   = "["
	map_[next.y][next.x+1] = "]"
}

func moveBoxUD(obj position, dir position, map_ *[][]string) bool {
	prevMap := copyMap(*map_)

	np := position{obj.x+dir.x, obj.y+dir.y}
	L := (*map_)[np.y][np.x]
	R := (*map_)[np.y][np.x+1]

	if L == "." && R == "." {
		swap(obj, np, *map_)
		return true
	}
	if L == "]" && R == "." {
		hasMoved := moveBoxUD(position{np.x-1, np.y}, dir, map_)
		if hasMoved {
			swap(obj, np, *map_)
			return true
		}
	}
	if L == "." && R == "[" {
		hasMoved := moveBoxUD(position{np.x+1, np.y}, dir, map_)
		if hasMoved {
			swap(obj, np, *map_)
			return true
		}
	}
	if L == "[" && R == "]" {
		hasMoved := moveBoxUD(position{np.x, np.y}, dir, map_)
		if hasMoved {
			swap(obj, np, *map_)
			return true
		}
	}
	if L == "]" && R == "[" {
		hasMovedL := moveBoxUD(position{np.x-1, np.y}, dir, map_)
		hasMovedR := moveBoxUD(position{np.x+1, np.y}, dir, map_)
		if hasMovedL && hasMovedR {
			swap(obj, np, *map_)
			return true
		}
	}
	*map_ = prevMap
	return false
}

func getBoxGPSCoordsSum(map_ [][]string) int {
	sum := 0
	for y := range map_ {
		for x, s := range map_[y] {
			if s == "O" || s == "[" {
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
		dir := directions[m]
		if hasMoved := move(robotPos, dir, map_); hasMoved {
			robotPos.x += dir.x
			robotPos.y += dir.y
		}
	}
	return getBoxGPSCoordsSum(map_)
}

func moveRobot(rp *position, dir position, map_ [][]string) {
	map_[rp.y][rp.x] = "."
	rp.x += dir.x
	rp.y += dir.y
	map_[rp.y][rp.x] = "@"
}

func part2() int {
	map_, moves, rp := readInput("input.txt")
	map_, rp = widenMap(map_)
	for _, m := range moves {
		if m == '\n' {
			continue
		}
		dir := directions[m]
		np := position{rp.x+dir.x, rp.y+dir.y}
		npSpot := map_[np.y][np.x]
		if npSpot == "." {
			moveRobot(&rp, dir, map_)
		} else if npSpot == "[" && (m == '^' || m == 'v') {
			hasMoved := moveBoxUD(np, dir, &map_)
			if hasMoved {
				moveRobot(&rp, dir, map_)
			}
		} else if npSpot == "]" && (m == '^' || m == 'v') {
			hasMoved := moveBoxUD(position{np.x-1, np.y}, dir, &map_)
			if hasMoved {
				moveRobot(&rp, dir, map_)
			}
		} else if (npSpot == "[" || npSpot == "]") && (m == '<' || m == '>') {
			hasMoved := move(np, dir, map_)
			if hasMoved {
				moveRobot(&rp, dir, map_)
			}
		}
	}
	return getBoxGPSCoordsSum(map_)
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
