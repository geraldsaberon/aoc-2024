package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PRUNE int = 16777216

func readInput(inputFile string) []int {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	secrets := []int{}
	for _, line := range lines {
		s, _ := strconv.Atoi(line)
		secrets = append(secrets, s)
	}
	return secrets
}

func genSecret(secret, nth int) int {
	for i := 0; i < nth; i++ {
		a := ((secret * 64) ^ secret) % PRUNE
		b := ((a / 32) ^ a) % PRUNE
		c := ((b * 2048) ^ b) % PRUNE
		secret = c
	}
	return secret
}

func part1() int {
	secrets := readInput("input.txt")
	sum := 0
	for _, s := range secrets {
		sum += genSecret(s, 2000)
	}
	return sum
}

func main() {
	fmt.Println("Part 1:", part1())
}
