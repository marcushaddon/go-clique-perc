package main

import (
	"fmt"

	"github.com/soniakeys/bits"
	"github.com/soniakeys/graph"
)

// GetSetIntersection returns a map with keys equal to
// numbers present in both a1 and a2
func GetSetIntersection(a1, a2 []int) map[int]bool {
	// Create dictionaries from each set
	var m1 = make(map[int]bool)
	var m2 = make(map[int]bool)
	var intersection = make(map[int]bool)
	for _, n := range a1 {
		m1[n] = true
	}

	for _, n := range a2 {
		m2[n] = true
	}

	for k := range m1 {
		if m2[k] {
			intersection[k] = true
		}
	}

	return intersection
}

// GetPercolatedCliques returns an array of cliques of size k,
// and an undirected graph in which an edge between nodes n and m
// indicates that the nodes at indices n and m in the clique array
// are adjacent.
func GetPercolatedCliques(g graph.LabeledUndirected, k int) ([][]int, graph.LabeledUndirected) {
	// Our cliques (index of clique is it's ID)
	var cliques [][]int

	// Our graph of clique adjacencies
	var cliqueAdjacencyGraph graph.LabeledUndirected

	g.BronKerbosch1(func(c bits.Bits) bool {
		// Add each clique to percolation graph
		clique := c.Slice()
		if len(clique) == k {
			cliques = append(cliques, clique)
		}
		return true
	})

	for id1, clique1 := range cliques {
		for id2 := id1 + 1; id2 < len(cliques); id2++ {
			clique2 := cliques[id2]
			if len(GetSetIntersection(clique1, clique2)) == k-1 {
				cliqueAdjacencyGraph.AddEdge(
					graph.Edge{
						N1: graph.NI(id1),
						N2: graph.NI(id2),
					},
					0)
			}
		}
	}

	return cliques, cliqueAdjacencyGraph
}

func main() {
	fmt.Println("------------")
	var g graph.LabeledUndirected

	nodes := []graph.Edge{
		graph.Edge{1, 0},
		graph.Edge{1, 2},
		graph.Edge{2, 0},

		graph.Edge{0, 3},
		graph.Edge{3, 1},

		graph.Edge{4, 5},
		graph.Edge{4, 3},
		graph.Edge{5, 3},
		graph.Edge{5, 1},
	}

	for _, n := range nodes {
		g.AddEdge(n, 0)
	}

	fmt.Println("GRAPH: ", g)
	cliques, adjs := GetPercolatedCliques(g, 3)

	adjs.Edges(func(e graph.LabeledEdge) {
		fmt.Printf("%d %c\n", e.Edge, e.LI)
	})
	fmt.Println("CLIQUES: ", cliques)
}
