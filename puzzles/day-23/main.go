package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type vertex string
type graph map[vertex][]vertex

func (g graph) vertices() []vertex {
	vs := []vertex{}
	for v := range g {
		vs = append(vs, v)
	}
	return vs
}

type set[T comparable] map[T]bool

func (s set[T]) add(a ...T) {
	for _, v := range a {
		s[v] = true
	}
}

func (s set[T]) remove(a ...T) {
	for _, v := range a {
		delete(s, v)
	}
}

func (s set[T]) union(t set[T]) set[T] {
	st := set[T]{}
	for si := range s {
		st[si] = true
	}
	for ti := range t {
		st[ti] = true
	}
	return st
}

func (s set[T]) intersection(t set[T]) set[T] {
	i := set[T]{}
	for si := range s {
		if _, ok := t[si]; ok {
			i[si] = true
		}
	}
	return i
}

func (s set[T]) values() []T {
	vs := []T{}
	for v := range s {
		vs = append(vs, v)
	}
	return vs
}

func (s set[T]) String() string {
	out := "set["
	for v := range s {
		out += fmt.Sprintf("%v ", v)
	}
	return strings.TrimSpace(out) + "]"
}

func readInput(inputFile string) graph {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	g := graph{}
	for _, line := range lines {
		split := strings.Split(line, "-")
		a, b := vertex(split[0]), vertex(split[1])
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	return g
}

func areAdjacent(g graph, a, b vertex) bool {
	if slices.Contains(g[a], b) || slices.Contains(g[b], a) {
		return true
	}
	return false
}

func bronKerbosch(R, P, X set[vertex], g graph, max *[]vertex) {
	if len(R) > len(*max) {
		*max = R.values()
	}
	if len(P) == 0 && len(X) == 0 {
		return
	}
	for v := range P {
		vNeighbors := set[vertex]{}
		vNeighbors.add(g[v]...)
		sv := set[vertex]{}
		sv.add(v)
		bronKerbosch(R.union(sv), P.intersection(vNeighbors), X.intersection(vNeighbors), g, max)
		P.remove(v)
		X.add(v)
	}
}

func part1() int {
	g := readInput("input.txt")
	threes := map[string]bool{}
	for    a := range g {
	for _, b := range g[a] {
	for _, c := range g[a] {
		if areAdjacent(g, b, c) && (a[0] == 't' || b[0] == 't' || c[0] == 't') {
			t := []vertex{a, b, c}
			slices.Sort(t)
			threes[fmt.Sprint(t)] = true
		}
	}}}
	return len(threes)
}

func part2() string {
	g := readInput("input.txt")
	R, P, X := set[vertex]{}, set[vertex]{}, set[vertex]{}
	P.add(g.vertices()...)
	res := []vertex{}
	bronKerbosch(R, P, X, g, &res)
	slices.Sort(res)
	return strings.Trim(strings.Replace(fmt.Sprint(res), " ", ",", -1), "[]")
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
