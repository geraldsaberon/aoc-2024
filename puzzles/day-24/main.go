package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(inputFile string) (map[string]int, []string) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	values := map[string]int{}
	for _, line := range strings.Split(split[0], "\n") {
		s := strings.Split(line, ": ")
		i, _ := strconv.Atoi(s[1])
		values[s[0]] = i
	}
	operations := strings.Split(split[1], "\n")
	return values, operations
}

func part1() int {
	inputs, operations := readInput("input.txt")
	done := map[string]bool{}
	for i := 0; len(done) < len(operations); i++ {
		op := operations[i % len(operations)]
		if done[op] {
			continue
		}
		s := strings.Fields(strings.Replace(op, "-> ", "", 1))
		L, O, R, OUT := s[0], s[1], s[2], s[3]
		LV, okL := inputs[L]
		RV, okR := inputs[R]
		if !(okL && okR) {
			continue
		}
		switch O {
		case "OR":
			inputs[OUT] = LV | RV
		case "AND":
			inputs[OUT] = LV & RV
		case "XOR":
			inputs[OUT] = LV ^ RV
		}
		done[op] = true
	}
	bin := ""
	for i := 0 ;; i++ {
		k := fmt.Sprintf("z%02d", i)
		if v, ok := inputs[k]; ok {
			bin = strconv.Itoa(v) + bin
		} else {
			break
		}
	}
	out, _ := strconv.ParseInt(bin, 2, 64)
	return int(out)
}

func main() {
	fmt.Println("Part 1:", part1())
}
