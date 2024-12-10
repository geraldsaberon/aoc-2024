package main

import (
	"bytes"
	"fmt"
	"os"
)

type node struct {
	height byte
	x, y int
}

func readinput(inputfile string) [][]byte {
	input, _ := os.ReadFile(inputfile)
	return bytes.Split(bytes.TrimSpace(input), []byte("\n"))
}

func adjacent(n node, map_ [][]byte) []node {
	adjs := []node{}
	if n.x > 0 {
		adjs = append(adjs, node{map_[n.x-1][n.y], n.x-1, n.y})
	}
	if n.x < len(map_)-1 {
		adjs = append(adjs, node{map_[n.x+1][n.y], n.x+1, n.y})
	}
	if n.y > 0 {
		adjs = append(adjs, node{map_[n.x][n.y-1], n.x, n.y-1})
	}
	if n.y < len(map_[0])-1 {
		adjs = append(adjs, node{map_[n.x][n.y+1], n.x, n.y+1})
	}
	return adjs
}

func search(head node, map_ [][]byte) (int, int) {
	stack := []node{head}
	trails := map[[2]int]bool{}
	rating := 0
	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if last.height == '9' {
			trails[[2]int{last.x, last.y}] = true
			rating += 1
		} else if adjs := adjacent(last, map_); len(adjs) > 0 {
			for _, a := range adjs {
				if a.height == last.height+1 {
					stack = append(stack, a)
				}
			}
		}
	}
	return len(trails), rating
}

func part1and2() (int, int) {
	map_ := readinput("input.txt")
	scoreSum := 0
	ratingSum := 0
	for x := range map_ {
		for y, height := range map_[x] {
			if height == '0' {
				score, rating := search(node{height, x, y}, map_)
				scoreSum += score
				ratingSum += rating
			}
		}
	}
	return scoreSum, ratingSum
}

func main() {
	p1, p2 := part1and2()
    fmt.Println("Part 1:", p1)
    fmt.Println("Part 2:", p2)
}
