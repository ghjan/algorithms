package graph

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestMGraph() MGraph {
	var gg MGraph
	var vexes = []VertexType{"A", "B", "C", "D", "E", "F", "G"}
	gg.InitMGraph(vexes)

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

func buildVilvilGraph() MGraph{
	return MGraph{}
}

func TestMGraph_DFS(t *testing.T) {
	fmt.Println("--------TestMGraph_DFS--------")
	gg := createTestMGraph()
	//gg.PrintWeight(len(gg.vexes))
	vexs := ""
	gg.DFS(func(i int) {
		vexs += fmt.Sprintf("%s ", gg.vexes[i])
	})
	vexs = strings.TrimRight(vexs, " ")
	assert.Equal(t, "A B D C F E G", vexs)
	fmt.Println(vexs)
}

func TestMGraph_BFS(t *testing.T) {
	fmt.Println("--------TestMGraph_BFS--------")
	gg := createTestMGraph()
	vexs := ""
	gg.BFS(func(v VertexType) {
		vexs += fmt.Sprintf("%s ", v)
	})
	vexs = strings.TrimRight(vexs, " ")
	assert.Equal(t, "A B D E C F G", vexs)
	fmt.Println(vexs)
}

func TestMGraph_Dijkstra(t *testing.T) {
	fmt.Println("--------TestMGraph_Dijkstra--------")
	gg := createTestMGraph()
	start := 0
	dist, path := gg.Dijkstra(start)

	for i := 0; i < gg.vexNum; i++ {
		fmt.Printf("shortest %s->%s = %d;", gg.vexes[start], gg.vexes[i], dist[i])
		realPath := strings.Split(gg.GetPathDijkstra(path, start, i), " ")
		for _, rp := range realPath {
			if index, err := strconv.Atoi(rp); index >= 0 && index < gg.vexNum && err == nil {
				fmt.Printf("%s", gg.vexes[index])
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
	fmt.Println("--------TestMGraph_Floyd--------")
	gg := createTestMGraph()
	start := 0
	if dist, path, err := gg.Floyd(); err == nil {
		for i := 0; i < gg.vexNum; i++ {
			fmt.Printf("shortest %s->%s = %d;", gg.vexes[start], gg.vexes[i], dist[start][i])
			if realPath, err := gg.GetPathFloyd(path, start, i); err == nil {
				realPath := strings.Split(realPath, " ")
				for _, rp := range realPath {
					if index, err := strconv.Atoi(rp); index >= 0 && index < gg.vexNum && err == nil {
						fmt.Printf("%s", gg.vexes[index])
					}
				}
				fmt.Println()
			} else {
				fmt.Println(err)
			}

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
	mg := createTestMGraph()
	start := 0
	//mst, sum, parent, err 分别表示最小生成树，加权和，边的起始点数组，错误信息
	if mst, sum, parent, err := mg.Prim(start); err == nil {
		for _, vert := range mst {
			minV := mg.getPosition(vert)
			fmt.Println(mg.vexes[parent[minV]], vert, "权重", mg.matrix[parent[minV]][minV])
		}
		fmt.Println("sum:", sum)
		fmt.Println(mst)
		assert.Equal(t, 12, sum)
		expected := make([]VertexType, mg.vexNum, mg.vexNum)
		arr := strings.Split("A D B C E G F", " ")
		for i := 0; i < mg.vexNum; i++ {
			expected[i] = VertexType(arr[i])
		}
		assert.Equal(t, expected, mst)
	} else {
		assert.Equal(t, "", err)
	}

}
