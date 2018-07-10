package graph

import (
	"fmt"
	"strconv"
	"strings"
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
	fmt.Println("--------TestMGraph_DFS--------")
	gg := initMGraph()
	//gg.PrintMatrix(len(gg.vexs))
	vexs := ""
	gg.DFS(func(i int) {
		vexs += fmt.Sprintf("%s ", gg.vexs[i])
	})
	vexs = strings.TrimRight(vexs, " ")
	assert.Equal(t, "A B D C F E G", vexs)
	fmt.Println(vexs)
}

func TestMGraph_BFS(t *testing.T) {
	fmt.Println("--------TestMGraph_BFS--------")
	gg := initMGraph()
	vexs := ""
	gg.BFS(func(v VertexType) {
		vexs += fmt.Sprintf("%s ", v)
	})
	vexs = strings.TrimRight(vexs, " ")
	assert.Equal(t, "A B D E C F G", vexs)
	fmt.Println(vexs)
}

func TestMGraph_Dijkstra(t *testing.T) {
	gg := initMGraph()
	start := 0
	dist, path := gg.Dijkstra(start)

	for i := 0; i < gg.vexNum; i++ {
		fmt.Printf("shortest %s->%s = %d;", gg.vexs[start], gg.vexs[i], dist[i])
		realPath := strings.Split(gg.GetPathDijkstra(path, start, i), " ")
		for _, rp := range realPath {
			if index, err := strconv.Atoi(rp); index >= 0 && index < gg.vexNum && err == nil {
				fmt.Printf("%s", gg.vexs[index])
			}
		}
		fmt.Println()
	}
	assert.Equal(t, 0, dist[0])
	assert.Equal(t, 2, dist[1])
	assert.Equal(t, 3, dist[2])
	assert.Equal(t, 1, dist[3])
	assert.Equal(t, 3, dist[4])
	assert.Equal(t, 6, dist[5])
	assert.Equal(t, 5, dist[6])
}

func TestMGraph_Floyd(t *testing.T) {
	gg := initMGraph()
	start := 0
	if dist, path, err := gg.Floyd(); err == nil {
		for i := 0; i < gg.vexNum; i++ {
			fmt.Printf("shortest %s->%s = %d;", gg.vexs[start], gg.vexs[i], dist[start][i])
			realPath := strings.Split(gg.GetPathFloyd(path, start, i), " ")
			for _, rp := range realPath {
				if index, err := strconv.Atoi(rp); index >= 0 && index < gg.vexNum && err == nil {
					fmt.Printf("%s", gg.vexs[index])
				}
			}
			fmt.Println()
		}
		assert.Equal(t, 0, dist[start][0])
		assert.Equal(t, 2, dist[start][1])
		assert.Equal(t, 3, dist[start][2])
		assert.Equal(t, 1, dist[start][3])
		assert.Equal(t, 3, dist[start][4])
		assert.Equal(t, 6, dist[start][5])
		assert.Equal(t, 5, dist[start][6])

	} else {
		fmt.Println(err)
	}

}

func TestMGraph_Prim(t *testing.T) {
	gg := initMGraph()
	start := 0
	gg.Prim(start)
	gg.PrintMatrix(gg.vexNum)
}
