package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

type position struct {
	x, y int
}

func readinput(inputfile string) [][]byte {
	input, _ := os.ReadFile(inputfile)
	return bytes.Split(bytes.TrimSpace(input), []byte("\n"))
}

func perimeter(ps []position) int {
	adjacent := 0
    for _, a := range ps {
    for _, b := range ps {
        dx := math.Abs(float64(a.x-b.x))
        dy := math.Abs(float64(a.y-b.y))
        if dx + dy == 1 {
            adjacent += 1
        }
    }}
    perimeter := len(ps)*4 - adjacent
    return perimeter
}

func adjacent(p position, map_ [][]byte) []position {
	adjs := []position{}
	if p.x > 0 {
		adjs = append(adjs, position{p.x-1, p.y})
	}
	if p.x < len(map_)-1 {
		adjs = append(adjs, position{p.x+1, p.y})
	}
	if p.y > 0 {
		adjs = append(adjs, position{p.x, p.y-1})
	}
	if p.y < len(map_[0])-1 {
		adjs = append(adjs, position{p.x, p.y+1})
	}
	return adjs
}

func getregions(map_ [][]byte) [][]position {
	seen := map[position]bool{}
	regions := [][]position{}
	for x, row := range map_ {
		for y, plot := range row {
			p := position{x, y}
			stack := []position{p}
			region := []position{}
			for len(stack) > 0 {
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if _, visited := seen[last]; visited {
					continue
				}
				seen[last] = true
				region = append(region, last)
				for _, adj := range adjacent(last, map_) {
					if map_[adj.x][adj.y] == plot {
						stack = append(stack, adj)
					}
				}
			}
			regions = append(regions, region)
		}
	}
	return regions
}

func part1() int {
	map_ := readinput("input.txt")
	regions := getregions(map_)
	total := 0
	for _, region := range regions {
		total += len(region) * perimeter(region)
	}
	return total
}

func part2() int {
	return -1
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
