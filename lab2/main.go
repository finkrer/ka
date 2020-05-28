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
	nodes   map[point]graph.Node
	coords  map[graph.Node]point
	visited map[graph.Node]bool
	parent  map[graph.Node]graph.Node
}

type point struct {
	x int
	y int
}

func main() {
	in := "In2.txt"
	m, o, d := readFile(in)
	g := createGraph(m)
	orig := g.nodes[o]
	dest := g.nodes[d]
	findPath(&g, orig, dest)
	if _, ok := g.parent[dest]; !ok {
		fmt.Println("N")
		return
	}
	fmt.Println("Y")
	n := dest
	output := make([]string, 0)
	for n != orig {
		c := g.coords[n]
		s := fmt.Sprintf("%d %d\n", c.y+1, c.x+1)
		output = append(output, s)
		n = g.parent[n]
	}
	s := fmt.Sprintf("%d %d\n", o.y+1, o.x+1)
	output = append(output, s)
	for i := len(output) - 1; i >= 0; i-- {
		fmt.Print(output[i])
	}
}

func createGraph(matrix [][]bool) graphInfo {
	h := len(matrix)
	w := len(matrix[0])
	g := *graph.New(graph.Undirected)

	nodes := make(map[point]graph.Node)
	coords := make(map[graph.Node]point)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := point{x, y}
			nodes[p] = g.MakeNode()
			coords[nodes[p]] = p
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if matrix[y][x] {
				adj := getAdjacent(matrix, point{x, y})
				for _, v := range adj {
					if matrix[v.y][v.x] {
						g.MakeEdge(nodes[point{x, y}], nodes[point{v.x, v.y}])
					}
				}
			}
		}
	}

	visited := make(map[graph.Node]bool)
	parent := make(map[graph.Node]graph.Node)

	return graphInfo{g, nodes, coords, visited, parent}
}

func readMatrix(scanner *bufio.Scanner, x int, y int) [][]bool {
	free := make([][]bool, y)
	for i := 0; i < y && scanner.Scan(); i++ {
		free[i] = make([]bool, x)
		line := strings.Split(scanner.Text(), " ")
		for j := 0; j < x; j++ {
			if line[j] == "0" {
				free[i][j] = true
			}
		}
	}
	return free
}

func getAdjacent(matrix [][]bool, p point) []point {
	h := len(matrix)
	w := len(matrix[0])
	adj := make([]point, 0)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 || i != 0 && j != 0 {
				continue
			}
			x := p.x + i
			y := p.y + j
			if x > 0 && x < w && y > 0 && y < h {
				adj = append(adj, point{x, y})
			}
		}
	}

	return adj
}

func readFile(file string) (matrix [][]bool, origin point, destination point) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Could not open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	y := readNumber(scanner)
	x := readNumber(scanner)
	matrix = readMatrix(scanner, x, y)
	origin = readPoint(scanner)
	destination = readPoint(scanner)
	return
}

func readNumber(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		os.Exit(-1)
	}
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func readPoint(scanner *bufio.Scanner) point {
	if !scanner.Scan() {
		os.Exit(-1)
	}
	t := scanner.Text()
	cs := strings.Split(t, " ")
	x, _ := strconv.Atoi(cs[0])
	y, _ := strconv.Atoi(cs[1])
	return point{y - 1, x - 1}
}

func findPath(i *graphInfo, u graph.Node, v graph.Node) {
	queue := make([]graph.Node, 0)
	i.visited[u] = true
	queue = append(queue, u)
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		if t == v {
			return
		}
		for _, s := range i.graph.Neighbors(t) {
			if !i.visited[s] {
				i.visited[s] = true
				i.parent[s] = t
				queue = append(queue, s)
			}
		}
	}
}
