package graph

import (
	"testing"
)

func TestAdjacentEdges(t *testing.T) {
	g := NewGraph(6)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(1, 4, 1)
	g.AddEdge(3, 5, 3)
	g.AdjacentEdges()
}
