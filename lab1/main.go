package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/twmb/algoimpl/go/graph"
)

type graphInfo struct {
	graph   graph.Graph
	nodes   []graph.Node
	colors  map[graph.Node]bool
	visited map[graph.Node]bool
}

func main() {
	in := "In2.txt"
	m := readFile(in)
	g := createGraph(m)
	if !isBipartite(&g, g.nodes[0]) {
		fmt.Println("N")
		return
	}
	fmt.Println("Y")
	part1 := make([]int, 0)
	part2 := make([]int, 0)
	for i, v := range g.nodes {
		if g.colors[v] {
			part1 = append(part1, i+1)
		} else {
			part2 = append(part2, i+1)
		}
	}
	if part1[0] < part2[0] {
		printParts(part1, part2)
	} else {
		printParts(part2, part1)
	}
}

func printPart(part []int) {
	for _, v := range part {
		fmt.Println(v)
	}
}

func printParts(part1 []int, part2 []int) {
	printPart(part1)
	fmt.Println("0")
	printPart(part2)
}

func createGraph(matrix [][]bool) graphInfo {
	size := len(matrix)
	g := *graph.New(graph.Undirected)
	nodes := make([]graph.Node, size)
	for i := 0; i < size; i++ {
		nodes[i] = g.MakeNode()
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[i][j] {
				g.MakeEdge(nodes[i], nodes[j])
			}
		}
	}

	colors := make(map[graph.Node]bool)
	visited := make(map[graph.Node]bool)

	return graphInfo{g, nodes, colors, visited}
}

func readFile(file string) (adj [][]bool) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Could not open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	size := readNumber(scanner)
	return readMatrix(scanner, size, size)
}

func readNumber(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		os.Exit(-1)
	}
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func readMatrix(scanner *bufio.Scanner, x int, y int) [][]bool {
	adj := make([][]bool, x)
	for i := 0; i < x && scanner.Scan(); i++ {
		adj[i] = make([]bool, y)
		line := strings.Split(scanner.Text(), " ")
		for j := 0; j < y; j++ {
			if line[j] == "1" {
				adj[i][j] = true
			}
		}
	}
	return adj
}

func isBipartite(info *graphInfo, u graph.Node) bool {
	for _, v := range info.graph.Neighbors(u) {
		if !info.visited[v] {
			info.visited[v] = true
			info.colors[v] = !info.colors[u]
			if !isBipartite(info, v) {
				return false
			}
		} else if info.colors[u] == info.colors[v] {
			return false
		}
	}

	return true
}
