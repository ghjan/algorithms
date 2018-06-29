package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func initGraphDirection() Graph {
	var g Graph
	n1, n2, n3, n4, n5, n6, n7, n8 := GNode{1}, GNode{2}, GNode{3}, GNode{4}, GNode{5},
		GNode{6}, GNode{7}, GNode{8}
	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)
	g.AddNode(&n5)
	g.AddNode(&n6)
	g.AddNode(&n7)
	g.AddNode(&n8)
	g.AddEdgeDirection(&n1, &n2)
	g.AddEdgeDirection(&n1, &n5)
	g.AddEdgeDirection(&n2, &n3)
	g.AddEdgeDirection(&n2, &n4)
	g.AddEdgeDirection(&n2, &n5)
	g.AddEdgeDirection(&n3, &n4)
	g.AddEdgeDirection(&n4, &n6)
	g.AddEdgeDirection(&n4, &n8)
	g.AddEdgeDirection(&n8, &n7)
	g.AddEdgeDirection(&n5, &n8)
	return g
}

func TestAdd(t *testing.T) {
	g := initGraph()

	g.String()
}

func TestBFS(t *testing.T) {
	g := initGraph()
	g.BFS(func(node *GNode) {
		fmt.Printf("[Current Traverse GNode]: %v\n", node)
	})
}

func TestGraph_Unweighted(t *testing.T) {
	g := initGraphDirection()
	g.String()
	dist := make(map[GNode]int)
	path := make(map[GNode]*GNode)
	var Q GNodeQueue

	Q.New()
	for _, node := range g.nodes {
		if node == nil {
			continue
		}
		dist[*node] = -1
		path[*node] = nil
	}

	source := g.nodes[0]

	g.UnweightedShortestPath(*source, dist, path, Q)
	target := g.nodes[3]

	result := GetPath(path, source, target)
	fmt.Printf("shortest path from %d to %d:%s\n", source.value, target.value, result)

	assert.Equal(t, "1 2 4 ", result, fmt.Sprintf("expected path is:%s, actual is:%s", "1 2 4 ", result))

	target = g.nodes[7]
	result = GetPath(path, source, target)
	fmt.Printf("shortest path from %d to %d:%s\n", source.value, target.value, result)

	assert.Equal(t, "1 5 8 ", result, fmt.Sprintf("expected path is:%s, actual is:%s", "1 5 8 ", result))

}
