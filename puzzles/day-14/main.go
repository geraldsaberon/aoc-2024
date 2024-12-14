package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func readInput(inputFile string) [][2]coord {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	config := [][2]coord{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		position := strings.Split(strings.TrimLeft(split[0], "p="), ",")
		velocity := strings.Split(strings.TrimLeft(split[1], "v="), ",")
		positionX, _ := strconv.Atoi(position[0])
		positionY, _ := strconv.Atoi(position[1])
		velocityX, _ := strconv.Atoi(velocity[0])
		velocityY, _ := strconv.Atoi(velocity[1])
		config = append(config, [2]coord{{positionX, positionY}, {velocityX, velocityY}})
	}
	return config
}

func makeMap(w, h int) [][]int {
	map_ := [][]int{}
	for i := 0; i < h; i++ {
		rows := make([]int, w)
		map_ = append(map_, rows)
	}
	return map_
}

func mod(a, b int) int {
    return (a % b + b) % b
}

func run(p, v coord, w, h int, time int) coord {
	x := mod(((time * v.x) + p.x), w)
	y := mod(((time * v.y) + p.y), h)
	return coord{x, y}
}

func getSafetyFactor(map_ [][]int) int {
	w, h := len(map_[0]), len(map_)
	quadrants := [4]int{}
	for y := range map_ {
		for x, n := range map_[y] {
			if x < w/2 && y < h/2 { quadrants[0] += n }
			if x > w/2 && y < h/2 { quadrants[1] += n }
			if x < w/2 && y > h/2 { quadrants[2] += n }
			if x > w/2 && y > h/2 { quadrants[3] += n }
		}
	}
	safetyFactor := 1
	for _, quadrant := range quadrants {
		if quadrant > 0 {
			safetyFactor *= quadrant
		}
	}
	return safetyFactor
}

func part1() int {
	config := readInput("input.txt")
	w, h := 101, 103
	time := 100
	map_ := makeMap(w, h)
	for _, pv := range config {
		p, v := pv[0], pv[1]
		pnew := run(p, v, w, h, time)
		map_[pnew.y][pnew.x] += 1
	}
	return getSafetyFactor(map_)
}

func main() {
	fmt.Println("Part 1:", part1())
}
