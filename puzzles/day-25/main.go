package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(inputFile string) (locks [][5]int, keys [][5]int) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	lockAndKeys := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	for _, item := range lockAndKeys {
		item = strings.Replace(item, "\n", "", -1)
		if item[0] == '#' {
			lock := [5]int{-1,-1,-1,-1,-1}
			for i, c := range item {
				if c == '#' {
					lock[i%5] += 1
				}
			}
			locks = append(locks, lock)
		} else if item[0] == '.' {
			key := [5]int{-1,-1,-1,-1,-1}
			for i, c := range item {
				if c == '#' {
					key[i%5] += 1
				}
			}
			keys = append(keys, key)
		}
	}
	return
}

func part1() int {
	locks, keys := readInput("input.txt")
	fits := 0
	for _, lock := range locks {
		for _, key := range keys {
			fit := true
			for i := 0; i < 5; i++ {
				if lock[i] + key[i] > 5 {
					fit = false
					break
				}
			}
			if fit {
				fits += 1
			}
		}
	}
	return fits
}

func main() {
	fmt.Println("Part 1:", part1())
}
