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

func getPrices(secret, n int) []int {
	prices := []int{secret % 10}
	for i := 0; i < n; i++ {
		secret = genSecret(secret, 1)
		prices = append(prices, secret % 10)
	}
	return prices
}

func getPriceChanges(prices []int) []int {
	changes := []int{}
	for i := 1; i < len(prices); i++ {
		changes = append(changes, prices[i] - prices[i-1])
	}
	return changes
}

func part1() int {
	secrets := readInput("input.txt")
	sum := 0
	for _, s := range secrets {
		sum += genSecret(s, 2000)
	}
	return sum
}

type sequence [4]int

func part2() int {
	secrets := readInput("input.txt")
	bananas := map[sequence]int{}
	for _, s := range secrets {
		prices := getPrices(s, 2000)
		priceChanges := getPriceChanges(prices)
		seen := map[sequence]bool{}
		for i := 0; i <= len(priceChanges)-4; i++ {
			seq := sequence(priceChanges[i:i+4])
			if _, ok := seen[seq]; !ok {
				bananas[seq] += prices[i+4]
				seen[seq] = true
			}
		}
	}
	mostBananas := -1
	for _, b := range bananas {
		if b > mostBananas {
			mostBananas = b
		}
	}
	return mostBananas
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
