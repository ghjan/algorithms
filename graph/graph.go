package graph

import "fmt"

type Graph struct {
	Vertexes int
	Edges    [][]int
}

// 创建具有 n 个节点的无权、无向图
func NewGraph(n int) *Graph {
	return &Graph{
		Vertexes: n,
		Edges:    make([][]int, n),
	}
}

// 向无向图中添加边
func (g *Graph) AddEdge(u, v int) {
	// 双向边的效果
	g.Edges[u] = append(g.Edges[u], v)
	g.Edges[v] = append(g.Edges[v], u)
}

// 将邻接好的图输出
func (g *Graph) AdjacentEdges() {
	for u, vertexes := range g.Edges {
		for _, v := range vertexes {
			// 输出从 u 到 v 的边
			fmt.Printf("[Eage]:\t%v -> %v\n", u, v)
		}
	}
}
