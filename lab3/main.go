package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/twmb/algoimpl/go/graph"
)

type graphInfo struct {
	graph    graph.Graph
	nodes    []graph.Node
	indices  map[graph.Node]int
	weights  map[edge]int
	visited  map[graph.Node]bool
	distance map[graph.Node]int
	parent   map[graph.Node]graph.Node
}

type edge struct {
	from, to graph.Node
}

func main() {
	in := "In5.txt"
	g, o, d := readFile(in)
	orig := g.nodes[o-1]
	dest := g.nodes[d-1]
	findPaths(&g, orig)
	path, length, success := getPath(&g, orig, dest)
	if !success {
		fmt.Println("N")
		return
	}
	fmt.Println("Y")
	for _, v := range path {
		fmt.Printf("%d ", g.indices[v]+1)
	}
	fmt.Println()
	fmt.Println(length)
}

func readFile(file string) (graph graphInfo, origin int, destination int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Could not open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	size := readNumber(scanner)
	graph = createGraph(scanner, size)
	origin = readNumber(scanner)
	destination = readNumber(scanner)
	return
}

func createGraph(scanner *bufio.Scanner, size int) graphInfo {
	g := *graph.New(graph.Directed)
	nodes := make([]graph.Node, size)
	indices := make(map[graph.Node]int)
	weights := make(map[edge]int)
	for i := 0; i < size; i++ {
		nodes[i] = g.MakeNode()
		indices[nodes[i]] = i
	}

	for i := 0; i < size && scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), " ")
		for j := 0; ; {
			v, _ := strconv.Atoi(line[j])
			if v == 0 {
				break
			}
			weight, _ := strconv.Atoi(line[j+1])
			g.MakeEdge(nodes[v-1], nodes[i])
			weights[edge{nodes[v-1], nodes[i]}] = weight
			j += 2
		}
	}

	visited := make(map[graph.Node]bool)
	parent := make(map[graph.Node]graph.Node)
	distance := make(map[graph.Node]int)

	return graphInfo{g, nodes, indices, weights, visited, distance, parent}
}

func readNumber(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		os.Exit(-1)
	}
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func findPaths(g *graphInfo, u graph.Node) {
	for _, n := range g.nodes {
		g.distance[n] = math.MaxInt64
	}
	g.distance[u] = 0
	empty := graph.Node{}
	for range g.nodes {
		v := empty
		for _, s := range g.nodes {
			if !g.visited[s] && (v == empty || g.distance[s] < g.distance[v]) {
				v = s
			}
		}
		if g.distance[v] == math.MaxInt64 {
			break
		}
		g.visited[v] = true
		for _, t := range g.graph.Neighbors(v) {
			if g.distance[v]+g.weights[edge{v, t}] < g.distance[t] {
				g.distance[t] = g.distance[v] + g.weights[edge{v, t}]
				g.parent[t] = v
			}
		}
	}
}

func getPath(g *graphInfo, u graph.Node, v graph.Node) (path []graph.Node, length int, success bool) {
	if !g.visited[v] {
		return
	}
	path = make([]graph.Node, 0)
	for t := v; t != u; t = g.parent[t] {
		path = append(path, t)
		length += g.weights[edge{g.parent[t], t}]
	}
	path = append(path, u)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	success = true
	return
}
