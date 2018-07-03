package graph

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func initMGraph() MGraph {
	var gg MGraph
	var vexs = []VertexType{"A", "B", "C", "D", "E", "F", "G"}
	gg.initMGraph(vexs)

	gg.matrix[0][1] = 2
	gg.matrix[0][3] = 1

	gg.matrix[1][3] = 3
	gg.matrix[1][4] = 10

	gg.matrix[2][0] = 4
	gg.matrix[2][5] = 5

	gg.matrix[3][2] = 2
	gg.matrix[3][4] = 2
	gg.matrix[3][5] = 8
	gg.matrix[3][6] = 4

	gg.matrix[4][6] = 6

	gg.matrix[6][5] = 1

	gg.edgeNum = 12
	return gg
}

func TestMGraph_DFS(t *testing.T) {
	gg := initMGraph()
	gg.PrintMatrix(len(gg.vexs))
	gg.DFS()
}

func TestMGraph_BFS(t *testing.T) {
	gg := initMGraph()
	gg.BFS()
}

func TestMGraph_Dijkstra(t *testing.T) {
	gg := initMGraph()
	start := 0
	dist := gg.Dijkstra(start)

	for i := 0; i < gg.vexNum; i++ {
		fmt.Printf("shortest %s->%s = %d\n", gg.vexs[start], gg.vexs[i], dist[i])
	}
	assert.Equal(t, 0, dist[0])
	assert.Equal(t, 3, dist[4])
	assert.Equal(t, 6, dist[5])
}

func TestMGraph_Prim(t *testing.T) {
	gg := initMGraph()
	start := 0
	gg.Prim(start)
	gg.PrintMatrix(gg.vexNum)
}
