package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initGraph() Graph {
	var g Graph
	for i := 0; i < 8; i++ {
		g.AddNode(&GNode{i})
	}
	g.AddEdge(1, 2, false)
	g.AddEdge(1, 5, false)
	g.AddEdge(2, 3, false)
	g.AddEdge(2, 4, false)
	g.AddEdge(2, 5, false)
	g.AddEdge(3, 4, false)
	g.AddEdge(4, 6, false)
	return g
}

func initGraphDirection() Graph {
	var g Graph
	for i := 0; i < 8; i++ {
		g.AddNode(&GNode{i})
	}
	g.AddEdgeDirection(1, 2, false)
	g.AddEdgeDirection(1, 5, false)
	g.AddEdgeDirection(2, 3, false)
	g.AddEdgeDirection(2, 4, false)
	g.AddEdgeDirection(2, 5, false)
	g.AddEdgeDirection(3, 4, false)
	g.AddEdgeDirection(4, 6, false)
	g.AddEdgeDirection(4, 8, false)
	g.AddEdgeDirection(8, 7, false)
	g.AddEdgeDirection(5, 8, false)
	return g
}

func TestGraph_Add(t *testing.T) {
	fmt.Println("--------TestGraph_Add--------")
	g := initGraph()

	g.String()
}

func TestGraph_BFS(t *testing.T) {
	fmt.Println("--------TestGraph_BFS--------")
	g := initGraph()
	g.BFS(func(nodeIndex int) {
		fmt.Printf("[Current Traverse GNode]: %v\n", g.nodes[nodeIndex])
	})
}

func TestGraph_Unweighted(t *testing.T) {
	fmt.Println("--------TestGraph_Unweighted--------")
	g := initGraphDirection()
	g.String()

	source := 0

	path := g.Unweighted(source)
	target := 3

	result := g.GetPath(path, source, target)
	fmt.Printf("shortest path from %d to %d:%s\n", g.nodes[source].value, g.nodes[target].value, result)

	assert.Equal(t, "0 1 3 ", result, fmt.Sprintf("expected path is:%s, actual is:%s", "0 1 3 ", result))

	target = 7
	result = g.GetPath(path, source, target)
	fmt.Printf("shortest path from %d to %d:%s\n", g.nodes[source].value, g.nodes[target].value, result)

	assert.Equal(t, "0 4 7 ", result, fmt.Sprintf("expected path is:%s, actual is:%s", "0 4 7 ", result))

}
