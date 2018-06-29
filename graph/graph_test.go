package graph

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {
	g := initGraph()

	g.String()
}

func initGraph() Graph {
	var g Graph
	n1, n2, n3, n4, n5 := GNode{1}, GNode{2}, GNode{3}, GNode{4}, GNode{5}
	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)
	g.AddNode(&n5)
	g.AddEdge(&n1, &n2)
	g.AddEdge(&n1, &n5)
	g.AddEdge(&n2, &n3)
	g.AddEdge(&n2, &n4)
	g.AddEdge(&n2, &n5)
	g.AddEdge(&n3, &n4)
	g.AddEdge(&n4, &n5)
	return g
}

func TestBFS(t *testing.T) {
	g := initGraph()
	g.BFS(func(node *GNode) {
		fmt.Printf("[Current Traverse GNode]: %v\n", node)
	})
}

func TestGraph_Unweighted(t *testing.T) {
	g := initGraph()
	var dist map[GNode]int
	var path map[GNode]*GNode
	var Q GNodeQueue
	for _, node := range g.nodes {
		dist[*node] = -1
		path[*node] = nil
	}
	Q.New()

	g.Unweighted(*g.nodes[0], dist, path, Q)

}
