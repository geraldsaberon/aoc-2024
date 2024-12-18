package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type register struct {
	a, b, c int
}

func readInput(inputFile string) (register, []int) {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	reg := register{}
	program := []int{}
	reg.a, _ = strconv.Atoi(strings.TrimLeft(lines[0], "Register A: "))
	reg.b, _ = strconv.Atoi(strings.TrimLeft(lines[1], "Register B: "))
	reg.c, _ = strconv.Atoi(strings.TrimLeft(lines[2], "Register C: "))
	for _, c := range strings.TrimLeft(lines[len(lines)-1], "Program: ") {
		if c != ',' {
			i, _ := strconv.Atoi(string(c))
			program = append(program, i)
		}
	}
	return reg, program
}

func combo(op int, reg register) int {
	if op == 4 {
		return reg.a
	}
	if op == 5 {
		return reg.b
	}
	if op == 6 {
		return reg.c
	}
	return op
}

func outJoin(s []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), ","), "[]")
}

func machine(reg register, program []int) string {
	out := []int{}
	for ip := 0; ip < len(program); ip += 2 {
		ins, op := program[ip], program[ip+1]
		switch ins {
		case 0:
			reg.a = reg.a / (1 << combo(op, reg))
		case 1:
			reg.b = reg.b ^ op
		case 2:
			reg.b = combo(op, reg) % 8
		case 3:
			if reg.a > 0 {
				ip = op-2
				continue
			}
		case 4:
			reg.b = reg.b ^ reg.c
		case 5:
			out = append(out, combo(op, reg) % 8)
		case 6:
			reg.b = reg.a / (1 << combo(op, reg))
		case 7:
			reg.c = reg.a / (1 << combo(op, reg))
		}
	}
	return outJoin(out)
}

func part1() string {
	reg, program := readInput("input.txt")
	out := machine(reg, program)
	return out
}

func part2() int {
	return -1
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
