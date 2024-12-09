package main

import (
	"fmt"
	"os"
	"strconv"
)

type block struct {
	length, free int
}

type file struct {
	index, length int
}

func readinput(inputfile string) []block {
	input, _ := os.ReadFile(inputfile)
	stringinput := string(input)
	diskmap := []block{}
	for i := 0; i < len(stringinput)-1; i+=2 {
		length, _ := strconv.Atoi(string(stringinput[i]))
		free, _ := strconv.Atoi(string(stringinput[i+1]))
		diskmap = append(diskmap, block{length, free})
	}
	return diskmap
}

func blocking(diskmap []block) []string {
	fileblocks := []string{}
	for id, block := range diskmap {
		for i := 0; i < block.length; i++ {
			fileblocks = append(fileblocks, strconv.Itoa(id))
		}
		for i := 0; i < block.free; i++ {
			fileblocks = append(fileblocks, ".")
		}
	}
	return fileblocks
}

func compact(fileblocks []string) []string {
	l, r := 0, len(fileblocks)-1
	for l < r {
		for i := r; fileblocks[r] == "."; i-- {
			r = i
		}
		if fileblocks[l] == "." {
			fileblocks[l], fileblocks[r] = fileblocks[r], fileblocks[l]
		}
		l += 1
	}
	return fileblocks
}

func getfiles(fileblocks []string) []file {
	files := make([]file, len(fileblocks))
	for i, b := range fileblocks {
		if b != "." {
			intb, _ := strconv.Atoi(b)
			files[intb].length += 1
			files[intb].index = i-files[intb].length+1
		}
	}
	return files
}

func compact2(fileblocks []string) []string {
	files := getfiles(fileblocks)
	for i := len(files)-1; i >= 0; i-- {
		file := files[i]
		if file.length == 0 {
			continue
		}
		freespace := 0
		for i, b := range fileblocks {
			if b == "." {
				freespace += 1
			} else if f := i-freespace; freespace >= file.length && f < file.index {
				for offset := 0; offset < file.length; offset++ {
					fileblocks[f+offset], fileblocks[file.index+offset] = fileblocks[file.index+offset], fileblocks[f+offset]
				}
				freespace = 0
				break
			} else {
				freespace = 0
			}
		}
	}
	return fileblocks
}

func checksum(compacted []string) int {
	sum := 0
	for pos, id := range compacted {
		if id != "." {
			idnum, _ := strconv.Atoi(id)
			sum += pos * idnum
		}
	}
	return sum
}

func part1() int {
	diskmap := readinput("input.txt")
	fileblocks := blocking(diskmap)
	compacted := compact(fileblocks)
	sum := checksum(compacted)
	return sum
}

func part2() int {
	diskmap := readinput("input.txt")
	fileblocks := blocking(diskmap)
	compacted := compact2(fileblocks)
	sum := checksum(compacted)
	return sum
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
