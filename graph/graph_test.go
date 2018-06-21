package graph

import (
	"testing"
)

func TestAdjacentEdges(t *testing.T) {
	g := NewGraph(6)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(3, 5)
	g.AdjacentEdges()
}
